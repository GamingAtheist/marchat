package main

import (
	"fmt"
	"net"
	"os"
)

func parseAddr(addr string) *net.IP {
	ipAddr, _, err := net.ParseCIDR(addr)
	if err != nil {
		ipAddr = net.ParseIP(addr)
	}
	if ipAddr == nil {
		return nil
	}
	return &ipAddr
}

func selectInterface() *net.UDPAddr {
	var NetInterface *net.Interface
	interfaceList, err := net.Interfaces()
	if err != nil {
		fmt.Println("[!] couldn't load interface list: ", err.Error())
		os.Exit(1)
	}

	for _, ifi := range interfaceList {
		addrList, err := ifi.Addrs()
		if err != nil {
			fmt.Println("[!] couldn't load interface list: ",
				err.Error())
			os.Exit(1)
		}
		for _, addr := range addrList {
			ip := parseAddr(addr.String())
			if !ip.IsLoopback() {
				NetInterface = &ifi
				break
			}
		}
		if NetInterface != nil {
			break
		}
	}

	if NetInterface == nil {
		fmt.Println("[!] couldn't find a valid interface")
		os.Exit(1)
	}

	chatSvc := fmt.Sprintf("239.255.255.250:%d", chatPort)
	gaddr, err := net.ResolveUDPAddr("udp", chatSvc)

	if err != nil {
		fmt.Println("[!] couldn't resolve multicast address: ", err.Error())
		os.Exit(1)
	}

	return gaddr
}
