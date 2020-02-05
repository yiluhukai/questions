package main

import (
	"fmt"
	"logger"
	"questions/controller/account"
	"questions/controller/category"
	"questions/controller/question"
	"questions/db"
	"questions/filter"
	"questions/gen_id"
	auth_middleware "questions/middleware/account"
	"questions/session"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 初始化日志
	err := logger.InitLogger("console", "", "", logger.LogLevelDebug, "")
	if err != nil {
		panic(err)
	}
	// 初始化过滤非法字的过滤器
	err = filter.Init("./data/filter.data.txt")
	if err != nil {
		panic(err)
	}
	err = db.InitDb()
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
	router.GET("/api/category/list", category.GetCategoryList)
	router.POST("/api/ask/submit", auth_middleware.AuthMiddleWare, question.QuestionSubmitHandle)
	_ = router.Run(":9090")
}
