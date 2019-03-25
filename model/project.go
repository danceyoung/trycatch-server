/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-06-28 15:45:19
 * @Last Modified by: Young
 * @Last Modified time: 2019-03-22 16:19:27
 */
package model

import (
	"database/sql"

	"github.com/danceyoung/trycatchserver/db"
)

type member struct {
	Email           string `json:"email" binding:"required"`
	Alias           string `json:"alias" binding:"required"`
	ReceiveFromList string `json:"receive_from_list" binding:"required"`
}

type Project struct {
	ProjectName string   `json:"project_name" binding:"required"`
	Language    string   `json:"language" binding:"required"`
	CreatedBy   string   `json:"created_by" binding:"required"`
	Members     []member `json:"members" binding:"required"`
}

type EditProject struct {
	UId           string       `json:"uid" binding:"required"`
	ProjectId     string       `json:"project_id" binding:"required"`
	ProjectName   string       `json:"project_name" binding:"required"`
	Language      string       `json:"language" binding:"required"`
	Members       []EditMember `json:"members" binding:"required"`
	DeleteMembers string       `json:"delete_members"`
}

type EditMember struct {
	UserId          string `json:"user_id" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Alias           string `json:"alias" binding:"required"`
	ReceiveFromList string `json:"receive_from_list" binding:"required"`
}

type ReceiveFromList struct {
	UserId    string `json:"uid" binding:"required"`
	ProjectId string `json:"project_id" binding:"required"`
}

type ProjectBugRequest struct {
	UserId      string   `json:"uid" binding:"required"`
	ProjectId   string   `json:"project_id" binding:"required"`
	DebuggerIds []string `json:"debugger_ids" binding:"required"`
	FetchPage   int      `json:"fetch_page" binding:"required"`
}

func ProjectMemberAlias(uid string) string {
	var aliasStr string
	err := db.DB.QueryRow("SELECT user_alias FROM tt_project_member where user_id = ?", uid).Scan(&aliasStr)
	switch {
	case err == sql.ErrNoRows:
		return ""
	case err != nil:
		panic(err.Error())
	default:
		return aliasStr
	}
}

func ProjectName(projectId string) string {
	var nameStr string
	err := db.DB.QueryRow("SELECT project_name FROM tt_project where project_id = ?", projectId).Scan(&nameStr)
	switch {
	case err == sql.ErrNoRows:
		return ""
	case err != nil:
		panic(err.Error())
	default:
		return nameStr
	}
}
