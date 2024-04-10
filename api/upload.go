package api

import (
	"cmdTest/internal/response"
	"cmdTest/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CsvUpload
//
//	@Tags		上传模块
//	@Summary	上传对应类型的csv文件到对应路径
//	@Accept		multipart/form-data
//	@Produce	application/json
//	@Param		file		formData	file	true	"csv文件"
//	@Param		file_type	query		string	true	"文件类型"
//	@Success	200			{object}	response.Response
//	@Router		/api/v1/upload/csv [post]
func CsvUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	fileType := c.Query("file_type")
	if err != nil {
		response.Fail(c.Writer, fmt.Sprintf("wrong data, err:%s", err.Error()), 400)
		return
	}

	path := util.MakePath(file.Filename, fileType)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		response.Fail(c.Writer, fmt.Sprintf("save file failed, err:%s", err.Error()), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}
