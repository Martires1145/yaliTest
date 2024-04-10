package api

import (
	"cmdTest/internal/response"
	"cmdTest/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// File
//
//	@Tags		文件选项模块
//	@Summary	查看对应类别的路径下有哪些文件
//	@Accept		multipart/form-data
//	@Produce	application/json
//	@Param		file_type	query		string	true	"文件类型"
//	@Success	200			{object}	response.Response
//	@Router		/api/v1/path [get]
func File(c *gin.Context) {
	fileType := c.Query("file_type")
	root, err := util.GetFile(fileType)
	if err != nil {
		response.Fail(c.Writer, fmt.Sprintf("os wrong, err:%#v", err.Error()), 500)
		return
	}
	response.Success(c.Writer, "success", gin.H{"files": root})
}
