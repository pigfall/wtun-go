package wtun

import(
		"unsafe"
		"github.com/pigfall/tzzGoUtil/syscall"
		"golang.org/x/sys/windows"
)

var (
	wintunDLL *syscall.DLL
)


func InitWinTun(wintunDLLPath string)(error) {
	var err error
	wintunDLL,err = syscall.LoadDLL(wintunDLLPath)
	if err != nil{
		return err
	}
	return nil
}


type Tun struct{
	handle uintptr
	sessionHandle uintptr
}

func NewTun(devName string)(*Tun,error){
	proc,err := wintunDLL.FindProcure("WintunCreateAdapter")
	if err != nil{
		return nil,err
	}
	devNamePtr,err := windows.UTF16PtrFromString(devName)
	if err != nil{
		return nil,err
	}
	devHandle,_,err := proc.Call(uintptr(unsafe.Pointer(devNamePtr)))
	if err != nil{
		return nil,err
	}

	startSessionProc,err := wintunDLL.FindProcure("WintunStartSession")
	if err != nil {
		return nil,err
	}
	var capacity = 0x400000
	sessionHandle,_,err :=startSessionProc.Call(devHandle,uintptr(capacity))
	if err != nil{
		return nil,err
	}

	return &Tun{
		handle:devHandle,
		sessionHandle:sessionHandle,
	},nil
}

