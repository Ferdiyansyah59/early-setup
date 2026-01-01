package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func NewRouter(ginMode string) *gin.Engine{
	gin.SetMode(ginMode)
	r := gin.Default()

	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	rootRouting := r.Group("/api") 
	{
		rootRouting.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "Halo Dunia Login")
		})
	}

	return r
}
