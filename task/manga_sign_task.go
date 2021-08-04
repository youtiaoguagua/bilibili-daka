package task

import (
	"bilibili/api"
	"bilibili/request"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type MangaSignTask struct {
}

func (m MangaSignTask) Run() {
	result, err := request.GetClient().R().
		SetFormData(map[string]string{"platform": "android"}).
		Post(api.Manga)
	if err != nil {
		logrus.Error("漫画签到失败！")
		return
	}
	data := gjson.Get(string(result.Body()), "data").Exists()
	if data {
		logrus.Info("完成漫画签到")
	} else {
		logrus.Info("哔哩哔哩漫画已经签到过了")
	}
}

func (m MangaSignTask) GetName() string {
	return "漫画签到"
}
