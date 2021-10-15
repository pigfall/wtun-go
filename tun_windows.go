package wtun

import(
	"errors"
		"unsafe"
		"github.com/pigfall/tzzGoUtil/syscall"
		"github.com/pigfall/tzzGoUtil/process"
		"fmt"
		"github.com/pigfall/tzzGoUtil/net"
		"golang.org/x/sys/windows"
)

var (
	wintunDLL *syscall.DLL
	procReadPacket *syscall.Procdure
	procReleaseRcvPacket *syscall.Procdure
	procAllocSendPacket *syscall.Procdure
	procSendPacket *syscall.Procdure
)


func InitWinTun(wintunDLLPath string)(error) {
	var err error
	wintunDLL,err = syscall.LoadDLL(wintunDLLPath)
	if err != nil{
		return err
	}
	procReadPacket,err = wintunDLL.FindProcure("WintunReceivePacket")
	if err != nil{
		return err
	}
	procReleaseRcvPacket, err = wintunDLL.FindProcure("WintunReleaseReceivePacket")
	if err != nil{
		return err
	}
	procAllocSendPacket,err = wintunDLL.FindProcure("WintunAllocateSendPacket")
	if err != nil{
		return err
	}
	procSendPacket,err = wintunDLL.FindProcure("WintunSendPacket")
	if err != nil{
		return err
	}

	return nil
}


type Tun struct{
	handle uintptr
	sessionHandle uintptr
	devName string
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
		devName:devName,
	},nil
}

func (this *Tun) Read(b []byte)(n int,err error){
	var packetSize int
	for{
		r0,_,err :=procReadPacket.Call(
			this.sessionHandle,
			uintptr(unsafe.Pointer(&packetSize)),
		)
		if err != nil{
			if errors.Is(err,windows.ERROR_NO_MORE_ITEMS){
				continue
			}
			return 0,err
		}
		packet := unsafe.Slice((*byte)(unsafe.Pointer(r0)), packetSize)
		n = copy(b,packet)
		procReadPacket.Call(
			this.sessionHandle,
			r0,
		)
		return n,nil
	}
}

func (this *Tun) Write(b []byte)(n int,err error){
	pacSize := len((b))
	bufAddr,_,err := procAllocSendPacket.Call(this.sessionHandle,uintptr(pacSize))
	if err != nil{
		return 0,err
	}
	bufToUse := unsafe.Slice((*byte)(unsafe.Pointer(bufAddr)),pacSize)
	_,_,err = procSendPacket.Call(this.sessionHandle,uintptr(unsafe.Pointer(&bufToUse[0])))
	if err != nil{
		return 0,err
	}

	return pacSize,nil
}
 
func (this *Tun) Name()(string,error) {
	return this.devName,nil
}

func (this *Tun) SetIp(ipNet *net.IpWithMask)error{
	devName := this.devName
	cmds :=[]string{
		"netsh","interface","ip","set","address",devName,"static",ipNet.Ip.String(),net.MaskFormatTo255(ipNet.Mask),
	}
	out,errOut,err := process.ExeOutput(cmds[0],cmds[1:]...)
	if err != nil{
		return fmt.Errorf("%w, %v, %v",err,errOut,out)
	}
	return nil
}

