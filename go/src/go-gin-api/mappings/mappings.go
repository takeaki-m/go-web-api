package mappings

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine;

func CreateUrlMappings() {
	Router = gin.Default()
	Router.LoadHTMLGlob("templates/*.html")

	data := "Hello Go/Gin!!"

	Router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	Router.Run()
	/**
	api := Router.Group("/api") 
	{
		api.GET("/test", func(ctx *gin.Context){
			ctx.JSON(200, gin.H{
				"message": "test successful",
			});
		})
	}
	**/
}