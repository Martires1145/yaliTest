package server

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/mq"
	"cmdTest/pkg/util"
	"context"
	"encoding/json"
	"errors"
	"github.com/spf13/viper"
)

var (
	cmdName    = viper.GetString("script.cmdName")
	successMsg = viper.GetString("script.successMsg")
)

func Call(params *model.Params) (bool, error) {
	// 用于收集运行信息的管道
	var runState chan string
	runState = make(chan string)

	// 处理参数
	args, err := params.Parse()
	if err != nil {
		return false, err
	}

	// 运行脚本
	go util.RunCmd(cmdName, args, runState)

	// 监听运行状态直至出错或成功
	msg := <-runState

	if msg != successMsg {
		return false, errors.New(msg)
	}

	// 查看是否要启动kafka流式传输
	return params.IsStream(), nil
}

func CallByKafka(params *model.Params) (bool, error) {
	// 处理数据
	paramsJson, _ := json.Marshal(params)

	// 发送参数
	err := mq.WriteParams(context.Background(), paramsJson)
	if err != nil {
		return false, err
	}

	// 获取结果
	useKafka, err := mq.ReadParamsResult(context.Background())
	return useKafka, err
}