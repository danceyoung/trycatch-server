/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-06-28 15:47:58
 * @Last Modified by: Young
 * @Last Modified time: 2018-07-26 10:33:08
 */
package response

import (
	"database/sql"
	"encoding/base64"

	"github.com/danceyoung/trycatchserver/db"
	"github.com/danceyoung/trycatchserver/model"
	"github.com/danceyoung/trycatchserver/utils"
	"github.com/gin-gonic/gin"
)

func NewProject(project model.Project) map[string]interface{} {
	//check uinque project name in created_by
	var existProjectName string
	err := db.DB.QueryRow("select project_name from tt_project where project_name=? and created_by=?", project.ProjectName, project.CreatedBy).Scan(&existProjectName)
	if err != sql.ErrNoRows {
		return gin.H{"msg": gin.H{"code": 11, "content": "The project name already exists on your account."}}
	}
	var base64ProjectID = base64.RawURLEncoding.EncodeToString([]byte(project.ProjectName + project.CreatedBy))

	var insertProjectSql = "insert into tt_project (`project_id`,`project_name`,`language`,`created_by`) values (?,?,?,?)"
	_, errProject := db.DB.Exec(insertProjectSql, base64ProjectID, project.ProjectName, project.Language, project.CreatedBy)

	if errProject == nil {
		var genmembers []model.EditMember
		for i := 0; i < len(project.Members); i++ {
			member := project.Members[i]
			genmembers = append(genmembers, model.EditMember{UserId: "", Email: member.Email, Alias: member.Alias, ReceiveFromList: member.ReceiveFromList})
		}
		return InsertMemberAction(genmembers, base64ProjectID)
	} else {
		panic(errProject.Error())
	}
}

func InsertMemberAction(members []model.EditMember, projectID string) map[string]interface{} {
	var insertMemberSql string = "insert into tt_project_member (`project_id`,`user_id`,`user_alias`,`receive_from_list`) values"
	for i := 0; i < len(members); i++ {
		user := SaveProjectMember(members[i].Email, utils.InitialPasswordMD5Value())
		var uid string
		switch user["uid"].(type) {
		case string:
			uid = user["uid"].(string)
		}

		var splitStr string
		if i != len(members)-1 {
			splitStr = ","
		} else {
			splitStr = ";"
		}

		insertMemberSql += ("('" + projectID + "','" + uid + "','" + members[i].Alias + "','" + members[i].ReceiveFromList + "')" + splitStr)
	}
	_, errMember := db.DB.Exec(insertMemberSql)
	if errMember == nil {
		return gin.H{"msg": gin.H{"code": 0, "content": "new project succeessfully."}}
	} else {
		panic(errMember.Error())
	}
}
