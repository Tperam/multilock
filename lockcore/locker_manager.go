/*
 * @Author: Tperam
 * @Date: 2022-04-25 22:17:14
 * @LastEditTime: 2022-04-25 22:33:38
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\lockcore\locker_manager.go
 */
package lockcore

import "github.com/tperam/multilock/locker"

type LockCore interface {
	GetLock(string) (locker.Locker, error)
}
