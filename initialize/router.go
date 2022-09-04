package initialize

import (
	v1 "ginDemo/api/v1"
	_ "ginDemo/docs"
	"ginDemo/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetupRouter(r *gin.Engine) {
	r.Use(cors.Default()) // 跨域

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/test", testGin)

	baseGroup := r.Group("/api/v1")
	{
		baseGroup.POST("/register", v1.Register)
		baseGroup.POST("/login", v1.Login)
	}

	userGroup := r.Group("/api/v1/user", middleware.AuthRequired())
	{
		userGroup.POST("/info", v1.GetUserInfo)
	}
}

func testGin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"success": true,
	})
}
