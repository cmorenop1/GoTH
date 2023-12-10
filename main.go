package main

import (
	"log"

	"github.com/GoTH/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	// App Custom made Renderer (renderer.go)
	app.HTMLRender = &TemplRender{}

	// App Handlers
	app.GET("/", handlers.Index)
	app.GET("/home", handlers.Home)

	// App listen
	log.Fatal(app.Run(":8080"))
}
