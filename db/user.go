package db

import (
	"database/sql"
	"logger"

	"github.com/jmoiron/sqlx"
	"questions/model"
	"questions/util"
)

const (
	PasswordSalt = "HBZciU2SiSDr4uPeJ1e7qlIlMbyusQ0v"
)

func Register(user *model.UserInfo) (err error) {
	var userId int64
	sqlStr := "select user_id from user where username=?"
	err = db.Get(&userId, sqlStr, user.Username)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	if userId > 0 {
		err = ErrUserExists
		return
	}
	// 给密码加密
	user.Password = util.MD5([]byte(PasswordSalt + user.Password))
	// 用户不存在，可以注册
	sqlStr = "insert into user(user_id,nickname,sex,username,password,email) values(?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, user.UserId, user.Nickname, user.Sex, user.Username, user.Password, user.Email)
	return
}

func Login(userInfo *model.UserInfo) (err error) {
	password := userInfo.Password
	sqlStr := "select username,password,user_id from user where username =?"
	err = db.Get(userInfo, sqlStr, userInfo.Username)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if err == sql.ErrNoRows {
		err = ErrUserNotExists
		return
	}

	//由于密码是加密过的，所以这块也要加密
	if userInfo.Password != util.MD5([]byte(PasswordSalt+password)) {
		err = ErrUserPasswordWrong
	}
	return
}

func GetUserInfoList(userIdList []int64) (userInfoList []*model.UserInfo, err error) {

	if len(userIdList) == 0 {
		return
	}
	sqlStr := `select
					user_id, nickname, sex, username, email
				from
					user
				where user_id in(?)`
	var tempUserIdList []interface{}

	for _, userId := range userIdList {
		tempUserIdList = append(tempUserIdList, userId)
	}
	query, args, err := sqlx.In(sqlStr, tempUserIdList)

	if err != nil {
		logger.LogError("sql err:%v", err)
		return

	}
	err = db.Select(&userInfoList, query, args...)
	if err != nil {
		logger.LogError("select userinfo failed :%v", err)
	}
	return
}
