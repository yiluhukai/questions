package account

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"questions/session"
)

//check user status,if login set session,userid,userstatus,or ceated a session

func ProcessRequest(context *gin.Context) {
	var userSession session.Session
	defer func() {
		if userSession == nil {
			userSession, err := session.CreateSession()
			if err != nil {
				fmt.Printf("%v", err)
			}
		}
		context.Set(MercurySessionName, userSession)
	}()

	sessionId, err := context.Cookie(MercurySessionName)
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
	return
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
	if userSession.IsModify() == false {
		return
	}
	err := userSession.Save()

	if err != nil {
		return
	}
	// reset cookie
	sessionId := userSession.Id()
	cookie := &http.Cookie{
		Name:     MercurySessionName,
		Value:    sessionId,
		Path:     "/",
		MaxAge:   CookieMaxAge,
		HttpOnly: true,
	}
	http.SetCookie(context.Writer, cookie)
}
