package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

/*
md5加密
参数：需要加密的字符串
返回：加密后的字符串
*/
func Md5(data string) string {
	hash := md5.New()          //初始化一个MD5对象
	hash.Write([]byte(data))   // 需要加密的字符串
	cipherStr := hash.Sum(nil) //计算出校验和
	fmt.Println("md5前的信息: ", data)
	fmt.Println("md5后的信息: ", hex.EncodeToString(cipherStr))
	return hex.EncodeToString(cipherStr) // 输出加密结果
}

func Base64(data string) string {
	fmt.Println("base64前的信息: ", data)
	bytes := []byte(data)
	result := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println("base64后的信息: ", result)
	return result
}
