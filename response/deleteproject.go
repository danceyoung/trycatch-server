/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-07-23 17:30:08
 * @Last Modified by: Young
 * @Last Modified time: 2018-07-25 10:55:07
 */
package response

import (
	"github.com/danceyoung/trycatchserver/db"
	"github.com/gin-gonic/gin"
)

func DeleteProject(uid, projectId string) map[string]interface{} {
	var response = make(map[string]interface{})
	var delProjectSql = "delete from tt_project where created_by='" + uid + "' and project_id='" + projectId + "' and id<>0"
	execDeleteAction(delProjectSql)
	delProjectSql = "delete from tt_project_member where project_id = '" + projectId + "' and id<>0"
	execDeleteAction(delProjectSql)

	response["msg"] = gin.H{"code": 0, "content": "delete project and members successfully"}
	return response
}

func execDeleteAction(delProjectSql string) {
	stmt, err := db.DB.Prepare(delProjectSql)
	if err == nil {
		_, err1 := stmt.Exec()
		if err1 != nil {
			panic(err1.Error())
		}
	} else {
		panic(err.Error())
	}
}
