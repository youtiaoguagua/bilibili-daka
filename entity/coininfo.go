package entity

type CoinInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		List []struct {
			Time   string `json:"time"`
			Delta  int    `json:"delta"`
			Reason string `json:"reason"`
		} `json:"list"`
		Count int `json:"count"`
	} `json:"data"`
}
