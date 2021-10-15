module github.com/pigfall/wtun-go

go 1.17

require (
	github.com/pigfall/tzzGoUtil v1.3.0
	golang.org/x/sys v0.0.0-20211013075003-97ac67df715c
)

require (
	github.com/google/gopacket v1.1.17 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/vishvananda/netlink v1.1.0 // indirect
	github.com/vishvananda/netns v0.0.0-20191106174202-0a2b9b5464df // indirect
)

replace github.com/pigfall/tzzGoUtil => ../tzzGoUtil

replace golang.zx2c4.com/wireguard => ../wintun-go
