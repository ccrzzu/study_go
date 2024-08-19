package test

import (
	"fmt"
	"net"
	"study_go/utils/crypto"
	"study_go/utils/iputils"
	"testing"
)

func TestIP(t *testing.T) {
	ori := "192.168.78.123"
	fmt.Println("origin:", ori)
	i := iputils.InetAtoN(ori)
	fmt.Println(i)
	/*fmt.Println(byte(i>>24))
	fmt.Println(byte(i>>16))
	fmt.Println(byte(i>>8))
	fmt.Println(byte(i))*/
	fmt.Println(iputils.InetNtoA(i))

	fmt.Println(iputils.IPtoUInt32(net.IP{12,13,14,15}))
	fmt.Println(net.IP{12,13,14,15}.String())
	fmt.Println(iputils.UInt32toIP(202182159))
}


func TestMD5(t *testing.T){
	fmt.Println(crypto.GetMD5UUID())
}
