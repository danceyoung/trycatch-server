/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-06-12 11:13:12
 * @Last Modified by: Young
 * @Last Modified time: 2018-06-12 11:30:45
 */
package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/try_catch")
	if err != nil {
		panic("mysql init occurs error " + err.Error())
	}

	DB = db
}
