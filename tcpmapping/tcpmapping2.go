package main

import (
	"fmt"
	"log"
	"net"
)

func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var localHost = "0.0.0.0"
var port = 28001

//var remoteHost = ""
//var remotePort = ""

func main() {
	//创建一个TCP服务端
	tcpaddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", localHost, port))
	chkError(err)
	//监听端口
	listen, err2 := net.ListenTCP("tcp", tcpaddr)
	chkError(err2)
	fmt.Printf("Server Start at %s:%d \n", localHost, port)
	//.建立链接并处理
	go func() {
		for {
			//如果有客户端链接过来，阻塞会返回
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			fmt.Println("收到客户端连接")

			//已经与客户端建立链接，处理业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						return
					}
					//回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write bak buf err", err)
						continue
					}
				}
			}()
		}
	}()
	//阻塞状态
	select {}
}
