package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

/**
编译为arm版本 （bash 终端下执行）
env GOOS=linux GOARCH=arm64 go build -o port_scan_arm64
*/

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
	Scan(ctx context.Context)
}

// Scan 端口扫描的主要执行逻辑
func (hs *HostScanner) Scan(ctx context.Context) {
	fmt.Println("开始扫描....")
	for _, host := range hs.hosts {
		address := fmt.Sprintf("%s:%d", host.Ip, host.Port)
		conn, err := net.DialTimeout("tcp", address, time.Duration(hs.timeout)*time.Millisecond)
		if err != nil {
			msg := fmt.Sprintf("%s 端口不通！, Error: %s", address, err)
			hs.msgCh <- msg
			//fmt.Println(msg)
		} else {
			msg := fmt.Sprintf("%s 端口可访问！", address)
			hs.msgCh <- msg
			// 关闭连接
			if err := conn.Close(); err != nil {
				fmt.Println("关闭连接错误！！！！！")
			}
		}
	}

	//wg.Done() // 任务结束
}

// readData 读取文本文件获取要扫描的主机端口
func readData(file string) (hosts []Host, err error) {

	hosts = []Host{}

	fp, err := os.Open(file)
	if err != nil {
		fmt.Printf("打开文件错误:%s \n", err) //打开文件错误
		return nil, err
	}

	defer fp.Close()
	buf := bufio.NewScanner(fp)
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		line = strings.Trim(line, "\n")
		line = strings.TrimSpace(line)
		items := strings.Split(line, ":")
		if len(items) != 2 {
			fmt.Printf("Warning: 有错误数据！ %s\n", line)
			continue
		}
		port, err := strconv.Atoi(items[1])
		if err != nil {
			fmt.Printf("Warning: 端口不是数字！ %s\n", line)
			continue
		}
		ip := strings.TrimSpace(items[0])
		host := Host{
			Ip:          ip,
			Port:        port,
			IsReachable: false,
			msg:         "",
		}

		hosts = append(hosts, host)
	}

	return
}

var (
	infoFlag = false
	step     = 10
	datafile = ""
)

func init() {
	//flag.BoolVar(&daemonFlag, "d", false, "start as Daemon")
	flag.StringVar(&datafile, "f", "D:\\github\\learning_golang\\tools\\address.dat", "config.yml path")
	flag.IntVar(&step, "s", 10, "batch size")
	flag.BoolVar(&infoFlag, "V", false, "version info")
}

func main() {

	// 解析命令行参数
	flag.Parse()
	if infoFlag {
		fmt.Println("usage:   port_scan -f .\\address.dat -s 100")
		return
	}

	fmt.Printf("data file:%s, step:%d \n", datafile, step)

	hosts, err := readData(datafile)
	if err != nil {
		fmt.Println("读取数据文件失败")
		return
	}
	if len(hosts) == 0 {
		fmt.Println("未读取到数据")
		return
	}
	ch := make(chan string)

	//wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	steps := step
	length := len(hosts)
	fmt.Printf("待扫描的端口数量:%d \n", length)

	cnt := 0

	for i := 0; i < length; i = i + steps {
		if i+steps > length {
			var scanner HostScan = &HostScanner{hosts: hosts[i:length], timeout: 2000, msgCh: ch}
			//wg.Add(1)
			go scanner.Scan(ctx)
			cnt++
			break
		} else {
			var scanner HostScan = &HostScanner{hosts: hosts[i : i+steps], timeout: 2000, msgCh: ch}
			//wg.Add(1)
			go scanner.Scan(ctx)
			cnt++
		}
	}

	fmt.Printf("启动的协程数: %d \n", cnt)

	i := 1
	for msg := range ch {
		fmt.Println(msg)
		if i == length {
			break
		}
		i++
	}

	cancel()

	//wg.Wait()
}
