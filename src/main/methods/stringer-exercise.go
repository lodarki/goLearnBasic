package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
//让 IPAddr 类型实现 fmt.Stringer 以便用点分格式输出地址。
//例如，`IPAddr{1,`2,`3,`4}` 应当输出 `"1.2.3.4"`。
func (ipAddr *IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a.String())
	}
}
