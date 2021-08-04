package config

type Config struct {
	// 每日设定的投币数 [0,5]
	NumberOfCoins int
	//预留硬币数
	ReserveCoins int
	//投币时是否点赞 [0,1]
	CoinAddPriority int
}

var config Config

func init() {
	config = Config{
		5, 50, 1,
	}
}

func GetConfig() Config {
	return config
}
