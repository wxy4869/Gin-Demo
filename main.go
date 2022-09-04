package main

import (
	"ginDemo/global"
	"ginDemo/initialize"
	"github.com/gin-gonic/gin"
)

// @title ginDemo
// @version 1.0
// @description MeowMeow
// @host localhost:8889
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	initialize.InitViper()

	initialize.InitMySQL()
	defer initialize.Close()

	r := gin.Default()
	initialize.SetupRouter(r)
	if err := r.Run(global.VP.GetString("host") + ":" + global.VP.GetString("port")); err != nil {
		panic(err)
	}
}
