package task

import (
	"bilibili/api"
	"bilibili/config"
	"bilibili/env"
	"bilibili/request"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type CoinAddTask struct {
}

func (c CoinAddTask) Run() {
	//投币最多操作数 解决csrf校验失败时死循环的问题
	addCoinOperateCount := 0
	//安全检查，最多投币数
	maxNumberOfCoins := 5
	//获取自定义配置投币数 配置写在src/main/resources/config.json中
	setCoin := config.GetConfig().NumberOfCoins
	// 预留硬币数
	reserveCoins := config.GetConfig().ReserveCoins

	coinAddPriority := config.GetConfig().CoinAddPriority
	useCoin := GetHaveCastCoin()

	logrus.Info("自定义投币数为: " + strconv.Itoa(setCoin) + "枚," + "程序执行前已投: " + strconv.Itoa(useCoin) + "枚")

	needCoins := setCoin - useCoin

	beforeAddCoinBalance := GetCoinBalance()
	coinBalance := int(math.Floor(beforeAddCoinBalance))

	if needCoins <= 0 {
		logrus.Info("已完成设定的投币任务，今日无需再投币了")
	} else {
		logrus.Info("投币数调整为: " + strconv.Itoa(needCoins) + "枚")
		//投币数大于余额时，按余额投
		if needCoins > coinBalance {
			logrus.Info("完成今日设定投币任务还需要投: " + strconv.Itoa(needCoins) + "枚硬币，但是余额只有: " + strconv.Itoa(coinBalance))
			logrus.Info("投币数调整为: " + strconv.Itoa(coinBalance))
			needCoins = coinBalance
		}
	}

	if coinBalance < reserveCoins {
		logrus.Infof("剩余硬币数为%v,低于预留硬币数%v,今日不再投币", beforeAddCoinBalance, reserveCoins)
		logrus.Infof("tips: 当硬币余额少于你配置的预留硬币数时，则会暂停当日投币任务")
		return
	}

	logrus.Infof("投币前余额为 : %v", beforeAddCoinBalance)
	/*
	 * 开始投币
	 * 请勿修改 max_numberOfCoins 这里多判断一次保证投币数超过5时 不执行投币操作
	 * 最后一道安全判断，保证即使前面的判断逻辑错了，也不至于发生投币事故
	 */
	for needCoins > 0 && needCoins <= maxNumberOfCoins {
		var bvid string
		if coinAddPriority == 1 && addCoinOperateCount < 7 {
			bvid = GetFollowUpRandomVideoBvid()
		} else {
			bvid = GetRegionRankingVideoBvid()
		}

		addCoinOperateCount++
		WatchVideo(bvid)
		flag := coinAdd(bvid)
		if flag {
			t := rand.Intn(4) + 2
			time.Sleep(time.Duration(t) * time.Second)
			needCoins--
		}
		if addCoinOperateCount > 15 {
			logrus.Info("尝试投币/投币失败次数太多")
			break
		}
	}
	logrus.Infof("投币任务完成后余额为: %v", GetCoinBalance())
}

func coinAdd(bvid string) bool {
	videoTitle := GetVideoTitle(bvid)
	//判断曾经是否对此av投币过
	if !isCoinAdded(bvid) {
		result, err := request.GetClient().
			SetHeader("Referer", "https://www.bilibili.com/video/"+bvid).
			SetHeader("Origin", "https://www.bilibili.com").R().
			SetFormData(map[string]string{
				"bvid":         bvid,
				"multiply":     strconv.Itoa(1),
				"select_like":  strconv.Itoa(1),
				"cross_domain": "true",
				"csrf":         env.GetUserEnv().Bili_jct,
			}).Post(api.CoinAdd)
		if err != nil {
			panic(err.Error())
		}
		jsonStr := string(result.Body())
		code := gjson.Get(jsonStr, "code").Num
		if code == 0 {
			logrus.Info("为 " + videoTitle + " 投币成功")
			return true
		} else {
			logrus.Debug("投币失败" + gjson.Get(jsonStr, "message").Str)
			return false
		}

	} else {
		logrus.Debug("已经为" + videoTitle + "投过币了")
		return false
	}

}

func isCoinAdded(bvid string) bool {
	result, err := request.GetClient().R().SetQueryParam("bvid", bvid).Get(api.IsCoin)
	if err != nil {
		panic(err.Error())
	}
	multiply := gjson.Get(string(result.Body()), "data.multiply").Num
	if multiply > 0 {
		logrus.Infof("之前已经为av %v 投过%v枚硬币啦", bvid, multiply)
		return true
	} else {
		return false
	}
}

func (c CoinAddTask) GetName() string {
	return "投币任务"
}
