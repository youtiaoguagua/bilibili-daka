package task

import (
	"bilibili/api"
	"bilibili/env"
	"bilibili/request"
	"github.com/tidwall/gjson"
	"math/rand"
	"strconv"
)

var followUpVideoList = make([]string, 0)
var rankVideoList = make([]string, 0)

var intit = false

func NewGetVideoId() {
	if intit {
		return
	}
	followUpVideoList = queryDynamicNew()
	rankVideoList = regionRanking()
	videoUpdate("14602398")
	intit = true
}

func videoUpdate(mid string) {
	urlParam := "?mid=" + mid + "&ps=30&tid=0&pn=1&keyword=&order=pubdate&jsonp=jsonp"
	result, err := request.GetClient().R().Get(api.GetBvidByCreate + urlParam)
	if err != nil {
		panic(err.Error())
	}
	array := gjson.Get(string(result.Body()), "data.list.vlist.#.bvid").Array()

	for _, v := range array {
		rankVideoList = append(rankVideoList, v.Str)
	}

}

func regionRanking() []string {
	arr := []int{1, 3, 4, 5, 160, 22, 119}
	rid := arr[rand.Intn(len(arr))]
	const day = 3
	urlParam := "?rid=" + strconv.Itoa(rid) + "&day=" + strconv.Itoa(day)
	result, err := request.GetClient().R().Get(api.GetRegionRanking + urlParam)
	if err != nil {
		panic(err.Error())
	}
	array := gjson.Get(string(result.Body()), "data.#.bvid").Array()
	for _, v := range array {
		rankVideoList = append(rankVideoList, v.Str)
		followUpVideoList = append(followUpVideoList, v.Str)
	}
	return rankVideoList
}

func queryDynamicNew() []string {
	urlParameter := "?uid=" + env.GetUserEnv().DedeUserID + "&type_list=8&from=&platform=web"
	result, err := request.GetClient().R().Get(api.QueryDynamicNew + urlParameter)
	if err != nil {
		panic(err.Error())
	}
	array := gjson.Get(string(result.Body()), "data.cards.#.desc.bvid").Array()
	for _, v := range array {
		followUpVideoList = append(followUpVideoList, v.Str)
	}
	return followUpVideoList
}

func GetRegionRankingVideoBvid() string {
	return rankVideoList[rand.Intn(len(rankVideoList))]
}

func GetFollowUpRandomVideoBvid() string {
	if len(followUpVideoList) == 0 {
		return GetRegionRankingVideoBvid()
	}
	return followUpVideoList[rand.Intn(len(followUpVideoList))]
}
