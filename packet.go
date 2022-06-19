//go:build windows
// +build windows

package divert

import (
	"fmt"
	"net"

	"github.com/tboerc/divert-go/header"
)

type Packet struct {
	Raw   []byte
	Addr  *Address
	IpHdr header.IPHeader
}

func (p *Packet) String() string {
	return fmt.Sprintf("Packet {\n"+
		"\tIPHeader=%v\n"+
		"\tWinDivertAddr=%v\n"+
		"\tRawData=%v\n"+
		"}",
		p.IpHdr, p.Addr, p.Raw)
}

// Shortcut for IpHdr.SrcIP()
func (p *Packet) SrcIP() net.IP {
	return p.IpHdr.SrcIP()
}

// Shortcut for IpHdr.SetSrcIP()
func (p *Packet) SetSrcIP(ip net.IP) {
	p.IpHdr.SetSrcIP(ip)
}

// Shortcut for IpHdr.DstIP()
func (p *Packet) DstIP() net.IP {
	return p.IpHdr.DstIP()
}

// Shortcut for IpHdr.SetDstIP()
func (p *Packet) SetDstIP(ip net.IP) {
	p.IpHdr.SetDstIP(ip)
}

func NewPacket(buff []byte, addr *Address) *Packet {
	p := &Packet{Raw: buff, Addr: addr}

	version := int(buff[0] >> 4)

	if version == 4 {
		p.IpHdr = header.NewIPv4Header(buff)
	} else {
		p.IpHdr = header.NewIPv6Header(buff)
	}

	return p
}
