package push

import (
	"bilibili/api"
	"bilibili/entity"
	"bilibili/env"
	"bilibili/request"
	"github.com/sirupsen/logrus"
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
	logrus.Info(string(post.Body()))
}
