/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-08-08 09:38:15
 * @Last Modified by: Young
 * @Last Modified time: 2019-03-25 15:51:39
 */
package response

import (
	"fmt"
	"strings"

	"github.com/danceyoung/trycatchserver/constant"
	"github.com/danceyoung/trycatchserver/db"
	"github.com/danceyoung/trycatchserver/model"
	"github.com/gin-gonic/gin"
)

func TryCatch(header model.Header, info string) map[string]interface{} {
	var tokens = strings.Split(header.Ttftoken, constant.TtfJsonTokenSplit)
	if len(tokens) != 2 {
		return nil
	}
	var userid = tokens[0]
	var projectid = tokens[1]

	var projectidvar string
	err := db.DB.QueryRow(`select tt_project.project_id from tt_project, tt_project_member
	 where tt_project.project_id=tt_project_member.project_id 
	 and tt_project.project_id=? and user_id=?`, projectid, userid).Scan(&projectidvar)
	if err != nil {
		return nil
	}

	stmt, err := db.DB.Prepare("insert into tt_catch_info (`user_id`,`project_id`,`catch_info`) values (?,?,?)")
	defer stmt.Close()
	if err == nil {
		_, errIns := stmt.Exec(userid, projectid, info)
		if errIns == nil {
			var uidsArray []string
			var currentAccountName string
			errCAN := db.DB.QueryRow("select account_name from tt_user where user_id = ?", userid).Scan(&currentAccountName)
			if errCAN != nil {
				panic(errCAN.Error())
			} else {
				sql := "select user_id from tt_project_member where receive_from_list like '%?%'"
				sql = strings.Replace(sql, "?", currentAccountName, -1)
				fmt.Println(sql)
				rows, errUids := db.DB.Query(sql)
				if errUids != nil {
					panic(errUids.Error())
				}
				defer rows.Close()
				for rows.Next() {
					var tempUid string
					if errTempUid := rows.Scan(&tempUid); errTempUid != nil {
						panic(errTempUid.Error())
					}
					uidsArray = append(uidsArray, tempUid)
				}
			}
			var debuger = model.ProjectMemberAlias(userid)
			var projectName = model.ProjectName(projectid)
			var title = debuger + " " + projectName
			var deviceTokens []string
			for i := 0; i < len(uidsArray); i++ {
				var userid = uidsArray[i]
				var deviceToken = model.DeviceTokenByUserId(userid)
				deviceTokens = append(deviceTokens, deviceToken)
			}
			XgPush(deviceTokens, title, info)

		} else {
			panic(errIns.Error())
		}
	} else {
		panic(err.Error())
	}
	return gin.H{"msg": gin.H{"code": 0, "content": "try a catch succeessfully."}}
}
