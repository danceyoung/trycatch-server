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
	"github.com/danceyoung/trycatchserver/model"
	"github.com/gin-gonic/gin"
)

func Signin(username, password string) map[string]interface{} {
	return SigninOrUp(username, password, true)
}

func SigninFromMobile(username, password string) map[string]interface{} {
	return SigninOrUp(username, password, false)
}

func SigninOrUp(username, password string, inup bool) map[string]interface{} {
	var response = make(map[string]interface{})

	var user = model.User{UserName: username, Password: password}
	var isExisted, userid, scanpassword = user.IsExisted()
	if isExisted == false && inup == true {
		tempUid := user.AddUser()
		response["uid"] = tempUid
		response["msg"] = gin.H{"code": 0, "content": "sign up successfully"}
		return response
	} else {
		if scanpassword != password {
			return gin.H{"msg": gin.H{"code": 2, "content": "The account and password are not matching"}}
		} else {
			response["uid"] = userid
			response["msg"] = gin.H{"code": 0, "content": "sign in successfully"}
			return response
		}
	}
}

func SaveProjectMember(username, password string) map[string]interface{} {
	var response = make(map[string]interface{})

	var user = model.User{UserName: username, Password: password}
	var isExisted, userid, _ = user.IsExisted()
	if isExisted == false {
		tempUid := user.AddUser()
		response["uid"] = tempUid
		response["msg"] = gin.H{"code": 0, "content": "sign up successfully"}
		return response
	} else {
		response["uid"] = userid
		response["msg"] = gin.H{"code": 0, "content": "sign in successfully"}
		return response
	}
}
