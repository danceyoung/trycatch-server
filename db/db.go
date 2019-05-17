/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-06-12 11:13:12
 * @Last Modified by: Young
 * @Last Modified time: 2019-04-02 14:58:01
 */
package db

import (
	"database/sql"

	"github.com/danceyoung/trycatchserver/constant"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var dataSourceName = "root:root@tcp(127.0.0.1:3306)/try_catch"
	if constant.DEBUG == false {
		dataSourceName = "root:Try1@@300@tcp(127.0.0.1:3306)/try_catch_db"
	}
	//39.105.65.207
	//127.0.0.1
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic("mysql init occurs error " + err.Error())
	}

	DB = db
}
