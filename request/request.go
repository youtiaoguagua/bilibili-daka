package request

import (
	"bilibili/env"
	"github.com/go-resty/resty/v2"
	"net/http"
)

var client *resty.Client

func init() {
	client = resty.New()
	cookies := []*http.Cookie{
		{
			Name:  "DedeUserID",
			Value: env.GetUserEnv().DedeUserID,
		},
		{
			Name:  "SESSDATA",
			Value: env.GetUserEnv().SESSDATA,
		}, {
			Name:  "bili_jct",
			Value: env.GetUserEnv().Bili_jct,
		},
	}
	client.SetCookies(cookies)
	client.SetHeader("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	client.SetHeader("Connection", "keep-alive")
	client.SetHeader("Accept-Encoding", "gzip, deflate, br")
	client.SetHeader("accept", "*/*")
}

func GetClient() *resty.Client {
	return client
}
