/*
 * @Author: Tperam
 * @Date: 2022-04-24 23:25:23
 * @LastEditTime: 2022-04-24 23:34:02
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\multilocker\multilocker.go
 */
package multilocker

import (
	"multilock/locker"
	"sync"
)

// 需要实现 sync.Locker
type MapLock struct {
	RWlock       sync.RWMutex
	m            map[string]locker.Locker
	generateLock locker.GenerateLocker
}

//
func (ml *MapLock) Lock(lockName string) (locker.Locker, error) {
	var err error
	ml.RWlock.RLock()
	lock, ok := ml.m[lockName]
	ml.RWlock.RUnlock()
	if !ok {
		ml.RWlock.Lock()
		if lock, ok = ml.m[lockName]; !ok {
			lock, err = ml.generateLock.New(lockName)
			ml.m[lockName] = lock
		}
		ml.RWlock.Unlock()
	}

	lock.Lock()

	return lock, err
}
