package api

import (
	"context"
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
//	@Router		/api/v1/path [get]
func File(c *gin.Context) {
	context.TODO()
}

// CsvUpload
//
//	@Tags		文件模块
//	@Summary	上传数据集文件到对应路径
//	@Accept		multipart/form-data
//	@Produce	application/json
//	@Param		file		formData	file	true	"csv文件"
//	@Param		filePath	query		string	true	"文件路径"
//	@Success	200			{object}	response.Response
//	@Router		/api/v1/upload/csv [post]
func CsvUpload(c *gin.Context) {
	context.TODO()
}

// NewFilePath
//
//	@Tags		文件模块
//	@Summary	新建数据集文件夹
//	@Produce	application/json
//	@Param		filePath	formData	string	true	"文件路径"
//	@Success	200			{object}	response.Response
//	@Router		/api/v1/upload/new [post]
func NewFilePath(c *gin.Context) {
	context.TODO()
}
