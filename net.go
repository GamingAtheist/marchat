package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

var loopback = regexp.MustCompile("^lo")

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

func selectInterface() (*net.UDPAddr, *net.Interface) {
	var netInterface *net.Interface
	interfaceList, err := net.Interfaces()
	if err != nil {
		fmt.Println("[!] couldn't load interface list: ", err.Error())
		os.Exit(1)
	}

	for _, ifi := range interfaceList {
		if loopback.MatchString(ifi.Name) {
			continue
		}
		addrList, err := ifi.Addrs()
		if err != nil {
			fmt.Println("[!] couldn't load interface list: ",
				err.Error())
			os.Exit(1)
		}
		for _, addr := range addrList {
			ip := parseAddr(addr.String())
			if !ip.IsLoopback() {
				netInterface = &ifi
				break
			}
		}
		if netInterface != nil {
			break
		}
	}

	if netInterface == nil {
		fmt.Println("[!] couldn't find a valid interface")
		os.Exit(1)
	}

	chatSvc := fmt.Sprintf("239.255.255.250:%d", chatPort)
	gaddr, err := net.ResolveUDPAddr("udp", chatSvc)

	if err != nil {
		fmt.Println("[!] couldn't resolve multicast address: ", err.Error())
		os.Exit(1)
	}

	return gaddr, netInterface
}
