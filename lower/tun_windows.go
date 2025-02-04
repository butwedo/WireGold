//go:build windows
// +build windows

package lower

import "net"

func (n *NIC) prepare() {
	_, ipn, err := net.ParseCIDR(n.subnet)
	if err != nil {
		panic(err)
	}
	execute("cmd", "/c", "netsh interface ip set address name=\""+n.ifce.Name()+"\" source=static addr=\""+n.ip+"\" mask=\""+(net.IP)(ipn.Mask).String()+"\" gateway=none")
}

func (n *NIC) Up() {
	// execute("netsh", "interface", "set", "interface", n.ifce.Name(), "enabled")
	// don't need to bring up the device by hand
}

func (n *NIC) Down() {
	// execute("netsh", "interface", "set", "interface", n.ifce.Name(), "disabled")
	// don't need to bring up the device by hand
}
