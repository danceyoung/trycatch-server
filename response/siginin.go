/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-06-28 14:28:46
 * @Last Modified by: Young
 * @Last Modified time: 2018-07-19 14:49:19
 */
package response

import (
	"database/sql"
	"encoding/base64"

	"github.com/danceyoung/trycatchserver/db"
	"github.com/gin-gonic/gin"
)

func Signin(username, password string) map[string]interface{} {
	return prepareUser(username, password, true)
}

func SaveProjectMember(username, password string) map[string]interface{} {
	return prepareUser(username, password, false)
}

func prepareUser(username, password string, signFlag bool) map[string]interface{} {
	var (
		userid       string
		scanpassword string
	)
	var response = make(map[string]interface{})

	err := db.DB.QueryRow("select user_id, password from tt_user where account_name = ?", username).Scan(&userid, &scanpassword)
	if err == sql.ErrNoRows {
		base64uid := base64.RawURLEncoding.EncodeToString([]byte(username))
		_, err := db.DB.Exec("insert into tt_user (`user_id`, `account_name`, `password`) values (?,?,?)", base64uid, username, password)

		if err == nil {
			response["uid"] = base64uid
			response["msg"] = gin.H{"code": 0, "content": "sign up successfully"}
			return response
		} else {
			panic(err.Error())
		}
	} else if scanpassword != password && signFlag {
		return gin.H{"msg": gin.H{"code": 2, "content": "The account and password are not matching"}}
	} else if err != nil {
		panic(err.Error())
	} else {
		response["uid"] = userid
		response["msg"] = gin.H{"code": 0, "content": "sign in successfully"}
		return response
	}
}
