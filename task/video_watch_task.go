package task

import (
	"bilibili/api"
	"bilibili/env"
	"bilibili/request"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"math/rand"
	"strconv"
)

type VideoWatchTask struct {
}

func (v VideoWatchTask) Run() {
	NewGetVideoId()
	rewardInfo, err := getDailyTaskStatus()
	if err != nil {
		return
	}
	bvid := GetRegionRankingVideoBvid()
	if !rewardInfo.Data.Watch {
		WatchVideo(bvid)
	} else {
		logrus.Info("本日观看视频任务已经完成了，不需要再观看视频了")
	}

	if !rewardInfo.Data.Share {
		dailyAvShare(bvid)
	} else {
		logrus.Info("本日分享视频任务已经完成了，不需要再分享视频了")
	}

}

func dailyAvShare(bvid string) {
	result, err := request.GetClient().R().
		SetFormData(map[string]string{
			"bvid": bvid,
			"csrf": env.GetUserEnv().Bili_jct,
		}).Post(api.AvShare)
	if err != nil {
		panic(err.Error())
	}
	videoTitle := GetVideoTitle(bvid)
	t := gjson.Get(string(result.Body()), "code")
	code := t.Num
	if code == 0 {
		logrus.Infof("视频: " + videoTitle + " 分享成功")
	} else {
		logrus.Debugf("视频分享失败，原因: " + gjson.Get(string(result.Body()), "message").Str)
		logrus.Debug("开发者提示: 如果是csrf校验失败请检查BILI_JCT参数是否正确或者失效")
	}
}

func WatchVideo(bvid string) {
	playedTime := rand.Intn(90) + 1
	result, err := request.GetClient().R().
		SetFormData(map[string]string{
			"bvid":        bvid,
			"played_time": strconv.Itoa(playedTime),
		}).
		Post(api.VideoHeartbeat)
	if err != nil {
		panic(err.Error())
	}
	videoTitle := GetVideoTitle(bvid)
	t := gjson.Get(string(result.Body()), "code")
	code := t.Num
	if code == 0 {
		logrus.Infof("视频: " + videoTitle + "播放成功,已观看到第" + strconv.Itoa(playedTime) + "秒")
	} else {
		logrus.Debugf("视频: " + videoTitle + "播放失败,原因: " + gjson.Get(string(result.Body()), "message").Str)
	}
}

func getDailyTaskStatus() (RewardInfo, error) {
	rewardInfo := RewardInfo{}
	_, err := request.GetClient().R().SetResult(&rewardInfo).Get(api.REWARD)
	if err != nil {
		logrus.Error("fail to get reward info!")
		return rewardInfo, err
	}
	if rewardInfo.Code == 0 {
		logrus.WithFields(logrus.Fields{
			"reward res is:": rewardInfo,
		}).Debug()
		logrus.Info("请求本日任务完成状态成功!")
		return rewardInfo, nil
	} else {
		//偶尔会失败再试一次
		_, err := request.GetClient().R().SetResult(&rewardInfo).Get(api.REWARD)
		return rewardInfo, err
	}

}

type RewardInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Login        bool `json:"login"`
		Watch        bool `json:"watch"`
		Coins        int  `json:"coins"`
		Share        bool `json:"share"`
		Email        bool `json:"email"`
		Tel          bool `json:"tel"`
		SafeQuestion bool `json:"safe_question"`
		IdentifyCard bool `json:"identify_card"`
	} `json:"data"`
}

func (r VideoWatchTask) GetName() string {
	return "看视频分享视频"
}
