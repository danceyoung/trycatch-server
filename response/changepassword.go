/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-10-26 14:33:17
 * @Last Modified by: Young
 * @Last Modified time: 2018-10-26 15:00:25
 */
package response

import (
	"database/sql"

	"github.com/danceyoung/trycatchserver/db"
	"github.com/gin-gonic/gin"
)

func ChangePassword(uid, old, new string) map[string]interface{} {
	//check uid and old password are if matching
	var userid string
	err := db.DB.QueryRow("select user_id from tt_user where user_id=? and password=?", uid, old).Scan(&userid)
	if err == sql.ErrNoRows {
		return gin.H{"msg": gin.H{"code": 2, "content": "The old password are incorrect"}}
	} else if err == nil {
		stmt, err := db.DB.Prepare("update tt_user set password=? where user_id=?")
		if err == nil {
			_, err1 := stmt.Exec(new, uid)
			if err1 == nil {
				return gin.H{"msg": gin.H{"code": 0, "content": "Change password successfully."}}
			}
		}
	}
	return nil
}
