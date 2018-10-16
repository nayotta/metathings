package net_helper

import "net"

var InterfaceAddrs = net.InterfaceAddrs

// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go/31551220#31551220
func GetLocalIP() string {
	addrs, err := InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !ipnet.IP.IsLinkLocalUnicast() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
