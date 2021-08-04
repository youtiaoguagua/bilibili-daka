package task

import (
	"bilibili/api"
	"bilibili/request"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"strings"
)

func GetVideoTitle(bvid string) string {
	result, err := request.GetClient().R().
		SetQueryParam("bvid", bvid).
		Get(api.VideoView)
	if err != nil {
		panic(err.Error())
	}
	t := gjson.Get(string(result.Body()), "code")
	code := t.Num
	var title string
	t = gjson.Get(string(result.Body()), "data")
	if code == 0 {
		title = t.Get("owner.name").Str + ": "
		title += t.Get("title").Str
	} else {
		title = "未能获取标题"
	}
	return strings.ReplaceAll(title, "&", "-")
}

func GetHaveCastCoin() int {
	result, err := request.GetClient().R().Get(api.NeedCoinNew)
	if err != nil {
		panic(err.Error())
	}
	count := gjson.Get(string(result.Body()), "data").Num
	return int(count) / 10
}

func GetCoinBalance() float64 {
	result, err := request.GetClient().R().Get(api.GetCoinBalance)
	if err != nil {
		panic(err.Error())
	}

	jsonStr := string(result.Body())
	code := gjson.Get(jsonStr, "code").Num
	data := gjson.Get(jsonStr, "data")
	if code == 0 {
		if data.Get("money").Exists() {
			return data.Get("money").Num
		} else {
			return 0.0
		}
	}

	logrus.Debug("请求硬币余额接口错误，请稍后重试。错误请求信息：" + jsonStr)
	return 0.0
}
