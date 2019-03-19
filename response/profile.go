/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-10-26 10:38:25
 * @Last Modified by: Young
 * @Last Modified time: 2018-10-26 11:35:03
 */
package response

import (
	"github.com/danceyoung/trycatchserver/db"
	"github.com/gin-gonic/gin"
)

func Profile(userid string) map[string]interface{} {
	var accountname string
	err := db.DB.QueryRow("select account_name from tt_user where user_id =?", userid).Scan(&accountname)
	if err != nil {
		panic(err.Error())
	} else {
		return gin.H{"account_name": accountname, "msg": gin.H{"code": 0, "content": "profile"}}
	}
}
