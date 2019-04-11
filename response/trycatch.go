/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-08-08 09:38:15
 * @Last Modified by: Young
 * @Last Modified time: 2019-04-03 09:35:34
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
	var tokens = strings.Split(header.TtfAccessToken, constant.TtfJsonTokenSplit)
	if len(tokens) != 2 {
		panic("ttf_access_token is not enough content.")
	}
	var userid = tokens[0]
	var projectid = tokens[1]

	var projectidvar string
	err := db.DB.QueryRow(`select tt_project.project_id from tt_project, tt_project_member
	 where tt_project.project_id=tt_project_member.project_id 
	 and tt_project.project_id=? and user_id=?`, projectid, userid).Scan(&projectidvar)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.DB.Prepare("insert into tt_catch_info (`user_id`,`project_id`,`catch_info`,`log_timestamp`) values (?,?,?,?)")
	defer stmt.Close()
	if err == nil {
		_, errIns := stmt.Exec(userid, projectid, info, header.TtfLogTimestamp)
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

			var deviceTokenObjects = model.DeviceTokenObjectsByUserIds(uidsArray)
			deviceTokenObjects = checkDeviceTokenObjects(deviceTokenObjects)

			var deviceTokens []string
			var updateUids []string
			for i := 0; i < len(deviceTokenObjects); i++ {
				dto := deviceTokenObjects[i]
				updateUids = append(updateUids, dto.UID)
				deviceTokens = append(deviceTokens, dto.DeviceToken)
			}
			XgPush(deviceTokens, title, info)
			// model.UpdateXgPushState(updateUids)
		} else {
			panic(errIns.Error())
		}
	} else {
		panic(err.Error())
	}
	return gin.H{"msg": gin.H{"code": 0, "content": "try a catch succeessfully."}}
}
