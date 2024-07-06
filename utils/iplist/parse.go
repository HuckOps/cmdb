package iplist

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func ParseListFromString(source string) ([]string, error) {
	var result []string
	lines := strings.Split(source, "\n")

	ipsList := [][]string{}

	for _, line := range lines {
		ips := strings.Split(line, ",")
		// ip有效性校验
		for _, ip := range ips {
			i := net.ParseIP(ip)
			if i == nil {
				return nil, errors.New(fmt.Sprintf("%s is not a vaild ip", ip))
			}
		}
		ipsList = append(ipsList, ips)
	}
	return result, nil
}
