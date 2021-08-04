package env

import "os"

type UserEnv struct {
	DedeUserID    string
	SESSDATA      string
	Bili_jct      string
	DingDingToken string
}

var userEnv UserEnv

func init() {
	//userEnv = UserEnv{"12976104",
	//	"f72e28ba%2C1641476752%2Cb4a2c%2A71",
	//	"6f453f9726aaf71b65bd9a50cfea85a1"}
	DedeUserID := os.Getenv("DedeUserID")
	SESSDATA := os.Getenv("SESSDATA")
	BILI_JCT := os.Getenv("BILI_JCT")
	DingdingToken := os.Getenv("DINGDING")
	if DedeUserID == "" || SESSDATA == "" || BILI_JCT == "" {
		panic("env is not valid!")
	}
	userEnv = UserEnv{DedeUserID, SESSDATA, BILI_JCT, DingdingToken}
}

func GetUserEnv() UserEnv {
	return userEnv
}
