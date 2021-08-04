package task

import (
	"bilibili/api"
	"bilibili/request"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type LiveCheckinTask struct {
}

func (l LiveCheckinTask) Run() {
	result, err := request.GetClient().R().Get(api.LiveCheckin)
	if err != nil {
		logrus.Error("漫画签到失败！")
		return
	}
	code := gjson.Get(string(result.Body()), "code").Num
	if code == 0 {
		data := gjson.Get(string(result.Body()), "data")
		logrus.Info("直播签到成功，本次签到获得" + data.Get("text").Str + "," + data.Get("specialText").Str)
	} else {
		message := gjson.Get(string(result.Body()), "message")
		logrus.Debug("直播签到失败:" + message.Str)
	}
}

func (l LiveCheckinTask) GetName() string {
	return "直播签到"
}
