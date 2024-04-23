package api

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/response"
	"cmdTest/internal/server"
	"github.com/gin-gonic/gin"
)

// NewModel
//
//	@Summary	新增模型参数
//	@Tags		Model
//	@Param		md	body	model.JsonModel	true	"模型信息"
//	@Router		/api/v1/md/new [post]
func NewModel(c *gin.Context) {
	var modelData model.JsonModel
	err := c.ShouldBindJSON(&modelData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 400)
		return
	}

	err = server.NewModel(&modelData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}
