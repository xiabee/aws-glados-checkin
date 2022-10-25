package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func CheckIn(url string, cookie string) string {
	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	msg := "Check in Failed! Check your Cookie!"
	// if cookie is wrong

	token := "{\"token\":\"glados.network\"}"
	// request body, token is necessary

	glados, err := http.NewRequest(http.MethodPost, url, strings.NewReader(token))
	if err != nil {
		panic(err)
	}
	glados.Header.Set("Accept", "application/json, text/plain, */*")
	glados.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// set Content-Type, json is necessary

	glados.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	// you can change your User-Agent here
	glados.Header.Set("Cookie", cookie)
	request, err := (&http.Client{}).Do(glados)
	if err != nil {
		fmt.Println(err)
	}
	defer request.Body.Close()
	// to check in automatically

	res, _ := io.ReadAll(request.Body)
	var basket Response
	err = json.Unmarshal(res, &basket)
	if err != nil {
		fmt.Println(err)
	}
	msg = basket.Message
	fmt.Println(basket.Message)
	return msg
}
