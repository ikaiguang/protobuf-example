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

	// route
	engine.Handle(http.MethodPost, goserver.SayHelloPath, goserver.SayHello)

	// run
	log.Println("server add : ", goserver.ServerAddress)
	if err := engine.Run(goserver.ServerAddress); err != nil {
		panic(err)
	}
}
