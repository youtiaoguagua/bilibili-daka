package task

import (
	"bilibili/api"
	"bilibili/entity"
	"bilibili/request"
	"github.com/sirupsen/logrus"
)

type CoinLogsTask struct {
}

func (c CoinLogsTask) Run() {
	coinInfo := entity.CoinInfo{}
	_, err := request.GetClient().R().SetResult(&coinInfo).Get(api.GetCionLog)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	logrus.WithFields(logrus.Fields{
		"coin log info res is:": coinInfo,
	}).Debug()
	logrus.Infof("最近一周共计%v条硬币记录", coinInfo.Data.Count)

	income := 0.0
	expend := 0.0
	for _, v := range coinInfo.Data.List {
		if v.Delta > 0 {
			income += float64(v.Delta)
		} else {
			expend += float64(v.Delta)
		}
	}

	logrus.Infof("最近一周收入%v个硬币", income)
	logrus.Infof("最近一周支出%v个硬币", expend)
}

func (c CoinLogsTask) GetName() string {
	return "硬币日志"
}
