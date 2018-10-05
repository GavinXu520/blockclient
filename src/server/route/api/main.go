package api

import (
	"blockclient/src/server/controller"
	"gopkg.in/kataras/iris.v5"
	"log"
	"encoding/json"
)

var (
	ethCon = &controller.EthController{}
)
func Api() {

	iris.Get("/", func(ctx *iris.Context){
		ctx.Response.Header.Set("X-Frame-Options", "Allow-From https://www.ethereum.org/")
		//ctx.Render("index.html", map[string]interface{}{"title": "明星钻", "body": "hello"}, iris.RenderOptions{"gzip": true})
	})
	api := iris.Party("/api", func(ctx *iris.Context){
		ctx.Next()
		var res, method string
		method = string(ctx.Method())
		if "GET" == method {
			resByte, _ := json.Marshal(ctx.URLParams())
			res = string(resByte)
		}else{
			res = string(ctx.PostBody()[:])
		}
		//接口入参返参写入日志
		log.Println("Method", method, "\nResbody", res, "\nRespbody", string(ctx.Response.Body()))
	})

	// 发布合约
	api.Post("/v1/:appId/deploy",ethCon.DeployContract)

	// 查询账户的ether余额
	api.Get("/v1/:appId/getBlance", ethCon.GetBlanceByAddress)

	// 调用合约的transfer
	api.Post("/v1/:appId/tokentransfer", ethCon.TokenTransfer)

	// 查询账户的token余额
	api.Get("/v1/:appId/tokenBalance", ethCon.QueryTokenBalance)
}
