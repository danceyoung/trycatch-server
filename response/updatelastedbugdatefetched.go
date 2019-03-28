/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2019-03-28 14:46:11
 * @Last Modified by: Young
 * @Last Modified time: 2019-03-28 15:55:35
 */
package response

import (
	"database/sql"

	"github.com/danceyoung/trycatchserver/db"
)

func UpdateLastedBugDateFetched(uid, projectid string, state int) {
	var stateScan int
	err := db.DB.QueryRow("select lastedbug_id from tt_user_fetch_lastedbugdate_state where user_id = ? and project_id = ?", uid, projectid).Scan(&stateScan)
	switch {
	case err == sql.ErrNoRows:
		//insert
		_, errInsert := db.DB.Exec("insert into tt_user_fetch_lastedbugdate_state (`user_id`,`project_id`,`lastedbug_id`) values (?,?,?)", uid, projectid, state)
		if errInsert != nil {
			panic(errInsert.Error())
		}
	case err == nil:
		//update
		_, errUpdate := db.DB.Exec("update tt_user_fetch_lastedbugdate_state set lastedbug_id = ?", state)
		if errUpdate != nil {
			panic(errUpdate.Error())
		}
	case err != nil:
		panic(err.Error())
	}

}
