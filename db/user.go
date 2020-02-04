package db

import (
	"database/sql"
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
