/*
 * @Author: Tperam
 * @Date: 2022-04-23 22:09:19
 * @LastEditTime: 2022-04-23 22:52:15
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\mutlilock.go
 */
package multilock

import (
	"multilock/algorithm"
	"sync"
)

type Multilock struct {
	RWlock sync.RWMutex
	m      map[string]*sync.Mutex
}

// estimate 预估锁大小，减少分配。
func NewMultilock(estimate int) *Multilock {
	return &Multilock{
		RWlock: sync.RWMutex{},
		m:      make(map[string]*sync.Mutex, estimate),
	}
}

type ExecFunc func() (interface{}, error)

func (m *Multilock) Do(f ExecFunc, lockName ...string) (interface{}, error) {
	lockList := make([]*sync.Mutex, 0, len(lockName))

	// 多重锁排序，防止出现死锁
	lockName = algorithm.Bubblesort(lockName)

	defer func() {
		for i := range lockList {
			lockList[i].Unlock()
		}
	}()
	// 上锁
	// 循环需要锁的Key
	for i := range lockName {
		var lock *sync.Mutex
		var ok bool
		// 判断需要锁的key是否被初始化
		// double check
		// 当前锁粒度过高，待优化。
		m.RWlock.RLock()
		lock, ok = m.m[lockName[i]]
		m.RWlock.RUnlock()
		if !ok {
			m.RWlock.Lock()
			if lock, ok = m.m[lockName[i]]; !ok {
				lock = &sync.Mutex{}
				m.m[lockName[i]] = lock
			}
			m.RWlock.Unlock()
		}
		lockList = append(lockList, lock)
		lock.Lock()
	}
	return f()
}
