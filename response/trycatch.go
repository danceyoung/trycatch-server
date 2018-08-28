/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-08-08 09:38:15
 * @Last Modified by: Young
 * @Last Modified time: 2018-08-27 16:08:29
 */
package response

import (
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
		stmt.Exec(userid, projectid, info)
	} else {
		panic(err.Error())
	}
	return gin.H{"msg": gin.H{"code": 0, "content": "try a catch succeessfully."}}
}
