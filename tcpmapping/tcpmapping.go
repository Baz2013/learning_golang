package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var host = "0.0.0.0"
var port = 8000
var bufSize = 4096
var remoteServer = "127.0.0.1"
var remotePort = 7890

func main() {
	// Part 1: create a listener
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Error listener returned: %s", err)
	}
	defer l.Close()
	log.Printf("server listen at %s:%d", host, port)

	for {
		// Part 2: accept new connection
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("Error to accept new connection: %s", err)
		}

		// Part 3: create a goroutine that reads and write back data
		go HandleConn(c)
	}
}

func HandleConn(client net.Conn) {
	log.Printf("TCP session open")
	defer client.Close()

	clientCh := make(chan []byte)
	serverCh := make(chan []byte)

	remoteConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", remoteServer, remotePort))
	if err != nil {
		log.Fatalf("Error to open TCP connection: %s", err)
		return
	}

	defer remoteConn.Close()

	// 读取客户端数据，发送给远程主机
	go func() {
		for {
			d := make([]byte, bufSize)
			// Read from TCP buffer
			client.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
			n, err := client.Read(d)
			if err != nil {
				log.Printf("conn read %d bytes,  error: %s", n, err)
				if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
					continue
				}
				return
			}
			log.Printf("read %d bytes, content is %s\n", n, string(d[:n]))

			if n > 0 {
				clientCh <- d
			}

		}
	}()

	// 读取远程主机数据，发送给客户端
	go func() {
		for {
			d := make([]byte, bufSize)
			// Read from TCP buffer
			remoteConn.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
			n, err := remoteConn.Read(d)
			if err != nil {
				log.Printf("conn read %d bytes,  error: %s", n, err)
				if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
					continue
				}
				return
			}
			log.Printf("read %d bytes, content is %s\n", n, string(d[:n]))

			if n > 0 {
				serverCh <- d
			}
		}
	}()

	go func() {
		for {
			select {
			case c := <-clientCh:
				_, err := remoteConn.Write(c)
				if err != nil {
					return
				}
			case s := <-serverCh:
				_, err := client.Write(s)
				if err != nil {
					return
				}
			}
		}
	}()

}

// ConnectRemoteServer 连接远程主机z
func ConnectRemoteServer(host string, port int) (c *net.Conn, e error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Error to open TCP connection: %s", err)
		return &conn, err
	}
	log.Printf("已与远程主机 %s:%d 建立连接", host, port)
	return &conn, nil
}
