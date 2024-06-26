package server

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/mq"
	"cmdTest/internal/response"
	"cmdTest/pkg/util"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"log"
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

	// 向模型提交用于交互的topic
	topicR, topicW := util.GetTopic()
	mq.WriteTopicID(context.Background(), topicR, topicW)

	// 接受模型方的ack信息
	if !mq.ReadTopicAck(context.Background(), topicR) {
		response.Fail(w, "no ack", http.StatusInternalServerError)
		return
	}

	// 向模型方发送ack消息
	mq.WriteTopicAck(context.Background(), topicW)

	// 启动读写进程
	if isStream {
		go read(clientID, topicW)
	}
	go write(clientID, topicR)
}

func read(clientID, topicW string) {
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

		err = mq.WriteMsg(context.Background(), topicW, jsonD)
		if err != nil {
			log.Printf("[mqError] ==== err: %s", err.Error())
			break
		}
	}
}

func write(clientID, topicR string) {
	conn := Conns[clientID]
	pip := make(chan []byte, 3)

	defer func() {
		conn.Close()
		delete(Conns, clientID)
	}()

	go mq.ReadMsg(context.Background(), topicR, &pip)

	for {
		data := <-pip
		if string(data) == finishMsg {
			break
		} else if string(data) == mq.TimeOut {
			log.Printf("[mqError] ==== err: %s", mq.TimeOut)
			break
		}
		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
