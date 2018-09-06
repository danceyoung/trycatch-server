/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-08-23 14:41:44
 * @Last Modified by: Young
 * @Last Modified time: 2018-08-23 17:28:49
 */
package response

import (
	"fmt"
	"log"
	"strings"

	"github.com/danceyoung/trycatchserver/db"
	"github.com/danceyoung/trycatchserver/model"
	"github.com/gin-gonic/gin"
)

func Bugs(req model.ProjectBugRequest) map[string]interface{} {
	var response = make(map[string]interface{})

	var rflSql = "select receive_from_list from tt_project_member where user_id=? and project_id=?"
	var receivelist string
	err := db.DB.QueryRow(rflSql, req.UserId, req.ProjectId).Scan(&receivelist)
	if err != nil {
		panic(err.Error())
	} else {
		userids := "'" + strings.Join(req.DebuggerIds, "','") + "'"
		rows, err := db.DB.Query("select tt_catch_info.user_id,tt_project_member.user_alias,catch_info from tt_catch_info,tt_project_member where tt_catch_info.user_id =tt_project_member.user_id and  tt_catch_info.project_id =tt_project_member.project_id and tt_catch_info.user_id in (" + userids + ") and tt_catch_info.project_id ='" + req.ProjectId + "' limit 100")
		fmt.Println("select tt_catch_info.user_id,tt_project_member.user_alias,catch_info from tt_catch_info,tt_project_member where tt_catch_info.user_id =tt_project_member.user_id and  tt_catch_info.project_id =tt_project_member.project_id and tt_catch_info.user_id in (" + userids + ") and tt_catch_info.project_id ='" + req.ProjectId + "' limit 100")
		defer rows.Close()
		if err != nil {
			panic(err.Error())
		} else {
			var bugsAry []map[string]string
			for rows.Next() {
				var userId, alias, content string
				if err := rows.Scan(&userId, &alias, &content); err != nil {
					panic(err.Error())
				} else {
					var bug = make(map[string]string)
					bug["user_id"] = userId
					bug["alias"] = alias
					bug["content"] = content
					bugsAry = append(bugsAry, bug)
				}

			}
			response["bugs"] = bugsAry
			if len(bugsAry) == 0 {
				response["msg"] = gin.H{"code": 12, "content": "no bugs"}
			} else {
				response["msg"] = gin.H{"code": 0, "content": "fetch bugs successfully"}
			}
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return response
}
