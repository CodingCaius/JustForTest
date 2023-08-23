package main

import "fmt"


type DocContentResponse struct {
	Content   string `json:"content"`
	IsSee     bool   `json:"isSee"`     // 是否试看
	IsPwd     bool   `json:"isPwd"`     // 是否需要密码
	IsColl    bool   `json:"isColl"`    // 用户是否收藏
	LookCount int    `json:"lookCount"` // 浏览量
	DiggCount int    `json:"diggCount"` // 点赞量
	CollCount int    `json:"collCount"` // 收藏量
}

func main() {
	fmt.Println("铃木爱理嫁给我")
	fmt.Println("我的最爱")
}