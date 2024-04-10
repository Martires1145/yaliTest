package api

import (
	"cmdTest/internal/response"
	"cmdTest/internal/server"
	"github.com/gin-gonic/gin"
)

// DataChan
//
//	@Tags		数据传输模块
//	@Summary	指定形式传输数据
//	@Produce	application/json
//	@Param		client_id	query		string	true	"会话id"
//	@Param		is_stream	query		int		true	"是否为流式传输"
//	@Success	200			{object}	response.Response
//	@Router		/api/v1/dchan [get]
func DataChan(c *gin.Context) {
	clientID := c.Query("client_id")
	isStream := c.Query("is_stream")

	if _, ok := server.Conns[clientID]; !ok {
		response.Fail(c.Writer, "wrong clientID", 401)
		return
	}

	server.Ws(c.Writer, c.Request, clientID, isStream == "1")
}
