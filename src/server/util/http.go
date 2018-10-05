package util

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	"encoding/json"
	"log"
)

func HttpGet(url string, params interface{}) (string, error) {
	paramByte, err := json.Marshal(params)
	if nil != err {
		return "", errors.New("HttpGet请求第三方网络入参错误err:=" + err.Error())
	}
	log.Println("HttpGet请求链接地址:", url, ",HttpGet请求参数为:", string(paramByte))
	request := gorequest.New()
	resp, body, errs := request.Get(url).Query(params).End()
	for _, err := range errs {
		if err != nil {
			log.Println("HttpGet请求链接地址:", url, ",HttpGet请求参数为:", string(paramByte), "HttpGet请求出错: err=", errs)
			return "", errors.New("HttpGet请求第三方网络发生错误")
		}
	}
	log.Println("HttpGet请求链接地址:", resp.Request.URL, ",HttpGet请求参数为:", string(paramByte), "HttpGet返回的结果是: ", body)
	return body, nil
}

func HttpPostForm(url string, params interface{}) (string, error) {
	paramByte, err := json.Marshal(params)
	if nil != err {
		return "", errors.New("HttpPostForm请求第三方网络入参错误err:=" + err.Error())
	}
	log.Println("HttpPostForm请求链接地址:", url, ",HttpPostForm请求参数为:", string(paramByte))
	request := gorequest.New()
	resp, body, errs := request.Post(url).Type("multipart").Send(params).End()
	for _, err := range errs {
		if err != nil {
			log.Println("HttpPostForm请求链接地址:", url, ",HttpPostForm请求参数为:", string(paramByte), ",HttpPostForm请求出错: err=", errs)
			return "", errors.New("HttpPostForm请求第三方网络发生错误")
		}
	}
	log.Println("HttpPostForm请求链接地址:", resp.Request.URL, ",HttpPostForm请求参数为:", string(paramByte), ",HttpPostForm返回的结果是: ", body)
	return body, nil
}

func PostForm(url string, params interface{}) (string, error) {
	paramByte, err := json.Marshal(params)
	if nil != err {
		return "", errors.New("postForm请求第三方网络入参错误err:=" + err.Error())
	}
	log.Println("postForm请求链接地址:", url, ",postForm请求参数为:", string(paramByte))
	request := gorequest.New()
	resp, body, errs := request.Post(url).Type("form").Send(params).End()
	for _, err := range errs {
		if err != nil {
			log.Println("postForm请求链接地址:", url, ",postForm请求参数为:", string(paramByte), ",postForm请求出错: err=", errs)
			return "", errors.New("postForm请求第三方网络发生错误")
		}
	}
	log.Println("postForm请求链接地址:", resp.Request.URL, ",postForm请求参数为:", string(paramByte), ",postForm返回的结果是: ", body)
	return body, nil
}

func HttpPostJson(url string, params interface{}) (string, error) {
	paramByte, err := json.Marshal(params)
	if nil != err {
		return "", errors.New("HttpPostJson请求第三方网络入参错误err:=" + err.Error())
	}
	log.Println("HttpPostJson请求链接地址:", url, ",HttpPostJson请求json参数为:", string(paramByte))
	request := gorequest.New()
	resp, body, errs := request.Post(url).Send(params).End()
	for _, err := range errs {
		if err != nil {
			log.Println("HttpPostJson请求链接地址:", url, ",HttpPostJson请求json参数为:", string(paramByte), ",HttpPostJson请求出错: err=", errs)
			return "", errors.New("HttpPostJson请求第三方网络发生错误")
		}
	}
	log.Println("HttpPostJson请求链接地址:", resp.Request.URL, ",HttpPostJson请求json参数为:", string(paramByte), ",HttpPostJson返回的结果是: ", body)
	return body, nil
}

//method大写POST
func HttpCustomMethod(method, url string, params interface{}) (string, error) {
	paramByte, err := json.Marshal(params)
	if nil != err {
		return "", errors.New("postJson请求第三方网络入参错误err:=" + err.Error())
	}
	log.Println("postJson请求链接地址:", url, ",postJson请求json参数为:", string(paramByte))
	request := gorequest.New()
	resp, body, errs := request.CustomMethod(method, url).Send(params).End()
	for _, err := range errs {
		if err != nil {
			log.Println("postJson请求链接地址:", url, ",postJson请求json参数为:", string(paramByte), ",postJson请求出错: err=", errs)
			return "", errors.New("postJson请求第三方网络发生错误")
		}
	}
	log.Println("postJson请求链接地址:", resp.Request.URL, ",postJson请求json参数为:", string(paramByte), ",postJson返回的结果是: ", body)
	return body, nil
}
