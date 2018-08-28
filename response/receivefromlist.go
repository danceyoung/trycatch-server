/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-07-17 10:40:28
 * @Last Modified by: Young
 * @Last Modified time: 2018-07-24 15:01:07
 */
package response

import (
	"strings"

	"github.com/danceyoung/trycatchserver/db"
	"github.com/gin-gonic/gin"
)

func ReceiveFromList(uid, projectid string) map[string]interface{} {
	var response = make(map[string]interface{})
	var accountsSql = "select receive_from_list from tt_project_member where project_id = ? and user_id=?"
	var accounts string
	accountsErr := db.DB.QueryRow(accountsSql, projectid, uid).Scan(&accounts)
	if accountsErr == nil {
		var inArgs = "'" + strings.Replace(accounts, ",", "','", -1) + "'"
		var receiveFromListSql = "SELECT tt_user.user_id,tt_user.account_name,tt_project_member.user_alias FROM tt_user,tt_project_member WHERE account_name IN (" + inArgs + ") AND tt_user.user_id = tt_project_member.user_id AND tt_project_member.project_id =?"
		receiveFromList, err := db.DB.Query(receiveFromListSql, projectid)
		defer receiveFromList.Close()
		if err != nil {
			panic(err.Error())
		} else {
			var receiveFromListArray []interface{}
			for receiveFromList.Next() {
				var member = make(map[string]interface{})
				var (
					user_id, email, alias string
				)
				if errOne := receiveFromList.Scan(&user_id, &email, &alias); errOne == nil {
					member["user_id"] = user_id
					member["email"] = email
					member["alias"] = alias
					receiveFromListArray = append(receiveFromListArray, member)
				} else {
					panic(errOne.Error())
				}
			}

			response["receive_from_list"] = receiveFromListArray
			response["msg"] = gin.H{"code": 0, "content": "fetch receivefromlist successfully"}
		}

		if errP := receiveFromList.Err(); errP != nil {
			panic(errP.Error())
		}

	} else {
		panic(accountsErr.Error())
	}
	return response
}
