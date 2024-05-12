package system

import (
	"fmt"
	"github.com/apistd/uni-go-sdk"
	unisms "github.com/apistd/uni-go-sdk/sms"
	"log"
	"math/rand"
	"time"
)

func SendMessage(phone string) (*uni.UniResponse, error) {

	// 初始化
	client := unisms.NewClient("PddHRiyShWHL42QwM2QeSBMno8C5YMKGmZ9YBCkfghA4S9zYx", "mc7ve5BmjTShjMzKNrXPFHzLTyiqer")
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	// 构建信息
	message := unisms.BuildMessage()
	message.SetTo(phone)
	message.SetSignature("2333小站")
	message.SetTemplateId("pub_verif_basic")
	message.SetTemplateData(map[string]string{"code": code}) // 设置自定义参数 (变量短信)
	// 发送短信
	res, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, err
}
