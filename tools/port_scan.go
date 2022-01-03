package main

import (
	"fmt"
	"net"
	"time"
)

type Host struct {
	Ip          string
	Port        int
	IsReachable bool
	msg         string
}

type HostScanner struct {
	hosts   []Host      // 待扫描资产列表
	timeout int         // 超时时间，单位： 毫秒
	msgCh   chan string // 用于向外传递消息的channel
}

type HostScan interface {
	Scan()
}

// Scan 端口扫描的主要执行逻辑
func (hs *HostScanner) Scan() {
	fmt.Println("开始扫描....")
	for _, host := range hs.hosts {
		address := fmt.Sprintf("%s:%d", host.Ip, host.Port)
		if _, err := net.DialTimeout("tcp", address, time.Duration(hs.timeout)*time.Millisecond); err != nil {
			fmt.Printf("%s 端口不通！", address)
		}
	}

}

// readConf 读取文本文件获取要扫描的主机端口
func readConf() (hosts []Host) {

	hosts = []Host{
		{"127.0.0.1", 2222, false, ""},
		{"127.0.0.1", 23339, false, ""},
		{"127.0.0.1", 16379, false, ""},
		{"127.0.0.1", 3389, false, ""},
		{"127.0.0.1", 8080, false, ""},
	}

	return
}

func main() {
	hosts := readConf()

	ch := make(chan string)

	var scanner HostScan = &HostScanner{hosts: hosts, timeout: 2000, msgCh: ch}
	scanner.Scan()

}
