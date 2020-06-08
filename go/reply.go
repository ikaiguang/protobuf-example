package goserver

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	pb "github.com/ikaiguang/propose_protobuf/go/proto"
	"net/http"
)

// replyPb protobuf
func replyPb(ctx *gin.Context, data proto.Message) {
	resp := &pb.Response{
		Code:    http.StatusOK,
		Message: "success",
	}

	// any
	dataAny, err := ptypes.MarshalAny(data)
	if err != nil {
		replyPbError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Data = dataAny

	// response
	b, err := proto.Marshal(resp)
	if err != nil {
		replyPbError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Data(http.StatusOK, binding.MIMEPROTOBUF, b)
}

// replyPbError error pb
func replyPbError(ctx *gin.Context, code int32, msg string) {
	resp := &pb.Response{
		Code:    code,
		Message: msg,
	}
	b, _ := proto.Marshal(resp)
	//if err != nil {
	//	replyPbError(ctx, http.StatusInternalServerError, err.Error())
	//	return
	//}
	ctx.Data(http.StatusOK, binding.MIMEPROTOBUF, b)
}

// replyJSON json
func replyJSON(ctx *gin.Context, data proto.Message) {
	resp := &pb.Response{
		Code:    http.StatusOK,
		Message: "success",
	}
	// any
	dataAny, err := ptypes.MarshalAny(data)
	if err != nil {
		replyJSONError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	resp.Data = dataAny

	// marshal json
	var buf bytes.Buffer
	jsonHandler := jsonpb.Marshaler{OrigName: true, EmitDefaults: true}
	if err := jsonHandler.Marshal(&buf, resp); err != nil {
		replyJSONError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	defer buf.Reset()
	ctx.Data(http.StatusOK, binding.MIMEJSON, buf.Bytes())
}

// replyJSONError error json
func replyJSONError(ctx *gin.Context, code int32, msg string) {
	resp := &pb.Response{
		Code:    code,
		Message: msg,
	}
	ctx.JSON(http.StatusOK, resp)
}
