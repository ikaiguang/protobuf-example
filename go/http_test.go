package goserver

import (
	"bytes"
	"crypto/tls"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	pb "github.com/ikaiguang/propose_protobuf/go/proto"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestSayHello_Pb(t *testing.T) {
	// requset name
	var name = "Go"
	helloReq := &pb.HelloReq{Name: name}
	helloReqBytes, err := proto.Marshal(helloReq)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	httpReader := bytes.NewReader(helloReqBytes)
	t.Log("pb request size : ", httpReader.Size())

	// http request
	httpURL := "http://" + ServerAddress + SayHelloPath
	httpReq, err := http.NewRequest(http.MethodPost, httpURL, httpReader)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	httpReq.Header.Set("content-type", binding.MIMEPROTOBUF)

	// req
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer httpResp.Body.Close()

	// bad request
	if httpResp.StatusCode != http.StatusOK {
		t.Errorf("request fail : code: %d, msg :%s \n \t request url : %s", httpResp.StatusCode, httpResp.Status, httpURL)
		t.FailNow()
	}

	// read body
	bodyBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll(resp.Body) fail : %s", err.Error())
		t.FailNow()
	}
	//t.Log("pb response size : ", len(bodyBytes))
	bodyReader := bytes.NewReader(bodyBytes)
	t.Log("pb response size : ", bodyReader.Size())

	// response
	var response = &pb.Response{}
	err = proto.Unmarshal(bodyBytes, response)
	if err != nil {
		t.Errorf("jsonpb.Unmarshal(httpResp.Body, helloResp) error : %s", err.Error())
		t.FailNow()
	}
	if response.Code != http.StatusOK {
		t.Errorf("request result is fail : code: %d, msg :%s", response.Code, response.Message)
		t.FailNow()
	}

	// hello resp
	var helloResp = &pb.HelloResp{}
	err = ptypes.UnmarshalAny(response.Data, helloResp)
	if err != nil {
		t.Errorf("ptypes.UnmarshalAny(response.Data, helloResp) error : %s", err.Error())
		t.FailNow()
	}

	// result
	if helloResp.Message != sayHelloMsg(name) {
		t.Error("helloResp.Message != sayHelloMsg(name)")
		t.FailNow()
	}

	t.Logf("say_hello pb test success : %s \n", response.String())
}

func TestSayHello_JOSN(t *testing.T) {
	// requset name
	var name = "Go"
	httpReader := strings.NewReader(`{"name":"` + name + `"}`)
	t.Log("json request size : ", httpReader.Size())

	// http request
	httpURL := "http://" + ServerAddress + SayHelloPath
	httpReq, err := http.NewRequest(http.MethodPost, httpURL, httpReader)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// req
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer httpResp.Body.Close()

	// bad request
	if httpResp.StatusCode != http.StatusOK {
		t.Errorf("request fail : code: %d, msg :%s \n \t request url : %s", httpResp.StatusCode, httpResp.Status, httpURL)
		t.FailNow()
	}

	// read body
	// 此处为了计算传输大小，中间多读了一次
	bodyBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll(resp.Body) fail : %s", err.Error())
		t.FailNow()
	}
	bodyReader := bytes.NewReader(bodyBytes)
	t.Log("json response size : ", bodyReader.Size())

	// response
	var response = &pb.Response{}
	err = jsonpb.Unmarshal(bodyReader, response)
	//err = jsonpb.Unmarshal(httpResp.Body, response)
	if err != nil {
		t.Errorf("jsonpb.Unmarshal(httpResp.Body, helloResp) error : %s", err.Error())
		t.FailNow()
	}
	if response.Code != http.StatusOK {
		t.Errorf("request fail : code: %d, msg :%s", response.Code, response.Message)
		t.FailNow()
	}

	// hello resp
	var helloResp = &pb.HelloResp{}
	err = ptypes.UnmarshalAny(response.Data, helloResp)
	if err != nil {
		t.Errorf("ptypes.UnmarshalAny(response.Data, helloResp) error : %s", err.Error())
		t.FailNow()
	}

	// result
	if helloResp.Message != sayHelloMsg(name) {
		t.Error("helloResp.Message != sayHelloMsg(name)")
		t.FailNow()
	}

	t.Logf("say_hello josn test success : %s \n", response.String())
}
