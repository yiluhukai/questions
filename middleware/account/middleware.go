package account

import (
	"github.com/gin-gonic/gin"
	"yiluhuakai/questions/util"
)

func AuthMiddleWare(c *gin.Context) {
	ProcessRequest(c)
	//  检验用户是否登录了
	if !IsLogin(c) {
		util.ResponseError(c, util.ErrCodeNotLogin)
		c.Abort()
	}
	c.Next()
}
