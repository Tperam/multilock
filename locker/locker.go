/*
 * @Author: Tperam
 * @Date: 2022-04-24 23:09:53
 * @LastEditTime: 2022-04-24 23:31:28
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\locker\locker.go
 */
package locker

// 支持分布式锁
type Locker interface {
	Lock() error
	Unlock() error
}

type GenerateLocker interface {
	New(lockname string) (Locker, error)
}
