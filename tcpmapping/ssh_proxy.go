package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"os"
	"os/user"
	"time"
)

var (
	DefaultSShTcpTimeout = 15 * time.Second // 与ssh建立连接的默认时间，自己设置一个就行
)

// 通过代理登录到远程主机执行命令
// local -> 127.0.0.1:3222 -> 172.18.0.5:22

func main() {

	session, client, err := Connect()
	defer session.Close()
	defer client.Close()

	if err != nil {
		fmt.Errorf("创建连接失败")
	}

	output, err := session.CombinedOutput("hostname")

	fmt.Printf("命令输出： %s", output)

}

func Connect() (*ssh.Session, *ssh.Client, error) {
	proxyServer := "127.0.0.1"
	proxyPort := "3222"
	authConfig := AuthConfig{User: "root", Password: "123456"}

	if err := authConfig.ApplyConfig(); err != nil {
		return nil, nil, err
	}
	proxyConfig := authConfig.ClientConfig

	targetServer := "172.18.0.5"
	targetPort := "22"
	targetConfig := proxyConfig // 用户密码相同

	proxyClient, err := ssh.Dial("tcp", net.JoinHostPort(proxyServer, proxyPort), proxyConfig)
	if err != nil {
		return nil, nil, err
	}

	conn, err := proxyClient.Dial("tcp", net.JoinHostPort(targetServer, targetPort))
	if err != nil {
		return nil, nil, err
	}

	ncc, chans, reqs, err := ssh.NewClientConn(conn, net.JoinHostPort(targetServer, targetPort), targetConfig)
	if err != nil {
		return nil, nil, err
	}

	client := ssh.NewClient(ncc, chans, reqs)

	session, err := client.NewSession()
	if err != nil {
		return nil, nil, err
	}

	return session, client, nil
}

type AuthConfig struct {
	*ssh.ClientConfig
	User     string
	Password string
	KeyFile  string
	Timeout  time.Duration
}

func (a *AuthConfig) setDefault() {
	if a.User == "" {
		a.User = getCurrentUser()
	}

	if a.KeyFile == "" {
		userHome, _ := os.UserHomeDir()
		a.KeyFile = fmt.Sprintf("%s/.ssh/id_rsa", userHome)
	}

	if a.Timeout == 0 {
		a.Timeout = DefaultSShTcpTimeout
	}
}

func (a *AuthConfig) SetAuthMethod() (ssh.AuthMethod, error) {
	a.setDefault()
	if a.Password != "" {
		return ssh.Password(a.Password), nil
	}
	data, err := ioutil.ReadFile(a.KeyFile)
	if err != nil {
		return nil, err
	}
	singer, err := ssh.ParsePrivateKey(data)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(singer), nil
}

func (a *AuthConfig) ApplyConfig() error {
	authMethod, err := a.SetAuthMethod()
	if err != nil {
		return err
	}
	a.ClientConfig = &ssh.ClientConfig{
		User:            a.User,
		Auth:            []ssh.AuthMethod{authMethod},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         a.Timeout,
	}
	return nil
}

// 返回当前用户名
func getCurrentUser() string {
	user, _ := user.Current()
	return user.Username
}
