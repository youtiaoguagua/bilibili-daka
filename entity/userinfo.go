package entity

import "encoding/json"

type UserInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		IsLogin       bool   `json:"isLogin"`
		EmailVerified int    `json:"email_verified"`
		Face          string `json:"face"`
		LevelInfo     struct {
			CurrentLevel int         `json:"current_level"`
			CurrentMin   int         `json:"current_min"`
			CurrentExp   int         `json:"current_exp"`
			NextExp      json.Number `json:"next_exp"`
		} `json:"level_info"`
		Mid            int     `json:"mid"`
		MobileVerified int     `json:"mobile_verified"`
		Money          float64 `json:"money"`
		Moral          int     `json:"moral"`
		Official       struct {
			Role  int    `json:"role"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
			Type  int    `json:"type"`
		} `json:"official"`
		OfficialVerify struct {
			Type int    `json:"type"`
			Desc string `json:"desc"`
		} `json:"officialVerify"`
		Pendant struct {
			Pid               int    `json:"pid"`
			Name              string `json:"name"`
			Image             string `json:"image"`
			Expire            int    `json:"expire"`
			ImageEnhance      string `json:"image_enhance"`
			ImageEnhanceFrame string `json:"image_enhance_frame"`
		} `json:"pendant"`
		Scores       int    `json:"scores"`
		Uname        string `json:"uname"`
		VipDueDate   int    `json:"vipDueDate"`
		VipStatus    int    `json:"vipStatus"`
		VipType      int    `json:"vipType"`
		VipPayType   int    `json:"vip_pay_type"`
		VipThemeType int    `json:"vip_theme_type"`
		VipLabel     struct {
			Path        string `json:"path"`
			Text        string `json:"text"`
			LabelTheme  string `json:"label_theme"`
			TextColor   string `json:"text_color"`
			BgStyle     int    `json:"bg_style"`
			BgColor     string `json:"bg_color"`
			BorderColor string `json:"border_color"`
		} `json:"vip_label"`
		VipAvatarSubscript int    `json:"vip_avatar_subscript"`
		VipNicknameColor   string `json:"vip_nickname_color"`
		Vip                struct {
			Type       int `json:"type"`
			Status     int `json:"status"`
			DueDate    int `json:"due_date"`
			VipPayType int `json:"vip_pay_type"`
			ThemeType  int `json:"theme_type"`
			Label      struct {
				Path        string `json:"path"`
				Text        string `json:"text"`
				LabelTheme  string `json:"label_theme"`
				TextColor   string `json:"text_color"`
				BgStyle     int    `json:"bg_style"`
				BgColor     string `json:"bg_color"`
				BorderColor string `json:"border_color"`
			} `json:"label"`
			AvatarSubscript    int    `json:"avatar_subscript"`
			NicknameColor      string `json:"nickname_color"`
			Role               int    `json:"role"`
			AvatarSubscriptURL string `json:"avatar_subscript_url"`
		} `json:"vip"`
		Wallet struct {
			Mid           int `json:"mid"`
			BcoinBalance  int `json:"bcoin_balance"`
			CouponBalance int `json:"coupon_balance"`
			CouponDueTime int `json:"coupon_due_time"`
		} `json:"wallet"`
		HasShop        bool   `json:"has_shop"`
		ShopURL        string `json:"shop_url"`
		AllowanceCount int    `json:"allowance_count"`
		AnswerStatus   int    `json:"answer_status"`
	} `json:"data"`
}

var userInfo UserInfo

func SetUserInfo(info UserInfo) {
	userInfo = info
}

func GetUserInfo() UserInfo {
	return userInfo
}
