/*
 * @Author: Tperam
 * @Date: 2022-04-24 23:37:47
 * @LastEditTime: 2022-04-24 23:43:10
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\locker\errors.go
 */
package locker

import "errors"

var (
	ErrRepeatUnlocked = errors.New("unlock of unlocked mutex")
)
