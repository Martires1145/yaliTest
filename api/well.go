package api

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/response"
	"cmdTest/internal/server"
	"github.com/gin-gonic/gin"
)

// NewWell
//
//	@Summary	新增井信息
//	@Tags		Well
//	@Param		wd	body	model.WellJson	true	"井信息"
//	@Router		/api/v1/well/new [post]
func NewWell(c *gin.Context) {
	var wellData model.WellJson
	err := c.ShouldBindJSON(&wellData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 400)
		return
	}

	err = server.NewWell(&wellData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// ReviseWell
//
//	@Summary	编辑井信息
//	@Tags		Well
//	@Param		wd	body	model.WellJson	true	"井信息"
//	@Router		/api/v1/well/rw [post]
func ReviseWell(c *gin.Context) {
	var wellData model.WellJson
	err := c.ShouldBindJSON(&wellData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 400)
		return
	}

	err = server.ReviseWell(&wellData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// DeleteWell
//
//	@Summary	删除井信息
//	@Tags		Well
//	@Param		wellID	body	int	true	"井ID"
//	@Router		/api/v1/well/d [post]
func DeleteWell(c *gin.Context) {
	id := c.PostForm("wellID")

	err := server.DeleteWell(id)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// GetBriefWellInfo
//
//	@Summary	获取所有井信息
//	@Tags		Well
//	@Router		/api/v1/well/all [get]
func GetBriefWellInfo(c *gin.Context) {
	briefInfos, err := server.GetBriefWellInfo()
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", briefInfos)
}

// GetWellInfo
//
//	@Summary	获取井的详细信息
//	@Tags		Well
//	@Param		id	path	int	true	"井ID"
//	@Router		/api/v1/well/{id} [get]
func GetWellInfo(c *gin.Context) {
	id := c.Param("id")
	info, err := server.GetWellInfo(id)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", info)
}
