/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2019-03-19 14:16:38
 * @Last Modified by: Young
 * @Last Modified time: 2019-03-19 15:15:08
 */
package model

import (
	"database/sql"

	"github.com/danceyoung/trycatchserver/db"
)

type DeviceToken struct {
	UID         string `json:"uid" binding:"required"`
	AccountName string `json:"account_name" binding:"required"`
	DeviceToken string `json:"device_token" binding:"required"`
}

func (dt DeviceToken) Checking() (new, modified bool) {
	new, modified = false, false
	var token string
	err := db.DB.QueryRow("SELECT device_token FROM tt_device_token where user_id = ?", dt.UID).Scan(&token)
	switch {
	case err == sql.ErrNoRows:
		new = true
		modified = false
	case err != nil:
		panic(err.Error())
	default:
		if dt.DeviceToken != token {
			new = false
			modified = true
		}
	}
	return
}
