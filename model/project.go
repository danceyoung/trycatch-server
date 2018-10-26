/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-06-28 15:45:19
 * @Last Modified by: Young
 * @Last Modified time: 2018-07-26 10:27:18
 */
package model

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
