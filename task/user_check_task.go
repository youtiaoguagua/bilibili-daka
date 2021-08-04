package task

import (
	"bilibili/api"
	"bilibili/entity"
	"bilibili/request"
	"github.com/sirupsen/logrus"
)

type UserCheckTask struct {
}

func (u UserCheckTask) Run() {
	userInfo := entity.UserInfo{}
	result, _ := request.GetClient().R().SetResult(&userInfo).Get(api.LOGIN)

	logrus.WithFields(logrus.Fields{
		"user check login res :": userInfo,
	}).Debug()

	if userInfo.Code != 0 || !userInfo.Data.IsLogin {
		logrus.Errorf("request info is %v", string(result.Body()))
		logrus.Info("Cookies可能失效了,请仔细检查配置中的DEDEUSERID SESSDATA BILI_JCT三项的值是否正确、过期")
		panic("cookies is not valid!")
	}

	logrus.WithFields(logrus.Fields{
		"用户名称:": userInfo.Data.Uname,
		"硬币余额:": userInfo.Data.Money,
	}).Info()

	entity.SetUserInfo(userInfo)
}

func (u UserCheckTask) GetName() string {
	return "登录检查"
}
