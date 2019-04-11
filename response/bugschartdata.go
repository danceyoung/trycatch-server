/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2019-04-03 11:49:21
 * @Last Modified by: Young
 * @Last Modified time: 2019-04-08 11:12:23
 */
package response

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danceyoung/trycatchserver/db"
	"github.com/gin-gonic/gin"
)

func BugsChartData(uids []string, projectId string) map[string]interface{} {
	var response = make(map[string]interface{})
	var chartData []int
	var querySql string
	var uidsInSql = "'" + strings.Join(uids, "','") + "'"
	for i := 24; i >= 1; i-- {
		querySql = querySql + "  select count(*) as y_value from tt_catch_info where timestampdiff(hour, from_unixtime(log_timestamp/1000), now()) =" + strconv.Itoa(i) + "  and user_id in (" + uidsInSql + ") and project_id ='" + projectId + "'   union all"
	}
	querySql = strings.TrimSuffix(querySql, "union all")
	fmt.Println(querySql)
	rows, err := db.DB.Query(querySql)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		var yValue int
		if err := rows.Scan(&yValue); err != nil {
			panic(err.Error())
		}
		chartData = append(chartData, yValue)
	}
	response["chart"] = chartData
	response["msg"] = gin.H{"code": 0, "content": "fetch bugs chartdata successfully"}
	return response
}
