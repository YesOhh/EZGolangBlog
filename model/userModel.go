package model

import (
	"database/sql"
	"goBlog/initialization"
	"goBlog/mylog"
	"goBlog/util"
)

type LoginModel struct {
	Email string `form:"email" binding:"email"`
	Password string `form:"password"`
}

type RegisterModel struct {
	Email string `form:"email" binding:"email"`
	Password string `form:"password"`
	PasswordAgain string `form:"passwordAgain"`
	Nickname string `form:"nickname"`
}

type UserModel struct {
	Id int64 `form:"id"`
	Email string `form:"email" binding:"email"`
	Password string `form:"password"`
	Nickname string `form:"nickname"`
}

func (user *UserModel) SaveUser() int64 {
	passwordHash := util.GenStringHash(user.Password)
	result, e := initialization.Db.Exec("INSERT INTO " + initialization.DbName + "(email, password, nickname) values (?, ?, ?);", user.Email, passwordHash, user.Nickname)
	if e != nil {
		mylog.MyLogger.Panicln("user insert error", e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		mylog.MyLogger.Panicln("user insert id error", e.Error())
	}
	return id
}

func (user *UserModel) QueryUserByEmail() UserModel {
	u := UserModel{}
	row := initialization.Db.QueryRow("SELECT id, email, password, nickname FROM " + initialization.DbName + " WHERE email = ?;", user.Email)
	err := row.Scan(&u.Id, &u.Email, &u.Password, &u.Nickname)
	if err != nil {
		if err == sql.ErrNoRows {
			// 没有用户
			return u
		}
		mylog.MyLogger.Panicln("query user err", err.Error())
	}
	return u
}

func (user *UserModel) ValidUser() bool {
	row := initialization.Db.QueryRow("SELECT password FROM " + initialization.DbName + " WHERE email = ?", user.Email)
	var passwordHash string
	err := row.Scan(&passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// 说明没有该用户
			return false
		}
		mylog.MyLogger.Panicln("err: ", err.Error())
	}
	currentPasswordHash := util.GenStringHash(user.Password)
	return currentPasswordHash == passwordHash
}