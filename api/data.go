package api

import (
	"cmdTest/internal/response"
	"cmdTest/internal/server"
	"github.com/gin-gonic/gin"
)

// NewForecastData
//
//	@Tags		历史数据模块
//	@Summary	上传数据集文件到对应路径
//	@Accept		multipart/form-data
//	@Produce	application/json
//	@Param		file	formData	file	true	"预测数据文件"
//	@Param		id		formData	int		true	"使用模型时产生的历史数据id"
//	@Success	200		{object}	response.Response
//	@Router		/api/v1/history/uf [post]
func NewForecastData(c *gin.Context) {
	id := c.PostForm("id")
	fileData, err := c.FormFile("file")
	if err != nil {
		response.Fail(c.Writer, "wrong file data", 400)
		return
	}

	err = server.NewForecastData(id, fileData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", nil)
}

// GetHistoryData
//
//	@Tags		历史数据模块
//	@Summary	查看所有历史数据
//	@Produce	application/json
//	@Success	200	{object}	response.Response
//	@Router		/api/v1/data/all [post]
func GetHistoryData(c *gin.Context) {
	histories, err := server.GetHistoryData()
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", histories)
}

// DeleteHistoryData
//
//	@Tags		历史数据模块
//	@Summary	删除历史数据
//	@Produce	application/json
//	@Param		id	formData	int	true	"使用模型时产生的历史数据id"
//	@Success	200	{object}	response.Response
//	@Router		/api/v1/data/delete [post]
func DeleteHistoryData(c *gin.Context) {
	id := c.PostForm("id")

	err := server.DeleteHistoryData(id)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", nil)
}

// DataDetailOpen
//
//	@Tags		历史数据模块
//	@Summary	打开查看历史数据进程
//	@Produce	application/json
//	@Param		id	formData	int	true	"使用模型时产生的历史数据id"
//	@Success	200	{object}	response.Response
//	@Router		/api/v1/data/delete [post]
func DataDetailOpen(c *gin.Context) {
	id := c.PostForm("id")

	err := server.DataDetailOpen(id)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", nil)
}

// GetDataDetailRanged
//
//	@Tags		历史数据模块
//	@Summary	打开查看历史数据进程
//	@Produce	application/json
//	@Param		id		formData	int	true	"使用模型时产生的历史数据id"
//	@Param		from	formData	int	true	"起始时间"
//	@Param		to		formData	int	true	"截止时间"
//	@Success	200		{object}	response.Response
//	@Router		/api/v1/data/delete [post]
func GetDataDetailRanged(c *gin.Context) {
	id := c.PostForm("id")
	from := c.PostForm("from")
	to := c.PostForm("to")

	rangeDataT, rangeDataP, err := server.GetRangeData(id, from, to)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", gin.H{
		"trueData":    rangeDataT,
		"predictData": rangeDataP,
	})
}

// DataDetailClose
//
//	@Tags		历史数据模块
//	@Summary	关闭查看历史数据进程
//	@Produce	application/json
//	@Param		id	formData	int	true	"使用模型时产生的历史数据id"
//	@Success	200	{object}	response.Response
//	@Router		/api/v1/data/delete [post]
func DataDetailClose(c *gin.Context) {
	id := c.PostForm("id")

	err := server.DataDetailClose(id)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
	}

	response.Success(c.Writer, "success", nil)
}
