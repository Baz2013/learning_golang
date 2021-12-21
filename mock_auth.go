package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

//func main() {
//	auth()
//}

func auth() {
	signature := makeSignature("24d05dfb-2703-4976-9e58-969162a96458", "Tue, 21 Dec 2021 07:03:24 GMT")
	fmt.Println("      :" + signature) // MTEzM2Y1ZGI4ZGFhMzQ5ZDA3ZTBkZmNiZjY1YTIyY2I=
	fmt.Println("Target:" + "MTEzM2Y1ZGI4ZGFhMzQ5ZDA3ZTBkZmNiZjY1YTIyY2I=")
	fmt.Println(signature == "MTEzM2Y1ZGI4ZGFhMzQ5ZDA3ZTBkZmNiZjY1YTIyY2I=")
}

/**
签名算法
*/
func makeSignature(secret string, date string) string {
	data := fmt.Sprintf("%s\n%s", secret, date)
	h := md5.New()
	h.Write([]byte(data))
	//d5 := h.Sum(nil)
	//sEnc := base64.StdEncoding.EncodeToString(d5)

	return base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(h.Sum(nil))))
}
