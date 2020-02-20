package main

import (
	"fmt"
	"logger"
	"questions/controller/account"
	"questions/controller/answer"
	"questions/controller/category"
	"questions/controller/comment"
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
	//err = session.Init("redis", "192.168.1.12:6379")
	err = session.Init("memory", "")
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
	router.GET("/api/category/list", category.GetCategoryListHandle)
	router.POST("/api/ask/submit", auth_middleware.AuthMiddleWare, question.QuestionSubmitHandle)
	router.GET("/api/question/list", category.GetQuestionListHandle)
	router.GET("/api/question/detail", question.QuestionDetailHandle)
	router.GET("/api/answer/list", answer.AnswerListHandle)
	router.POST("/api/answer/commit", auth_middleware.AuthMiddleWare, answer.AnswerCommitHandle)
	group := router.Group("/api/comment", auth_middleware.AuthMiddleWare)

	group.POST("/post_comment", comment.PostCommentHandle)
	_ = router.Run(":9090")
}
