package api

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/response"
	"cmdTest/internal/server"
	"cmdTest/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// RunScript
//
//	@Tags		脚本模块
//	@Summary	运行脚本
//	@Produce	application/json
//	@Param		params	body		model.Params	true	"脚本运行的参数"
//	@Success	200		{object}	response.Response
//
//	@Router		/api/v1/rs [post]
func RunScript(c *gin.Context) {
	// 解析参数
	var params model.Params
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.Fail(c.Writer, fmt.Sprintf("params parse failed, err: %s", err.Error()), 400)
		return
	}

	// 运行脚本
	IsStream, err := server.Call(&params)
	if err != nil {
		response.Fail(c.Writer, fmt.Sprintf("wrong params, err: %s", err.Error()), 400)
		return
	}

	// 生成当次会话的id
	clientID := util.GetClientID()
	server.Conns[clientID] = nil

	response.Success(c.Writer, "success", gin.H{
		"is_stream": IsStream,
		"client_id": clientID,
	})
}
