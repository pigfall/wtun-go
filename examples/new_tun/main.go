package main

import (
	"time"
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
	time.Sleep(time.Second*100)
}
