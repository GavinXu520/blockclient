package controller

import (
	"gopkg.in/kataras/iris.v5"
	"blockclient/src/server/service"
	"fmt"
	"encoding/json"
)

type MonitorController struct {
}

var (
	mservice = &service.MonitorService{}
)

// deprecated
func (mc *MonitorController) RunningStats (ctx *iris.Context) {
	//ctx.JSON(iris.StatusOK, mservice.RunningStats(ctx))
	res := mservice.RunningStats()
	//ctx.JSON(iris.StatusOK, *res)
	bt, err := json.Marshal(*(res))
	if nil != err {
		fmt.Println("err:=" + err.Error())
	}
	fmt.Println("str:=", string(bt))
	ctx.WriteString(string(bt))
}

// Real Running stats
func (mc *MonitorController) RunningStats2(ctx *iris.Context){
	mservice.RunningStats2(ctx)
}
