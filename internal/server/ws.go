package server

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/mq"
	"cmdTest/internal/response"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"net/http"
)

var finishMsg = viper.GetString("kafka.finishMsg")

var UP = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var Conns map[string]*websocket.Conn

func init() {
	Conns = make(map[string]*websocket.Conn)
}

func Ws(w http.ResponseWriter, rq *http.Request, clientID string, isStream bool) {
	// 升级为webSocket连接
	conn, err := UP.Upgrade(w, rq, nil)
	if err != nil {
		response.Fail(w, err.Error(), http.StatusUpgradeRequired)
		return
	}

	// 将生成的连接加入全局连接字典中
	Conns[clientID] = conn

	// 启动读写进程
	if isStream {
		go read(clientID)
	}
	go write(clientID)
}

func read(clientID string) {
	conn := Conns[clientID]

	defer func() {
		conn.Close()
		delete(Conns, clientID)
	}()

	for {
		var list []string
		err := conn.ReadJSON(&list)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		data := model.MakeData(list)

		jsonD, err := data.Json()
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		mq.WriteMsg(context.Background(), jsonD)
	}
}

func write(clientID string) {
	conn := Conns[clientID]
	pip := make(chan []byte, 3)

	defer func() {
		conn.Close()
		delete(Conns, clientID)
	}()

	go mq.ReadMsg(context.Background(), &pip)

	for {
		data := <-pip
		if string(data) == finishMsg {
			break
		}
		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
