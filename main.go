package main

import (
	"github.com/gin-gonic/gin"
	goserver "github.com/ikaiguang/propose_protobuf/go"
	"log"
	"net/http"
)

func main() {
	// gin
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	// html
	engine.LoadHTMLFiles("./js/index.html")
	engine.StaticFile("/favicon.ico", "./js/favicon.ico")
	engine.StaticFS("/dist", http.Dir("./js/dist"))
	//engine.StaticFS("/js", gin.Dir("./js", true))
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// api
	engine.Handle(http.MethodPost, goserver.SayHelloPath, goserver.SayHello)

	// run
	log.Println("server add : ", goserver.ServerAddress)
	if err := engine.Run(goserver.ServerAddress); err != nil {
		panic(err)
	}
}
