package util

import (
    "net"
)

func GetOutboundIP() (ip net.IP, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err == nil {
		defer conn.Close()
		localAddr := conn.LocalAddr().(*net.UDPAddr)
		ip = localAddr.IP
	}
	return
}
