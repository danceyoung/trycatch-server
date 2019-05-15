/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-07-02 15:53:17
 * @Last Modified by: Young
 * @Last Modified time: 2018-07-24 14:55:59
 */
package response

import (
	"database/sql"
	"strings"

	"github.com/danceyoung/trycatchserver/db"
	"github.com/gin-gonic/gin"
)

func Projects(uid string) map[string]interface{} {
	var response = make(map[string]interface{})

	var referProjectIDsSql = "SELECT project_id FROM tt_project_member WHERE user_id = ? AND project_id NOT IN (SELECT project_id FROM tt_project WHERE created_by = ?)"
	rrows, rerr := db.DB.Query(referProjectIDsSql, uid, uid)
	defer rrows.Close()
	var pids string
	var pidArray []string
	if rerr != nil {
		panic(rerr.Error())
	} else {
		var pid string

		for rrows.Next() {
			e := rrows.Scan(&pid)
			if e == nil {
				pidArray = append(pidArray, pid)
			}
		}
		pids = "'" + strings.Join(pidArray, "','") + "'"
	}

	var projectArray []interface{}

	var creatorProjectsSql = "SELECT project_id, project_name, language, created_by,create_date,(select count(*) as members from tt_project_member where project_id = tt_project.project_id ) as members FROM tt_project where created_by=? order by create_date desc"
	projectArray = append(projectArray, createProjectArray(creatorProjectsSql, uid, 0)...)
	var referProjectsSql = "SELECT project_id, project_name, language, created_by,create_date,(select count(*) as members from tt_project_member where project_id = tt_project.project_id ) as members FROM tt_project where project_id in (?) order by create_date desc"
	referProjectsSql = strings.Replace(referProjectsSql, "?", pids, -1)
	// // fmt.Println(referProjectsSql)
	projectArray = append(projectArray, createProjectArray(referProjectsSql, "", 1)...)

	if len(projectArray) == 0 {
		response["msg"] = gin.H{"code": 12, "content": "No projects, please new one."}
	} else {
		response["projects"] = projectArray
		response["msg"] = gin.H{"code": 0, "content": "more projects in this account"}
	}
	return response
}

func createProjectArray(querySql string, params string, creator int) []interface{} {
	var (
		id, name, language, createdby, createdate string
		count                                     int
	)
	// var tempParams interface{}
	// if len(params) > 0 {
	// 	tempParams = params
	// } else {
	// 	tempParams = nil
	// }
	var rows *sql.Rows
	var err error
	if len(params) > 0 {
		rows, err = db.DB.Query(querySql, params)
	} else {
		rows, err = db.DB.Query(querySql)
	}
	defer rows.Close()
	var projectsAry = []interface{}{}
	if err != nil {
		panic(err.Error())
	} else {
		for rows.Next() {
			err1 := rows.Scan(&id, &name, &language, &createdby, &createdate, &count)
			if err1 == nil {
				var project = make(map[string]interface{})
				project["project_id"] = id
				project["project_name"] = name
				project["source_code"] = language
				project["created_by"] = createdby
				project["create_date"] = createdate
				project["members"] = count
				project["creator"] = creator
				projectsAry = append(projectsAry, project)
			}
		}
		if err2 := rows.Err(); err2 != nil {
			panic(err2.Error())
		}
	}
	return projectsAry
}
