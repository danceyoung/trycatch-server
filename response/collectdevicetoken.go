/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2019-03-19 15:08:04
 * @Last Modified by: Young
 * @Last Modified time: 2019-03-19 15:38:16
 */
package response

import (
	"github.com/danceyoung/trycatchserver/db"
	"github.com/danceyoung/trycatchserver/model"
	"github.com/gin-gonic/gin"
)

func CollectDeviceToken(dt model.DeviceToken) map[string]interface{} {
	new, modified := dt.Checking()
	response := make(map[string]interface{})
	if new == true {
		_, errNew := db.DB.Exec("INSERT INTO tt_device_token (`user_id`,`account_name`,`device_token`) VALUES (?,?,?)", dt.UID, dt.AccountName, dt.DeviceToken)
		if errNew != nil {
			panic(errNew.Error())
		}
		response["msg"] = gin.H{"code": 0, "content": "new a token"}
	} else if modified == true {
		_, errUpdate := db.DB.Exec("UPDATE tt_device_token SET `device_token` = ? WHERE `user_id` = ?", dt.DeviceToken, dt.UID)
		if errUpdate != nil {
			panic(errUpdate.Error())
		}
		response["msg"] = gin.H{"code": 0, "content": "update a token"}
	}
	return response
}
