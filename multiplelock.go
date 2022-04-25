/*
 * @Author: Tperam
 * @Date: 2022-04-23 22:09:19
 * @LastEditTime: 2022-04-25 22:34:08
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\multiplelock.go
 */
package multilock

import (
	"multilock/algorithm"
	"multilock/lockcore"
	"multilock/locker"
)

type Multilock struct {
	sorter algorithm.Sorter
	lms    []lockcore.LockCore
}

// estimate 预估锁大小，减少分配。
func NewMultilock(sorter algorithm.Sorter, lms ...lockcore.LockCore) *Multilock {
	return &Multilock{
		sorter: sorter,
		lms:    lms,
	}
}

type ExecFunc func() (interface{}, error)

func (m *Multilock) Do(f ExecFunc, lockNames ...string) (result interface{}, err error) {
	lockList := make([]locker.Locker, 0, len(lockNames))

	// 多重锁排序，防止出现死锁
	lockNames = m.sorter.Sort(lockNames)

	defer func() {
		for i := range lockList {
			lockList[i].Unlock()
		}
	}()
	// 上锁
	// 循环需要锁的Key
	for i := range m.lms {

		for j := range lockNames {
			lock, err := m.lms[i].GetLock(lockNames[j])
			if err != nil {
				return nil, err
			}
			lockList = append(lockList, lock)
		}
	}
	return f()
}
