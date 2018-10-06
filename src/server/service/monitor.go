package service

import (
	"blockclient/src/server/entity"
	"blockclient/src/server/common/stacode"
	"expvar"
	//"gopkg.in/kataras/iris.v5"
	"fmt"
	"gopkg.in/kataras/iris.v5"
)

type MonitorService struct {
}


// deprecated
func (ms *MonitorService) RunningStats() *entity.Result {

	res := make(map[string]interface{},  0)
	report := func(key string, value interface{}) {
		fmt.Println("k:=", key, "\nval:=", value)
		res[key] = value
	}
	// Do 方法会遍历底层注册表的 varKeys 和 vars
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	fmt.Println("res", res)
	return errcode.SUCCESS.Result(res)
}

// Use IO stream
func (ms *MonitorService) RunningStats2(ctx *iris.Context) {
	ctx.Response.Header.Set("Content-Type", "application/json; charset=utf-8")
	first := true
	report := func(key string, value interface{}) {
		if !first {
			fmt.Fprintf(ctx.Response.BodyWriter(), ",\n")
		}
		first = false
		if str, ok := value.(string); ok {
			fmt.Fprintf(ctx.Response.BodyWriter(), "%q: %q", key, str)
		} else {
			fmt.Fprintf(ctx.Response.BodyWriter(), "%q: %v", key, value)
		}
	}

	fmt.Fprintf(ctx.Response.BodyWriter(), "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	fmt.Fprintf(ctx.Response.BodyWriter(), "\n}\n")
}
