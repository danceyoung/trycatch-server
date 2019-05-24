/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2019-03-20 14:20:52
 * @Last Modified by: Young
 * @Last Modified time: 2019-03-25 15:48:53
 */
package response

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/FrontMage/xinge"
	"github.com/FrontMage/xinge/auth"
	"github.com/FrontMage/xinge/req"
	"github.com/danceyoung/trycatchserver/constant"
	"github.com/danceyoung/trycatchserver/model"
)

func XgPush(deviceTokens []string, title, info string) {
	if len(deviceTokens) > 0 {
		pushDebug(deviceTokens, title, info)
	}

}

func pushDebug(deviceTokens []string, title, content string) {
	auther := auth.Auther{AppID: constant.XgAppId, SecretKey: constant.XgSecretKey}
	pushReq, _ := req.NewPushReq(
		&xinge.Request{},
		req.Platform(xinge.PlatformiOS),
		// req.EnvDev(),
		req.EnvProd(),
		req.AudienceType(xinge.AdTokenList),
		req.MessageType(xinge.MsgTypeNotify),
		req.TokenList(deviceTokens),
		req.PushID("0"),
		req.Message(xinge.Message{
			IOS: &xinge.IOSParams{
				Aps: &xinge.Aps{
					Alert: map[string]string{
						"title": title,
						"body":  content,
					},
					Badge: 1,
					Sound: "default",
				},
			},
		}),
	)
	auther.Auth(pushReq)

	c := &http.Client{}
	rsp, _ := c.Do(pushReq)
	defer rsp.Body.Close()
	body, _ := ioutil.ReadAll(rsp.Body)

	r := &xinge.CommonRsp{}
	json.Unmarshal(body, r)
	fmt.Printf("%+v", r)
	if r.RetCode != 0 {
		fmt.Errorf("Failed rsp=%+v", r)
	}
}

func checkDeviceTokenObjects(dtos []model.DeviceToken) []model.DeviceToken {
	var returnedDtos []model.DeviceToken
	for i := 0; i < len(dtos); i++ {
		dto := dtos[i]
		var nowYearDay = time.Now().YearDay()

		lastPushDate, err := time.Parse("2006-01-02 15:04:05", dto.LastPushDate)
		if err != nil {
			panic(err.Error())
		}
		lastPushYearDay := lastPushDate.YearDay()
		if nowYearDay == lastPushYearDay {
			if dto.PushTimes%constant.MaximumXGPushTimes != 0 {
				returnedDtos = append(returnedDtos, dto)
			}
		} else {
			returnedDtos = append(returnedDtos, dto)
		}
	}
	return returnedDtos
}
