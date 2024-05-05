package api

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/response"
	"cmdTest/internal/server"
	"github.com/gin-gonic/gin"
)

// NewHistoryData
//
//	@Tags		历史数据模块
//	@Summary	新建历史数据
//	@Accept		multipart/form-data
//	@Produce	application/json
//	@Param		fileTrue	formData	file					true	"真实数据文件"
//	@Param		filePredict	formData	file					true	"预测数据文件"
//	@Param		history		body		model.DataHistoryJson	true	"使用模型时产生的历史数据id"
//	@Success	200			{object}	response.Response
//	@Router		/api/v1/data/new [post]
func NewHistoryData(c *gin.Context) {
	fileDataTrue, err := c.FormFile("fileTrue")
	if err != nil {
		response.Fail(c.Writer, "wrong file data", 400)
		return
	}

	filePredict, err := c.FormFile("filePredict")
	if err != nil {
		response.Fail(c.Writer, "wrong file data", 400)
		return
	}

	var history model.DataHistoryJson
	err = c.ShouldBind(&history)
	if err != nil {
		response.Fail(c.Writer, "wrong data format", 400)
		return
	}

	err = server.NewHistoryData(fileDataTrue, filePredict, &history)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// GetHistoryData
//
//	@Tags		历史数据模块
//	@Summary	查看所有历史数据
//	@Produce	application/json
//	@Success	200	{object}	response.Response
//	@Router		/api/v1/data/all [get]
func GetHistoryData(c *gin.Context) {
	histories, err := server.GetHistoryData()
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
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
		return
	}

	response.Success(c.Writer, "success", nil)
}

// DataDetailOpen
//
//	@Tags		历史数据模块
//	@Summary	打开查看历史数据进程
//	@Produce	application/json
//	@Param		id	formData	int	true	"历史数据id"
//	@Success	200	{object}	response.Response
//	@Router		/api/v1/data/do [post]
func DataDetailOpen(c *gin.Context) {
	id := c.PostForm("id")

	history, err := server.DataDetailOpen(id)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", history)
}

// GetDataDetailRanged
//
//	@Tags		历史数据模块
//	@Summary	获取范围历史数据,0,0可以查看全部数据
//	@Produce	application/json
//	@Param		id		query		int	true	"历史数据id"
//	@Param		from	query		int	true	"起始时间"
//	@Param		to		query		int	true	"截止时间"
//	@Success	200		{object}	response.Response
//	@Router		/api/v1/data/range [get]
func GetDataDetailRanged(c *gin.Context) {
	id := c.Query("id")
	from := c.Query("from")
	to := c.Query("to")

	rangeDataT, rangeDataP, err := server.GetRangeData(id, from, to)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
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
//	@Param		id	formData	int	true	"历史数据id"
//	@Success	200	{object}	response.Response
//	@Router		/api/v1/data/dc [post]
func DataDetailClose(c *gin.Context) {
	id := c.PostForm("id")

	err := server.DataDetailClose(id)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}
