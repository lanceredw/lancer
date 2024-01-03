package network

import (
	"fmt"
	"net"
	"strings"
)

func GetLocalIP() (ipv4 string) {
	var (
		address []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
		err     error
	)
	// Get all network cards
	if address, err = net.InterfaceAddrs(); err != nil {
		fmt.Println("InterfaceAddress err:", err)
		return
	}
	// Take the first non lo network card IP address
	for _, addr = range address {
		//  ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// skip IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				// check "169.256" start
				if strings.HasPrefix(ipv4, "169.254") {
					continue // skip 169.254 start
				}
				return
			}
		}
	}

	return
}
