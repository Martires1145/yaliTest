package api

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/response"
	"cmdTest/internal/server"
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

	response.Success(c.Writer, "success", nil)
}

// ModifyModel
//
//	@Summary	修改模型
//	@Tags		Model
//	@Param		modelID		formData	string	true	"模型ID"
//	@Param		name		formData	string	true	"模型新名称"
//	@Param		useKafka	formData	int		true	"模型use-kafka按钮"
//	@Router		/api/v1/md/revise [post]
func ModifyModel(c *gin.Context) {
	id := c.PostForm("modelID")
	name := c.PostForm("name")
	useKafka := c.PostForm("useKafka")

	err := server.ModifyModel(id, name, useKafka)

	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", nil)
}

// UseModel
//
//	@Summary	使用模型
//	@Tags		Model
//	@Router		/api/v1/md/use [post]
func UseModel(c *gin.Context) {
	// todo
	panic("todo")
}

// CopyModel
//
//	@Summary	复制模型
//	@Tags		Model
//	@Param		modelID	formData	string	true	"模型ID"
//	@Param		name	formData	string	true	"新模型名称"
//	@Router		/api/v1/md/copy [post]
func CopyModel(c *gin.Context) {
	id := c.PostForm("modelID")
	name := c.PostForm("name")

	err := server.CopyModel(id, name)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", nil)
}

// GetModel
//
//	@Summary	获取所有模型信息
//	@Tags		Model
//	@Router		/api/v1/md/all [get]
func GetModel(c *gin.Context) {
	models, err := server.GetAllModel()
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", models)
}

func UploadModelFile(c *gin.Context) {
	// todo
	panic("todo")
}

// GetModelParams
//
//	@Summary	获取模型参数
//	@Tags		Model
//	@Param		modelID	formData	string	true	"模型ID"
//	@Router		/api/v1/md/params [get]
func GetModelParams(c *gin.Context) {
	id := c.PostForm("modelID")

	params, err := server.GetModelParams(id)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", params)
}
