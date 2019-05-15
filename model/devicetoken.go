/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2019-03-19 14:16:38
 * @Last Modified by: Young
 * @Last Modified time: 2019-03-22 15:32:20
 */
package model

import (
	"database/sql"
	"strings"

	"github.com/danceyoung/trycatchserver/db"
)

type DeviceToken struct {
	UID          string `json:"uid" binding:"required"`
	AccountName  string `json:"account_name" binding:"required"`
	DeviceToken  string `json:"device_token" binding:"required"`
	LastPushDate string
	PushTimes    int
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

func DeviceTokenByUserId(uid string) string {
	var tokenStr string
	err := db.DB.QueryRow("select device_token from tt_device_token where user_id = ?", uid).Scan(&tokenStr)
	switch {
	case err == sql.ErrNoRows:
		return ""
	case err != nil:
		panic(err.Error())
	default:
		return tokenStr
	}
}

func DeviceTokenObjectsByUserIds(uids []string) []DeviceToken {
	var deviceTokenObjects []DeviceToken
	var sql = "select user_id, account_name, device_token from tt_device_token where user_id in ("
	for i := 0; i < len(uids); i++ {
		sql = sql + "'" + uids[i] + "',"

	}
	sql = strings.TrimSuffix(sql, ",") + ")"
	rows, err := db.DB.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var dtobjc = DeviceToken{}
		if err := rows.Scan(&dtobjc.UID, &dtobjc.AccountName, &dtobjc.DeviceToken); err != nil {
			panic(err.Error())
		}
		deviceTokenObjects = append(deviceTokenObjects, dtobjc)
	}

	return deviceTokenObjects
}

func UpdateXgPushState(uids []string) {
	for i := 0; i < len(uids); i++ {
		var updateSql = "update tt_device_token set last_push_date = now(), push_times = push_times+1 where user_id = ?"
		_, err := db.DB.Exec(updateSql, uids[i])
		if err != nil {
			panic(err.Error())
		}
	}
}
