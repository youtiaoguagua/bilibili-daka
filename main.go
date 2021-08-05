package main

import (
	"bilibili/entity"
	"bilibili/push"
	"bilibili/task"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	logrus.Info("bilibili签到开始！")
	defer func() {
		if err := recover(); err != nil {
			logrus.Error(err)
		}
		pushMessage()
	}()
	task := []task.Task{
		task.UserCheckTask{},
		task.CoinLogsTask{},
		//task.VideoWatchTask{},
		//task.MangaSignTask{},
		//task.LiveCheckinTask{},
		//task.CoinAddTask{},
	}
	for _, v := range task {
		logrus.Infof("------%v开始------", v.GetName())
		v.Run()
		logrus.Infof("------%v结束------", v.GetName())
		time.Sleep(time.Duration(2) * time.Second)
	}
	logrus.Info("本日任务已全部执行完毕")
	calculateUpgradeDays()
}

func pushMessage() {
	file, err := os.OpenFile("./log.txt", os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	//发送推送通知
	dingPush := push.DingDingPush{}
	dingPush.DoPush(entity.MetaInfo{}, string(content))
}

func calculateUpgradeDays() {
	userInfo := entity.GetUserInfo()
	if userInfo == (entity.UserInfo{}) {
		logrus.Info("未请求到用户信息，暂无法计算等级相关数据")
		return
	}

	todayExp := 15
	todayExp += task.GetHaveCastCoin() * 10
	logrus.Infof("今日获得的总经验值为: %v", todayExp)

	levelInfo := userInfo.Data.LevelInfo
	var needExp int
	if "--" == levelInfo.NextExp.String() {
		needExp = levelInfo.CurrentExp
	} else {
		i, _ := levelInfo.NextExp.Int64()
		needExp = int(i)
	}
	needExp -= levelInfo.CurrentExp

	if levelInfo.CurrentLevel < 6 {
		logrus.Infof("按照当前进度，升级到升级到Lv%v 还需要: %v天", levelInfo.CurrentLevel+1, needExp/todayExp)
	} else {
		logrus.Infof("当前等级Lv6，经验值为：%v", levelInfo.CurrentExp)
	}
}
