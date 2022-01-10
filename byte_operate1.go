package main

import "fmt"

type clientFlag uint32

// 1, 2, 4, 8, 16, 32, 64, 128, 256, 512 ...
const (
	clientLongPassword clientFlag = 1 << iota
	clientFoundRows
	clientLongFlag
	clientConnectWithDB
	clientNoSchema
	clientCompress
	clientODBC
	clientLocalFiles
	clientIgnoreSpace
	clientProtocol41
	clientInteractive
	clientSSL
	clientIgnoreSIGPIPE
	clientTransactions
	clientReserved
	clientSecureConn
	clientMultiStatements
	clientMultiResults
	clientPSMultiResults
	clientPluginAuth
	clientConnectAttrs
	clientPluginAuthLenEncClientData
	clientCanHandleExpiredPasswords
	clientSessionTrack
	clientDeprecateEOF
)

/**
通过 |、^、& 运算，我们可以很方便快捷的对状态值进行操作。当然，位运算的应用不仅限于状态值，知道了其中的二进制运算原理后，还有更多的其他应用场景，等着你去发现。
from https://juejin.cn/post/6844903909333401607
*/

func main() {
	fmt.Println(clientLongPassword)
	fmt.Println(clientFoundRows)
	fmt.Println(clientLongFlag)
	fmt.Println(clientConnectWithDB)
	fmt.Println(clientNoSchema)
	fmt.Println(clientCompress)
	fmt.Println(clientODBC)
	fmt.Println(clientLocalFiles)
	fmt.Println(clientIgnoreSpace)
	fmt.Println(clientProtocol41)
	status := clientLongPassword | clientLongFlag | clientNoSchema | clientODBC | clientProtocol41
	fmt.Println(status)
	fmt.Println("Include clientLongPassword? ", status&clientLongPassword)   // 1
	fmt.Println("Include clientFoundRows? ", status&clientFoundRows)         // 0
	fmt.Println("Include clientLongFlag? ", status&clientLongFlag)           // 4
	fmt.Println("Include clientConnectWithDB? ", status&clientConnectWithDB) // 0
	fmt.Println("Include clientConnectWithDB? ", status&clientNoSchema)      // 16
	fmt.Println("Include clientConnectWithDB? ", status&clientCompress)      // 0
}
