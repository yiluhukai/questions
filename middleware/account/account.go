package account

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"yiluhuakai/questions/session"
)

//check user status,if login set session,userid,userstatus,or ceated a session

func ProcessRequest(context *gin.Context) {
	var userSession session.Session
	defer func() {
		if userSession == nil {
			userSession, _ = session.CreateSession()

		}
		context.Set(MercurySessionName, userSession)
	}()

	sessionId, err := context.Cookie(CookieSessionId)
	if err != nil {
		context.Set(MercuryUserLoginStatus, int64(0))
		context.Set(MercuryUserId, int64(0))
		return
	}
	if len(sessionId) == 0 {
		context.Set(MercuryUserLoginStatus, int64(0))
		context.Set(MercuryUserId, int64(0))
		return
	}
	// get session via cookie
	userSession, err = session.Get(sessionId)
	if err != nil {
		context.Set(MercuryUserId, int64(0))
		context.Set(MercuryUserLoginStatus, int64(0))
		return
	}

	temUserId, err := userSession.Get(MercuryUserId)
	if err != nil {
		context.Set(MercuryUserId, int64(0))
		context.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	userId, ok := temUserId.(int64)
	if !ok {
		context.Set(MercuryUserId, int64(0))
		context.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	context.Set(MercuryUserId, userId)
	context.Set(MercuryUserLoginStatus, int64(1))

}

func ProcessReponse(context *gin.Context) {
	tempSession, exist := context.Get(MercurySessionName)
	if !exist {
		return
	}
	userSession, ok := tempSession.(session.Session)
	if !ok {
		return
	}
	if userSession == nil {
		return
	}

	// session没有修改
	if !userSession.IsModify() {
		return
	}
	err := userSession.Save()

	if err != nil {
		fmt.Printf("userSession save failed = %v\n", err)
		return
	}
	// reset cookie
	sessionId := userSession.Id()
	cookie := &http.Cookie{
		Name:     CookieSessionId,
		Value:    sessionId,
		Path:     "/",
		MaxAge:   CookieMaxAge,
		HttpOnly: true,
	}
	http.SetCookie(context.Writer, cookie)
}

// 设置userId到session中
func SetUserId(context *gin.Context, userId int64) {
	tempSession, exist := context.Get(MercurySessionName)
	if !exist {
		return
	}
	session, ok := tempSession.(session.Session)
	if !ok {
		return
	}

	if session == nil {
		return
	}
	err := session.Set(MercuryUserId, userId)
	if err != nil {
		fmt.Printf("set data failed:%v\n", err)
	}

}

// 获取用户的id
func GetUserId(c *gin.Context) (uid int64, err error) {
	tempUid, exist := c.Get(MercuryUserId)
	if !exist {
		err = errors.New("user id not exists")
		return
	}
	uid, ok := tempUid.(int64)
	if !ok {
		err = errors.New("user id not exists")
		return
	}
	if uid == 0 {
		err = errors.New("user id not exists")
		return
	}
	return
}

func IsLogin(c *gin.Context) bool {
	tempStatus, exist := c.Get(MercuryUserLoginStatus)
	if !exist {
		return false
	}
	status, ok := tempStatus.(int64)
	if !ok {
		return false
	}
	return status != 0
}
