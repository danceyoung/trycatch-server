/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-07-20 16:13:23
 * @Last Modified by: Young
 * @Last Modified time: 2018-07-24 15:41:37
 */
package response

import (
	"github.com/danceyoung/trycatchserver/db"
	"github.com/gin-gonic/gin"
)

func ProjectDetail(uid, projectId string) map[string]interface{} {
	var projectSql = "select project_id,project_name,language,created_by,create_date from tt_project where created_by=? and project_id=?"
	var (
		projectIdVar, projectName, language, createdBy, createdDate string
	)
	err := db.DB.QueryRow(projectSql, uid, projectId).Scan(&projectIdVar, &projectName, &language, &createdBy, &createdDate)
	var project = make(map[string]interface{})
	if err == nil {
		project["project_id"] = projectIdVar
		project["project_name"] = projectName
		project["language"] = language
		project["created_by"] = createdBy
		project["created_date"] = createdDate
	}

	var memberSql = "select tt_project_member.user_id,tt_user.account_name,tt_project_member.user_alias,tt_project_member.receive_from_list from tt_user,tt_project_member where tt_user.user_id=tt_project_member.user_id and tt_project_member.project_id = ?"

	rows, err1 := db.DB.Query(memberSql, projectId)
	if err1 != nil {
		panic(err1.Error())
	}
	defer rows.Close()
	var memberArray []interface{}
	for rows.Next() {
		var (
			userId, accountName, userAlias, receivefromlist string
		)
		if errtemp := rows.Scan(&userId, &accountName, &userAlias, &receivefromlist); errtemp != nil {
			panic(errtemp.Error())
		}
		var member = make(map[string]interface{})
		member["user_id"] = userId
		member["email"] = accountName
		member["alias"] = userAlias
		member["receive_from_list"] = receivefromlist
		memberArray = append(memberArray, member)
	}
	if endErr := rows.Err(); endErr != nil {
		panic(endErr.Error())
	}

	var response = make(map[string]interface{})
	response["project"] = project
	response["members"] = memberArray
	response["msg"] = gin.H{"code": 0, "content": "project detail retrive successfully"}
	return response
}
