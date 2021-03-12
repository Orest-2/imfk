package main

import (
	"mime"
	"net/http"

	"github.com/Orest-2/imfk/controllers"
	"github.com/Orest-2/imfk/mocks"
	"github.com/Orest-2/imfk/models"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// CORS Middleware
func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		c.AbortWithStatus(http.StatusOK)

	}
}

func main() {
	r := gin.Default()

	models.ConnectDataBase()
	mocks.Init()

	r.Use(CORS)
	r.Use(static.Serve("/", static.LocalFile("./clients/main/dist", true)))

	v1 := r.Group("api/v1")
	{
		// settings
		v1.GET("/settings", controllers.GetSettings)
		//mf
		v1.POST("/mf/2d/plot", controllers.Make2DPlot)
		v1.POST("/mf/3d/plot", controllers.Make3DPlot)
		//mf
		v1.POST("/mf/2d/eval", controllers.Eval2D)
	}

	mime.AddExtensionType(".js", "text/javascript")

	r.Run(":1447")
}
