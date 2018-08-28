/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-06-22 14:06:13
 * @Last Modified by: Young
 * @Last Modified time: 2018-07-02 15:58:27
 */
package model

type SigninUser struct {
	AccountName string `json:"account_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type UID struct {
	UserId string `json:"uid" binding:"required"`
}
