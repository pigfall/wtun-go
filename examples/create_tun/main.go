package main

import(
	"unsafe"
	"syscall"
	"fmt"
	"golang.org/x/sys/windows"
)

func main() {
	dll,err :=syscall.LoadDLL("wintun.dll")
	if err != nil{
		panic(err)
	}
	proc,err := dll.FindProc("WintunCreateAdapter")
	if err != nil{
		panic(err)
	}
	deleteDriver ,err := dll.FindProc("WintunDeleteDriver")
	if err != nil{
		panic(err)
	}
	r0,_,err := deleteDriver.Call()
	if r0 == 0{
		panic(err)
	}
	fmt.Println("deleted driver")


	fmt.Println(proc)
	name,err := windows.UTF16PtrFromString("devName")
	if err != nil{
		panic(err)
	}
	tpe ,err := windows.UTF16PtrFromString("wintunType")
	if err != nil{
		panic(err)
	}
	r0,_,err = proc.Call(uintptr(unsafe.Pointer(name)),uintptr(unsafe.Pointer(tpe)))
	if r0 == 0{
		panic(err)
	}

	procOpenAdaptor,err := dll.FindProc("WintunOpenAdapter")
	if err != nil{
		panic(err)
	}
	fmt.Println(procOpenAdaptor)
	invalidName,err := windows.UTF16PtrFromString("noThisDevice")
	if err != nil{
		panic(err)
	}
	fmt.Println("call")
	r0,_,err = procOpenAdaptor.Call(uintptr(unsafe.Pointer(name)))
	if r0 == 0 {
		panic(err)
	}else{
		fmt.Println("open adaptor suc")
	}
	r0,_,err = procOpenAdaptor.Call(uintptr(unsafe.Pointer(invalidName)))
	if r0 == 0 {
		panic(err)
	}
	fmt.Println("over")
}
