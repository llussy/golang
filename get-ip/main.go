package main

import (
	"fmt"
	"net"
	"regexp"
)

// REGLOCALIP ip
const REGLOCALIP = `^((172\.([1][6-9]|[2]\d|3[01]))(\.([2][0-4]\d|[2][5][0-5]|[01]?\d?\d)){2}|10(\.([2][0-4]\d|[2][5][0-5]|[01]?\d?\d)){3})|192\.168(\.([2][0-4]\d|[2][5][0-5]|[01]?\d?\d)){2}$`

func main() {
	fmt.Println(GetLocalIP())
}

// GetLocalIP func
func GetLocalIP() (string, error) {
	var res = "0.0.0.0"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return res, err
	}
	for _, a := range addrs {
		ipnet, ok := a.(*net.IPNet)
		matchedIP, err := regexp.MatchString(REGLOCALIP, ipnet.IP.String())
		if err != nil {
			return res, err
		}

		if ok && !ipnet.IP.IsLoopback() && matchedIP {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return res, fmt.Errorf("not found local IP")
}
