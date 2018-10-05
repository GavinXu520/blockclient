package util

import (
	"encoding/json"
	"fmt"
)

func Interface2Json(argsName string, args interface{}) {
	jsonByte, err := json.Marshal(args)
	if nil != err {
		fmt.Println("入参不能被解析成json, args:=" + fmt.Sprint(args))
	}
	fmt.Println(argsName + ":=" + string(jsonByte) + " \n")
}

/**
参数转换成string 用于日志
 */
func Interface2JsonStr(args interface{}) (string, error) {
	jsonByte, err := json.Marshal(args)
	if nil != err {
		return "", err
	}
	return string(jsonByte), nil
}