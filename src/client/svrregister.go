package client

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"reflect"
)

type SvrRegisterObj struct {
	Service     string
	ServiceInfo []SvrInfo
}

type SvrInfo struct {
	Path     string
	FuncName interface{}
}

func RegisterAndListen(registerObjs []SvrRegisterObj) {
	router := fasthttprouter.New()

	for _, obj := range registerObjs {
		fmt.Printf("Service %s Start\n", obj.Service)
		register(router, obj)
	}

	if err := fasthttp.ListenAndServe("localhost:8090", router.Handler); err != nil {
		fmt.Printf("出错: %v", err)
	}

}

func register(router *fasthttprouter.Router, obj SvrRegisterObj) {
	for _, v := range obj.ServiceInfo {
		path := "/" + v.Path
		generateFunc(router, path, v.FuncName)
	}
}

func generateFunc(r *fasthttprouter.Router, funcPath string, funcName interface{}) {
	h := reflect.ValueOf(funcName)
	t := h.Type()
	handleFunc := func(c *fasthttp.RequestCtx) {
		reqT := t.In(0).Elem()
		reqV := reflect.New(reqT)
		fmt.Println(reqV)
	}
	r.POST(funcPath, handleFunc)
	r.GET(funcPath, handleFunc)
}
