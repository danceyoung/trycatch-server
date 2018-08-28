/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-07-25 14:53:27
 * @Last Modified by: Young
 * @Last Modified time: 2018-07-26 10:38:08
 */
package response

import (
	"strings"

	"github.com/danceyoung/trycatchserver/db"
	"github.com/danceyoung/trycatchserver/model"
)

func SaveProject(project model.EditProject) map[string]interface{} {
	//save project ,delete members,update members,insert members
	var updateProjectSql = "update tt_project set project_name=?,language=? where id<>0 and project_id=? and created_by=?"
	stmt, err := db.DB.Prepare(updateProjectSql)
	if err == nil {
		_, err1 := stmt.Exec(project.ProjectName, project.Language, project.ProjectId, project.UId)
		if err1 != nil {
			panic(err1.Error())
		}
	} else {
		panic(err.Error())
	}

	if len(project.DeleteMembers) > 0 {
		var emails = strings.Replace(project.DeleteMembers, ",", "','", -1)
		emails = "('" + emails + "')"
		var deleteMembersSql = "delete from tt_project_member where project_id=? and user_id in (select user_id from tt_user where account_name in " + emails + ")"
		stmtd, errd := db.DB.Prepare(deleteMembersSql)
		if errd == nil {
			_, err1 := stmtd.Exec(project.ProjectId)
			if err1 != nil {
				panic(err1.Error())
			}
		} else {
			panic(errd.Error())
		}
	}

	var genmembers []model.EditMember

	for i := 0; i < len(project.Members); i++ {
		member := project.Members[i]
		if len(member.UserId) > 0 {
			var updateMemberSql = "update tt_project_member set receive_from_list='" + member.ReceiveFromList + "' where project_id='" + project.ProjectId + "' and user_id='" + member.UserId + "'"
			_, err := db.DB.Exec(updateMemberSql)
			if err != nil {
				panic(err.Error())
			}
		} else {
			genmembers = append(genmembers, member)
		}
	}

	return InsertMemberAction(genmembers, project.ProjectId)
}
