package main

import "fmt"

/**
3 字节的小端序转大端序
*/
//func main() {
//	demoByteOperate1()
//}

func demoByteOperate1() {
	data := []byte{74, 00, 00, 00} // mysql 协议 header
	pktLen := int(uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16)
	/**
	uint32(data[0])     000000000000000001001010
	uint32(data[1])<<8  000000000000000000000000
	uint32(data[1])<<16 000000000000000000000000
	                  | 000000000000000001001010
	*/
	fmt.Println(pktLen)

	//https://blog.erratasec.com/2016/11/how-to-teach-endian.html#.YdgX82hByUk
	//https://www.ruanyifeng.com/blog/2016/11/byte-order.html

	//32位整数的求值公式
	/**
	大端字节序
	i = (data[3]<<0) | (data[2]<<8) | (data[1]<<16) | (data[0]<<24)

	小端字节序
	i = (data[0]<<0) | (data[1]<<8) | (data[2]<<16) | (data[3]<<24)
	*/

	/**
	mysql 协议中，包长为 int<3> 即 24 位整数，所以其求值公式为
	i = uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16
	*/
}
