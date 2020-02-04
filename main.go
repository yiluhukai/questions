package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"questions/controller/account"
	"questions/db"
	"questions/gen_id"
	"questions/session"
)

func main() {

	router := gin.Default()
	err := db.InitDb()
	if err != nil {
		panic(err)
	}
	// 设置生成id
	err = gen_id.Init(uint16(0))
	if err != nil {
		panic(err)
	}
	// 初始化Session
	err = session.Init("redis", "192.168.1.12:6379")

	if err != nil {
		fmt.Printf("init redis error:%v", err)
		panic(err)
	}
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.Static("/js", "./static/js")
	router.Static("/css", "./static/css")
	router.Static("/fonts", "./static/fonts")
	router.POST("/api/user/register", account.RegisterHandle)
	router.POST("/api/user/login", account.LoginHandle)
	_ = router.Run(":9090")
}
