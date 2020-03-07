package account

import (
	"fmt"
	"yiluhuakai/questions/db"
	"yiluhuakai/questions/gen_id"
	"yiluhuakai/questions/middleware/account"
	"yiluhuakai/questions/model"
	"yiluhuakai/questions/util"

	"github.com/gin-gonic/gin"
)

func RegisterHandle(c *gin.Context) {
	var userInfo model.UserInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	if len(userInfo.Email) == 0 || len(userInfo.Password) == 0 ||
		len(userInfo.Username) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	//sex=1表示男生，sex=2表示女生
	if userInfo.Sex != model.UserSexMan && userInfo.Sex != model.UserSexWomen {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	//produce a user id
	userId, err := gen_id.GetID()
	if err != nil {
		fmt.Println(err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	userInfo.UserId = int64(userId)
	err = db.Register(&userInfo)
	if err == db.ErrUserExists {
		util.ResponseError(c, util.ErrCodeUserExist)
		return
	}
	if err != nil {
		fmt.Println(err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(c, nil)
}

func LoginHandle(c *gin.Context) {
	var userInfo model.UserInfo
	var err error
	// 检测用户是否登录以及是否存在session
	account.ProcessRequest(c)
	defer func() {
		// 可以之前的代码已经报错了，这里还会执行，所以需要检测下err
		if err != nil {
			return
		}
		account.SetUserId(c, userInfo.UserId)
		// 检测session是否改变，session改变重新设置设置cookie
		account.ProcessReponse(c)
		util.ResponseSuccess(c, nil)
	}()
	err = c.BindJSON(&userInfo)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	if len(userInfo.Username) == 0 || len(userInfo.Password) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	err = db.Login(&userInfo)
	if err == db.ErrUserNotExists {
		util.ResponseError(c, util.ErrCodeUserNotExist)
		return
	}
	if err == db.ErrUserPasswordWrong {
		util.ResponseError(c, util.ErrCodeUserPasswordWrong)
		return
	}

	if err != nil {
		fmt.Println(err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

}
