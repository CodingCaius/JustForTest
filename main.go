package main

import "fmt"


type DocContentResponse struct {
	content   string `json:"content"`
	IsSee     bool   `json:"isSee"`     // 是否试看
	IsPwd     bool   `json:"isPwd"`     // 是否需要密码
	IsColl    bool   `json:"isColl"`    // 用户是否收藏
	LookCount int    `json:"lookCount"` // 浏览量
	DiggCount int    `json:"diggCount"` // 点赞量
	CollCount int    `json:"collCount"` // 收藏量
}

func main() {
	fmt.Println("铃木爱理嫁给我")

	//写一个for循环输出1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	//编写程序，输出五角星
	for i := 1; i <= 5; i++ {
		fmt.Print("*")
	}




}