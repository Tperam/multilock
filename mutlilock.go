/*
 * @Author: Tperam
 * @Date: 2022-04-23 22:09:19
 * @LastEditTime: 2022-04-23 22:45:01
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\mutlilock.go
 */
package multilock

import (
	"sync"
)

type Multilock struct {
	lock sync.Mutex
	m    map[string]*sync.Mutex
}

// estimate 预估锁大小，减少分配。
func NewMultilock(estimate int) *Multilock {
	return &Multilock{
		lock: sync.Mutex{},
		m:    make(map[string]*sync.Mutex, estimate),
	}
}

type ExecFunc func() (interface{}, error)

func (m *Multilock) Do(f ExecFunc, lockName ...string) (interface{}, error) {
	lockList := make([]*sync.Mutex, 0, len(lockName))

	// 多重锁排序，防止出现死锁
	// N^2 ，通常此处不会出现过多的锁
	for i := 0; i < len(lockName); i++ {
		for j := i + 1; j < len(lockName); j++ {
			if lockName[i] == lockName[j] {
				lockName[j] = lockName[len(lockName)-1]
				lockName = lockName[:len(lockName)-1]
				j = i + 1
			}
			if lockName[i] > lockName[j] {
				// 交换
				lockName[j], lockName[i] = lockName[i], lockName[j]
			}

		}
	}

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
		if lock, ok = m.m[lockName[i]]; !ok {
			m.lock.Lock()
			if lock, ok = m.m[lockName[i]]; !ok {
				lock = &sync.Mutex{}
				m.m[lockName[i]] = lock
			}
			m.lock.Unlock()
		}
		lockList = append(lockList, lock)
		lock.Lock()
	}
	return f()
}
