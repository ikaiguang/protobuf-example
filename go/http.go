package goserver

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	pb "github.com/ikaiguang/propose_protobuf/go/proto"
	"io"
	"io/ioutil"
	"net/http"
)

// SayHelloPath http route
const (
	ServerAddress = "127.0.0.1:51002"
	SayHelloPath  = "/say_hello"
)

// SayHello say hello
/*
curl -k -X POST -H "Content-Type: application/json" -d "{\"name\":\"ZhangSan\"}" http://127.0.0.1:51002/say_hello

curl -k -X POST -H "Content-Type: application/json" -d "{\"name\":\"\u5f20\u4e09\"}" http://127.0.0.1:51002/say_hello
*/
func SayHello(ctx *gin.Context) {
	// pb & json
	switch ctx.GetHeader("content-type") {
	case binding.MIMEPROTOBUF:
		sayHelloPb(ctx)
	default:
		sayHello(ctx)
		return
	}
}

// sayHelloPb context type = pb
func sayHelloPb(ctx *gin.Context) {
	// bodyBytes
	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		replyJSONError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// unmarshal pb
	helloReq := &pb.HelloReq{}
	if err := proto.Unmarshal(bodyBytes, helloReq); err != nil {
		replyPbError(ctx, http.StatusInternalServerError, "proto.Unmarshal(bodyBytes, helloReq) error : "+err.Error())
		return
	}

	// resp
	if len(helloReq.Name) == 0 {
		helloReq.Name = "World"
	}
	helloResp := &pb.HelloResp{Message: sayHelloMsg(helloReq.Name)}

	replyPb(ctx, helloResp)
}

// sayHello normal context type
func sayHello(ctx *gin.Context) {
	// unmarshal json
	helloReq := &pb.HelloReq{}
	err := jsonpb.Unmarshal(ctx.Request.Body, helloReq)
	if err != nil {
		if err == io.EOF {
			err = nil
		} else {
			replyJSONError(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// resp
	if len(helloReq.Name) == 0 {
		helloReq.Name = "World"
	}
	helloResp := &pb.HelloResp{Message: sayHelloMsg(helloReq.Name)}

	replyJSON(ctx, helloResp)
}

// sayHelloMsg hello msg
func sayHelloMsg(name string) string {
	return "Hello " + name
}
