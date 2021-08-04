package api

const (
	GetCionLog       = "https://api.bilibili.com/x/member/web/coin/log?jsonp=jsonp"
	LOGIN            = "https://api.bilibili.com/x/web-interface/nav"
	REWARD           = "https://api.bilibili.com/x/member/web/exp/reward"
	QueryDynamicNew  = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/dynamic_new"
	GetRegionRanking = "https://api.bilibili.com/x/web-interface/ranking/region"
	GetBvidByCreate  = "https://api.bilibili.com/x/space/arc/search"
	//VideoHeartbeat 上报观看进度
	VideoHeartbeat = "https://api.bilibili.com/x/click-interface/web/heartbeat"
	//VideoView
	VideoView      = "https://api.bilibili.com/x/web-interface/view"
	AvShare        = "https://api.bilibili.com/x/web-interface/share/add"
	Manga          = "https://manga.bilibili.com/twirp/activity.v1.Activity/ClockIn"
	LiveCheckin    = "https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/DoSign"
	NeedCoinNew    = "https://api.bilibili.com/x/web-interface/coin/today/exp"
	GetCoinBalance = "https://account.bilibili.com/site/getCoin"
	IsCoin         = "https://api.bilibili.com/x/web-interface/archive/coins"
	CoinAdd        = "https://api.bilibili.com/x/web-interface/coin/add"
	DingDing       = "https://oapi.dingtalk.com/robot/send"
)
