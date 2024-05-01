package route

import (
	"cmdTest/api"
	_ "cmdTest/docs"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetGin() *gin.Engine {
	r := gin.Default()

	r.Use(api.Cors)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/api/v1/rs", api.RunScript)
	r.GET("/api/v1/path", api.File)
	r.GET("/api/v1/dchan", api.DataChan)

	user := r.Group("/api/v1/user")
	{
		user.POST("/new", api.Register)
		user.POST("/v", api.Verify)
		user.POST("/login", api.Login)
		user.POST("/ru", api.UserRevise)
		user.POST("/d", api.DeleteUser)
		user.POST("/rp", api.ReSetUser)
		user.GET("/cu", api.CheckUserName)
		user.GET("/info", api.UserInfo)
		user.GET("/all", api.GetAllUser)
	}

	well := r.Group("/api/v1/well")
	{
		well.POST("/new", api.NewWell)
		well.POST("/rw", api.ReviseWell)
		well.POST("/d", api.DeleteWell)
		well.GET("/all", api.GetBriefWellInfo)
		well.GET("/:id", api.GetWellInfo)
	}

	engineering := r.Group("/api/v1/en")
	{
		engineering.POST("/new", api.NewEngineering)
		engineering.POST("/re", api.ReviseEngineering)
		engineering.POST("/device/add", api.AddDevices)
		engineering.POST("/device/delete", api.DeleteDevices)
		engineering.POST("/delete", api.DeleteEngineering)
		engineering.GET("/all", api.GetBriefEngineeringInfos)
		engineering.GET("/:id", api.GetEngineeringInfo)
	}

	model := r.Group("/api/v1/md")
	{
		model.POST("/new", api.NewModel)
		model.POST("/delete", api.DeleteModel)
		model.POST("/revise", api.ModifyModel)
		model.POST("/copy", api.CopyModel)
		model.GET("/all", api.GetModel)
		model.GET("/params", api.GetModelParams)
	}

	file := r.Group("/api/v1/file")
	{
		file.POST("/new", api.NewFilePath)
		file.POST("/dp", api.DeletePath)
		file.POST("/df", api.DeleteFile)
		file.POST("/csv", limits.RequestSizeLimiter(2<<20), api.CsvUpload)
		file.GET("/all", api.GetAllPath)
		file.GET("/", api.File)
	}

	history := r.Group("/api/v1/history")
	{
		history.POST("/uf", api.NewForecastData)
		history.POST("/delete", api.DeleteHistoryData)
		history.POST("/do", api.DataDetailOpen)
		history.POST("/dc", api.DataDetailClose)
		history.GET("/all", api.GetHistoryData)
		history.GET("/range", api.GetDataDetailRanged)
	}
	return r
}
