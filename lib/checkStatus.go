package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	Email    string `json:"email"`
	LeftDays string `json:"leftDays"`
}
type Response struct {
	Code int  `json:"code"`
	Data Data `json:"data"`
}

func CheckStatus(url string, cookie string) string {

	glados, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

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
	fmt.Println(basket.Data.Email)
	fmt.Println(basket.Data.LeftDays)
	msg := "Email: " + basket.Data.Email + "\n" + "Left days: " + basket.Data.LeftDays
	return msg
}
