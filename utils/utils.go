/*
 * @Author: Young
 * DSHARP
 * @flow
 * @Date: 2018-07-03 10:36:40
 * @Last Modified by: Young
 * @Last Modified time: 2018-08-27 16:02:38
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/danceyoung/trycatchserver/constant"
)

func InitialPasswordMD5Value() string {
	md5value := md5.Sum([]byte(constant.InitialPassword))
	return hex.EncodeToString(md5value[:])
}

func IndexOf(s, substr string, fromIndex int, direct int) int {
	var index = -1

	if direct == constant.Forward {
		for i := fromIndex; i > -1; i-- {
			if s[i-1:i] == substr {
				index = i - 1
				break
			}
		}
	} else {
		for i := fromIndex + 1; i < len(s); i++ {
			if s[i:i+1] == substr {
				index = i
				break
			}
		}
	}

	return index
}
