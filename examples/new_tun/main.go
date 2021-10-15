package main

import (
		"github.com/pigfall/tzzGoUtil/net"
	"fmt"
	wingo "github.com/pigfall/wtun-go"
)		

func main() {
	err := wingo.InitWinTun("wintun.dll")
	if err !=nil{
		panic(err)
	}
	tun,err := wingo.NewTun("testDevName")
	if err != nil{
		panic(err)
	}
	fmt.Println(tun)
	ipToSet,err := net.FromIpSlashMask("10.7.0.1/8")
	if err != nil{
		panic(err)
	}
	err = tun.SetIp(ipToSet)
	if err != nil{
		panic(err)
	}
	var buf = make([]byte,1024*4)
	for{
		n,err := tun.Read(buf)
		if err != nil{
			panic(err)
		}
		fmt.Println(buf[:n])
	}
}
