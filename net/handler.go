/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-06-28 15:12:45
 * @Last Modified by: Young
 * @Last Modified time: 2019-04-03 15:10:55
 */
package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/danceyoung/trycatchserver/constant"
	"github.com/danceyoung/trycatchserver/model"
	"github.com/danceyoung/trycatchserver/response"
	"github.com/danceyoung/trycatchserver/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Signin(c *gin.Context) {
	var user model.SigninUser
	err := c.ShouldBindWith(&user, binding.JSON)
	if err == nil {
		result := response.Signin(user.AccountName, user.Password)
		c.JSON(200, result)
	}
}

func SigninFromMobile(c *gin.Context) {
	var user model.SigninUser
	err := c.ShouldBindWith(&user, binding.JSON)
	if err == nil {
		result := response.SigninFromMobile(user.AccountName, user.Password)
		c.JSON(200, result)
	}
}

func Profile(c *gin.Context) {
	var uid model.UID
	err := c.ShouldBindWith(&uid, binding.JSON)
	if err == nil {
		result := response.Profile(uid.UserId)
		c.JSON(200, result)
	}
}

func ChangePassword(c *gin.Context) {
	var cp model.ChangePassword
	err := c.ShouldBindWith(&cp, binding.JSON)
	if err == nil {
		result := response.ChangePassword(cp.UserId, cp.Old, cp.New)
		c.JSON(200, result)
	}
}

func Projects(c *gin.Context) {
	var uid model.UID
	err := c.ShouldBindWith(&uid, binding.JSON)
	if err == nil {
		result := response.Projects(uid.UserId)
		c.JSON(200, result)
	}
}

func ReceiveFromList(c *gin.Context) {
	var rfl model.ReceiveFromList
	err := c.ShouldBindWith(&rfl, binding.JSON)
	if err == nil {
		result := response.ReceiveFromList(rfl.UserId, rfl.ProjectId)
		c.JSON(200, result)
	}
}

func NewProject(c *gin.Context) {
	var project model.Project
	err := c.ShouldBindWith(&project, binding.JSON)
	if err == nil {
		result := response.NewProject(project)
		c.JSON(200, result)
	}
}

func ProjectDetail(c *gin.Context) {
	var detail model.ReceiveFromList
	err := c.ShouldBindWith(&detail, binding.JSON)
	if err == nil {
		c.JSON(200, response.ProjectDetail(detail.UserId, detail.ProjectId))
	}
}

func DeleteProject(c *gin.Context) {
	var delete model.ReceiveFromList
	err := c.ShouldBindWith(&delete, binding.JSON)
	if err == nil {
		c.JSON(200, response.DeleteProject(delete.UserId, delete.ProjectId))
	}
}

func SaveProject(c *gin.Context) {
	var editproject model.EditProject
	err := c.ShouldBindWith(&editproject, binding.JSON)
	if err == nil {
		c.JSON(200, response.SaveProject(editproject))
	}
}

func Bugs(c *gin.Context) {
	var req model.ProjectBugRequest
	err := c.ShouldBindWith(&req, binding.JSON)
	if err == nil {
		c.JSON(200, response.Bugs(req))
	}
}

func BugsChart(c *gin.Context) {
	var req model.ProjectBugsChartRequest
	err := c.ShouldBindWith(&req, binding.JSON)
	if err == nil {
		c.JSON(200, response.BugsChartData(req.DebuggerIds, req.ProjectId))
	}
}

func TryCatch(c *gin.Context) {
	reqbody, _ := ioutil.ReadAll(c.Request.Body)

	var reqbodystr = string(reqbody)
	fmt.Println(reqbodystr)

	var ttftokenIdx = strings.Index(reqbodystr, constant.TtfAcessToken)
	if ttftokenIdx == -1 {
		return
	}
	var startIdx int = utils.IndexOf(reqbodystr, constant.TtfJsonPrefix, ttftokenIdx, constant.Forward)
	var endIdx int = utils.IndexOf(reqbodystr, constant.TtfJsonSuffix, ttftokenIdx, constant.Backward)
	if startIdx == -1 || endIdx == -1 {
		return
	}
	var ttfJsonStr = reqbodystr[startIdx : endIdx+1]
	fmt.Println(ttfJsonStr)
	if !strings.Contains(ttfJsonStr, constant.TtfAcessToken) {
		return
	}

	var header model.Header
	err := json.Unmarshal([]byte(ttfJsonStr), &header)

	if err == nil {
		c.JSON(200, response.TryCatch(header, strings.Replace(reqbodystr, ttfJsonStr, "", -1)))
		fmt.Println("diapering ttttoken is " + header.TtfAccessToken + " diapering content are " + string(reqbody))
	}
}

func DeviceToken(c *gin.Context) {
	var req model.DeviceToken
	err := c.ShouldBindWith(&req, binding.JSON)
	if err == nil {
		c.JSON(200, response.CollectDeviceToken(req))
	}
}

func XgPushTest(c *gin.Context) {
	// response.XgPush("cb5367ae2812c827421278dd6c112331d0f7f5f4c42e1d4edc6b3523258e0c46", "YoungDing >_ TestProject", "github.com/danceyoung/trycatchserver/netnet/handler.go:159:17: not enough arguments in call to response.XgPushhave ()want (string, string, string)")
	c.JSON(200, "123")
}
