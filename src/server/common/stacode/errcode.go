package errcode

import "blockclient/src/server/entity"

type ErrCode struct {
	Code int
	Msg  string
}

func (e *ErrCode)Result(data interface{}) *entity.Result {
	return &entity.Result{
		Status:e.Code,
		Msg:e.Msg,
		Data:data,
	}
}



func (e *ErrCode)ResultWithMsg(msg string) *entity.Result {
	return &entity.Result{
		Status:e.Code,
		Msg:msg,
		Data:"",
	}
}

func  (e *ErrCode)ReplaceMsg(msg string) *ErrCode {
	return &ErrCode{
		Code:e.Code,
		Msg:msg,
	}
}

var (
	SUCCESS = &ErrCode{0, "success"}

	SYSTEMBUSY_ERROR = &ErrCode{-1,"系统繁忙"}

	REQUEST_PARAM_ERR = &ErrCode{1005,"请求参数非法"}

	LIMIT_REQUEST = &ErrCode{1014, "请求过于频繁"}

)
