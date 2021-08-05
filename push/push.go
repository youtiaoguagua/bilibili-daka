package push

import (
	"bilibili/api"
	"bilibili/entity"
	"bilibili/env"
	"bilibili/request"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type Push interface {
	DoPush(info entity.MetaInfo, content string)
}

type DingDingPush struct {
}

func (d DingDingPush) DoPush(info entity.MetaInfo, content string) {
	post, err := request.GetClient().R().
		SetBody(map[string]interface{}{
			"msgtype": "text",
			"text":    map[string]interface{}{"content": content},
		}).SetQueryParam("access_token", env.GetUserEnv().DingDingToken).
		Post(api.DingDing)
	if err != nil {
		panic(err.Error())
	}
	code := gjson.Get(string(post.Body()), "errcode").Num
	if code == 0 {
		logrus.Info("钉钉推送成功！")
	} else {
		logrus.Infof("钉钉推送失败！请求结果：%v", string(post.Body()))
	}
}
