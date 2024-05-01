package api

import (
	"cmdTest/internal/response"
	"cmdTest/internal/server"
	"github.com/gin-gonic/gin"
)

// File
//
//	@Tags		文件模块
//	@Summary	查看对应路径下有哪些文件
//	@Accept		multipart/form-data
//	@Produce	application/json
//	@Param		filePath	query		string	true	"文件路径"
//	@Success	200			{object}	response.Response
//	@Router		/api/v1/file/ [get]
func File(c *gin.Context) {
	path := c.Query("filePath")
	files, err := server.GetFile(path)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", files)
}

// CsvUpload
//
//	@Tags		文件模块
//	@Summary	上传数据集文件到对应路径
//	@Accept		multipart/form-data
//	@Produce	application/json
//	@Param		file		formData	file	true	"csv文件"
//	@Param		filePath	formData	string	true	"文件路径"
//	@Success	200			{object}	response.Response
//	@Router		/api/v1/file/csv [post]
func CsvUpload(c *gin.Context) {
	path := c.PostForm("filePath")
	fileData, err := c.FormFile("file")
	if err != nil {
		response.Fail(c.Writer, "wrong file data", 400)
		return
	}

	err = server.SaveCsvFile(path, fileData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", nil)
}

// NewFilePath
//
//	@Tags		文件模块
//	@Summary	新建数据集文件夹
//	@Produce	application/json
//	@Param		filePath	formData	string	true	"文件路径"
//	@Success	200			{object}	response.Response
//	@Router		/api/v1/file/new [post]
func NewFilePath(c *gin.Context) {
	path := c.PostForm("filePath")

	err := server.NewFilePath(path)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", nil)
}

// GetAllPath
//
//	@Tags		文件模块
//	@Summary	获取所有数据文件夹
//	@Produce	application/json
//	@Success	200	{object}	response.Response
//	@Router		/api/v1/file/all [get]
func GetAllPath(c *gin.Context) {
	paths, err := server.GetAllPath()
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", paths)
}

// DeletePath
//
//	@Tags		文件模块
//	@Summary	删除数据文件夹
//	@Produce	application/json
//	@Param		filePath	formData	string	true	"文件路径"
//	@Success	200			{object}	response.Response
//	@Router		/api/v1/file/dp [post]
func DeletePath(c *gin.Context) {
	filePath := c.PostForm("filePath")

	err := server.DeletePath(filePath)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", nil)
}

// DeleteFile
//
//	@Tags		文件模块
//	@Summary	删除数据文件
//	@Produce	application/json
//	@Param		filePath	formData	string	true	"文件路径"
//	@Param		fileName	formData	string	true	"文件名称"
//	@Success	200			{object}	response.Response
//	@Router		/api/v1/file/df [post]
func DeleteFile(c *gin.Context) {
	filePath := c.PostForm("filePath")
	fileName := c.PostForm("fileName")

	err := server.DeleteFile(filePath, fileName)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", nil)
}
