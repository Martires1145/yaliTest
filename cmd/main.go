package main

import (
	_ "cmdTest/config"
	_ "cmdTest/internal/database"
	"cmdTest/internal/route"
	"github.com/spf13/viper"
)

var port = viper.GetString("web.port")

func main() {
	_ = route.GetGin().Run(port)
}
