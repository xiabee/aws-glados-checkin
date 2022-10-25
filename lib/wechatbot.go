package lib

// To send messages by we-chat bot
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func WeChatSend(content string, key string) {
	type TEXT struct {
		Content string `json:"content"`
	}

	type Body struct {
		Msgtype string `json:"msgtype"`
		Text    TEXT   `json:"text"`
	}

	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
	sendUrl := url + key

	var sendContent Body
	sendContent.Msgtype = "text"
	sendContent.Text.Content = content
	sendBody, _ := json.Marshal(sendContent)
	resp, _ := http.Post(sendUrl, "application/json", bytes.NewBuffer(sendBody))
	defer resp.Body.Close()

	resBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
}
