package util

import (
	"strings"
	"strconv"
	"github.com/go-errors/errors"
	"blockclient/src/server/common"
)

/**
字符串 IP 转 int
 */
func Inet_aton(ip string) (int, error) {
	ips := strings.Split(ip, ".")
	E := errors.New("Not A IP.")
	if len(ips) != 4 {
		return 0, E
	}
	var intIP int
	for k, v := range ips {
		i, err := strconv.Atoi(v)
		if err != nil || i > 255 {
			return 0, E
		}
		intIP = intIP | i << uint(8 * (3 - k))
	}
	log.Println("ip:", ip, ",经解析后的int类型IP:",intIP)
	return intIP, nil
}


