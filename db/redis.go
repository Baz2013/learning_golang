package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type HostLogin struct {
	UserId        string `json:"user_id"`
	User          string `json:"user"`
	SystemUserId  string `json:"system_user_id"`
	SystemUser    string `json:"system_user"`
	SystemUserPWD string `json:"system_user_pwd"`
	Host          string `json:"host"`
	Port          int    `json:"port"`
	Time          int64  `json:"time"` // unix 时间
	CocoInstance  string `json:"coco_instance"`
	Flag          string `json:"flag"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.249.206.1:18887",
		Password: "yHo@6JEcb_OjlNG^",
		DB:       0,
	})

	// 连接成功啦
	fmt.Println("redis 连接成功")

	// 订阅全部消息
	//pubsub := client.Subscribe("shduidi-*")
	//// 等待消息返回，原因是上一个方法不是立即返回的，囧
	//_, err := pubsub.Receive()
	//if err != nil {
	//	println(err)
	//	//logger.Error(err)
	//}

	hostLogin := HostLogin{
		User:          "gucb3",
		SystemUser:    "app",
		Host:          "10.249.206.1",
		Port:          22,
		SystemUserPWD: "*aN`QY'IyO;ZgU5q",
		Time:          1642041390,
	}

	jsons, err := json.Marshal(hostLogin)
	if err != nil {
		fmt.Println("转换错误")
	}
	n, err := client.Publish("shuidi-1", jsons).Result()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("%d clients received the message\n", n)

}
