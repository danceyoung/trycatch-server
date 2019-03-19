/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-06-22 14:06:13
 * @Last Modified by: Young
 * @Last Modified time: 2018-10-26 14:12:16
 */
package model

import (
	"database/sql"
	"encoding/base64"

	"github.com/danceyoung/trycatchserver/db"
)

type SigninUser struct {
	AccountName string `json:"account_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type UID struct {
	UserId string `json:"uid" binding:"required"`
}

type ChangePassword struct {
	UserId string `json:"uid" binding:"required"`
	Old    string `json:"old" binding:"required"`
	New    string `json:"new" binding:"required"`
}

type User struct {
	UserName, Password string
}

func (user User) IsExisted() (bool, string, string) {
	var (
		userid       string
		scanpassword string
	)
	err := db.DB.QueryRow("select user_id, password from tt_user where account_name = ?", user.UserName).Scan(&userid, &scanpassword)
	if err == sql.ErrNoRows {
		return false, "", ""
	} else if err != nil {
		panic(err)
	} else {
		return true, userid, scanpassword
	}
}

func (user User) AddUser() string {
	base64uid := base64.RawURLEncoding.EncodeToString([]byte(user.UserName))
	_, err := db.DB.Exec("insert into tt_user (`user_id`, `account_name`, `password`) values (?,?,?)", base64uid, user.UserName, user.Password)
	if err == nil {
		return base64uid
	}
	return ""
}
