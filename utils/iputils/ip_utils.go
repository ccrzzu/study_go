package iputils

import (
	"fmt"
	"math/big"
	"net"
	"strconv"
	"strings"
)

//ip 2 int
func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

//int 2 ip string
func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func UInt32toIP(intIP uint32) net.IP {
	var bytes [4]byte
	//0xFF == 255
	bytes[0] = byte(intIP & 0xFF)
	bytes[1] = byte((intIP >> 8) & 0xFF)
	bytes[2] = byte((intIP >> 16) & 0xFF)
	bytes[3] = byte((intIP >> 24) & 0xFF)
	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

func IPtoUInt32(ip net.IP) uint32 {
	strArr := strings.Split(ip.String(), ".")

	b0, _ := strconv.Atoi(strArr[0])
	b1, _ := strconv.Atoi(strArr[1])
	b2, _ := strconv.Atoi(strArr[2])
	b3, _ := strconv.Atoi(strArr[3])

	var sum uint32
	sum += uint32(b0) << 24
	sum += uint32(b1) << 16
	sum += uint32(b2) << 8
	sum += uint32(b3)

	return sum
}
