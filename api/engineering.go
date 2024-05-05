package api

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/response"
	"cmdTest/internal/server"
	"github.com/gin-gonic/gin"
)

// NewEngineering
//
//	@Summary	新增施工信息
//	@Tags		Engineering
//	@Param		ed	body	model.EngineeringJson	true	"施工信息"
//	@Router		/api/v1/en/new [post]
func NewEngineering(c *gin.Context) {
	var engineeringData model.EngineeringJson
	err := c.ShouldBindJSON(&engineeringData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 400)
		return
	}

	err = server.NewEngineering(&engineeringData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// ReviseEngineering
//
//	@Summary	修改施工信息
//	@Tags		Engineering
//	@Param		ed	body	model.EngineeringJson	true	"施工信息"
//	@Router		/api/v1/en/re [post]
func ReviseEngineering(c *gin.Context) {
	var engineeringData model.EngineeringJson
	err := c.ShouldBindJSON(&engineeringData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 400)
		return
	}

	err = server.ReviseEngineering(&engineeringData)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// AddDevices
//
//	@Summary	增加施工设备
//	@Tags		Engineering
//	@Param		dd	body	[]model.Device	true	"施工设备"
//	@Router		/api/v1/en/device/add [post]
func AddDevices(c *gin.Context) {
	var devices []model.Device
	err := c.ShouldBind(&devices)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 400)
		return
	}

	err = server.AddDevices(devices)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// DeleteDevices
//
//	@Summary	删除施工设备
//	@Tags		Engineering
//	@Param		deviceIDs	body	[]string	true	"要删除的设备ID"
//	@Router		/api/v1/en/device/delete [post]
func DeleteDevices(c *gin.Context) {
	var IDs []string
	err := c.ShouldBind(&IDs)
	if err != nil {
		response.Fail(c.Writer, "wrong data format", 400)
	}

	err = server.DeleteDevices(IDs)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// DeleteEngineering
//
//	@Summary	删除施工
//	@Tags		Engineering
//	@Param		engineeringID	formData	string	true	"要删除的施工ID"
//	@Router		/api/v1/en/delete [post]
func DeleteEngineering(c *gin.Context) {
	id := c.PostForm("engineeringID")

	err := server.DeleteEngineering(id)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// GetBriefEngineeringInfos
//
//	@Summary	查看全部施工的简述信息
//	@Tags		Engineering
//	@Router		/api/v1/en/all [get]
func GetBriefEngineeringInfos(c *gin.Context) {
	briefEngineeringInfos, err := server.GetBriefEngineeringInfos()
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", briefEngineeringInfos)
}

// GetEngineeringInfo
//
//	@Summary	查看施工详细信息
//	@Tags		Engineering
//	@Param		id	path	int	true	"施工ID"
//	@Router		/api/v1/en/{id} [get]
func GetEngineeringInfo(c *gin.Context) {
	id := c.Param("id")

	engineeringInfo, err := server.GetEngineeringInfo(id)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", *engineeringInfo)
}
