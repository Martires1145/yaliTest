package api

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/response"
	"cmdTest/internal/server"
	"context"
	"github.com/gin-gonic/gin"
)

// NewModel
//
//	@Summary	新增模型
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

// DeleteModel
//
//	@Summary	删除模型
//	@Tags		Model
//	@Param		modelID	formData	string	true	"模型信息"
//	@Router		/api/v1/md/delete [post]
func DeleteModel(c *gin.Context) {
	id := c.PostForm("modelID")

	err := server.DeleteModel(id)

	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", 200)
}

func ModifyModel(c *gin.Context) {
	context.TODO()
}

func UseModel(c *gin.Context) {
	context.TODO()
}

func CopyModel(c *gin.Context) {
	context.TODO()
}

func GetModel(c *gin.Context) {
	context.TODO()
}
