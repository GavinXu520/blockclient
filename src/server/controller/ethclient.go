package controller

import (
	"gopkg.in/kataras/iris.v5"
	"blockclient/src/server/service"
	"encoding/json"
	"blockclient/src/server/entity"
	"blockclient/src/server/common/stacode"
	"math/big"
)

type EthController struct {
}

var ethService = &service.EthService{}

// query account's ether balance
func (ec *EthController) GetBlanceByAddress(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, ethService.GetBalanceByAddress(ctx.URLParam("addresses"), ctx.URLParam("num")))
}

// deploy a contract
func (ec *EthController) DeployContract(ctx *iris.Context){
	var res entity.DeployEntity
	if err := json.Unmarshal(ctx.PostBody(), &res); nil != err {
		ctx.JSON(iris.StatusOK, errcode.REQUEST_PARAM_ERR.Result(err))
		return
	}
	key, err :=  json.Marshal(res.Key)
	if nil != err {
		ctx.JSON(iris.StatusOK, errcode.REQUEST_PARAM_ERR.Result("parse key err: " + err.Error()))
		return
	}
	ctx.JSON(iris.StatusOK, ethService.DeployContract(res.Total, res.Decimals, res.Name, res.Symbol, string(key), res.Pwd))
}

func (ec *EthController) TokenTransfer (ctx *iris.Context) {
	var res entity.ToekenTransferEntity
	if err := json.Unmarshal(ctx.PostBody(), &res); nil != err {
		ctx.JSON(iris.StatusOK, errcode.REQUEST_PARAM_ERR.Result(err))
		return
	}
	key, err :=  json.Marshal(res.Key)
	if nil != err {
		ctx.JSON(iris.StatusOK, errcode.REQUEST_PARAM_ERR.Result("parse key err: " + err.Error()))
		return
	}
	ctx.JSON(iris.StatusOK, ethService.TokenTransfer(res.Address, res.To, string(key), res.Pwd, big.NewInt(int64(res.Value))))
}

func (ec *EthController) QueryTokenBalance (ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, ethService.QueryTokenBalance(ctx.URLParam("contract"), ctx.URLParam("eoas")))
}
