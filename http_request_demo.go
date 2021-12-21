package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://172.16.6.126:18080/api/users/v1/profile/")

	req, _ := http.NewRequest("GET", "http://172.16.6.126:18080/api/users/v1/profile/", nil)
	// 比如说设置个token
	req.Header.Set("Date", "Tue, 21 Dec 2021 07:43:24 GMT")
	req.Header.Set("AUTHORIZATION", "Sign 08fb74b0-efdc-43f2-b8d2-ce48097e0879:M2ZjM2JlNjgyNTZiZWYxNzg1MGUzYTZhYThkMGFkNjQ=")
	// 再设置个js
	req.Header.Set("Content-Type", "application/json")

	resp, err := (&http.Client{}).Do(req)
	fmt.Println(resp.StatusCode)
	//resp, err := http.Get(serviceUrl + "/topic/query/false/lsj")
	if err != nil {
		fmt.Println("query topic failed", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	//var serviceResp ServiceResp
	//respByte, _ := ioutil.ReadAll(resp.Body)

}
