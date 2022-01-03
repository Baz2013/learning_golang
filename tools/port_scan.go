package main

import (
	"fmt"
	"net"
	"sync"
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
	Scan(wg *sync.WaitGroup)
}

// Scan 端口扫描的主要执行逻辑
func (hs *HostScanner) Scan(wg *sync.WaitGroup) {
	fmt.Println("开始扫描....")
	for _, host := range hs.hosts {
		address := fmt.Sprintf("%s:%d", host.Ip, host.Port)
		if _, err := net.DialTimeout("tcp", address, time.Duration(hs.timeout)*time.Millisecond); err != nil {
			msg := fmt.Sprintf("%s 端口不通！", address)
			//hs.msgCh <- msg
			fmt.Println(msg)
		}
	}

	wg.Done() // 任务结束
}

// readConf 读取文本文件获取要扫描的主机端口
func readConf() (hosts []Host) {

	hosts = []Host{
		{"127.0.0.1", 2222, false, ""},
		{"127.0.0.1", 23339, false, ""},
		{"127.0.0.1", 16379, false, ""},
		{"127.0.0.1", 3389, false, ""},
		{"127.0.0.1", 8080, false, ""},
		{"127.0.0.1", 8081, false, ""},
		{"127.0.0.1", 18080, false, ""},
		{"127.0.0.1", 3322, false, ""},
		{"127.0.0.1", 32222, false, ""},
		{"127.0.0.1", 8082, false, ""},
		{"127.0.0.1", 8083, false, ""},
		{"127.0.0.1", 3422, false, ""},
		{"127.0.0.1", 3500, false, ""},
		{"127.0.0.1", 5000, false, ""},
		{"127.0.0.1", 32222, false, ""},
		{"127.0.0.1", 32223, false, ""},
	}

	return
}

func main() {

	hosts := readConf()
	ch := make(chan string)

	wg := sync.WaitGroup{}

	steps := 5
	length := len(hosts)
	fmt.Printf("待扫描的端口数量:%d \n", length)

	cnt := 0

	for i := 0; i < length; i = i + 5 {
		if i+steps > length {
			var scanner HostScan = &HostScanner{hosts: hosts[i:length], timeout: 2000, msgCh: ch}
			wg.Add(1)
			go scanner.Scan(&wg)
			cnt++
			break
		} else {
			var scanner HostScan = &HostScanner{hosts: hosts[i : i+5], timeout: 2000, msgCh: ch}
			wg.Add(1)
			go scanner.Scan(&wg)
			cnt++
		}
	}

	fmt.Printf("启动的协程数: %d \n", cnt)

	wg.Wait()
}
