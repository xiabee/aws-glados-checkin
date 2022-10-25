package main

import (
	"checkin/lib"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func start() {
	host := "https://glados.rocks"
	checkInUrl := "/api/user/checkin"
	checkInStatusUrl := "/api/user/status"

	cookie := os.Getenv("COOKIE")
	wechatKey := os.Getenv("WECHAT_KEY")
	// use env or git action secrets

	msg1 := lib.CheckIn(host+checkInUrl, cookie)
	msg2 := lib.CheckStatus(host+checkInStatusUrl, cookie)
	msg := msg1 + "\n" + msg2
	lib.WeChatSend(msg, wechatKey)
	// send message by wechat-bot
}
func main() {
	lambda.Start(start)
}
