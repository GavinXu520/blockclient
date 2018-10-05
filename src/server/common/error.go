package common

import (
	"fmt"
	"gopkg.in/kataras/iris.v5"
	"github.com/go-errors/errors"
	"blockclient/src/server/common/stacode"
)
func SetErrorDeal() {
	iris.Use(iris.HandlerFunc(func(ctx *iris.Context) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("发生panic异常: %v\n", errors.Wrap(err, 2).ErrorStack())
				ctx.JSON(iris.StatusOK, errcode.SYSTEMBUSY_ERROR.ResultWithMsg(msg))
			}
		}()
		ctx.Next()
	}))

}
