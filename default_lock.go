/*
 * @Author: Tperam
 * @Date: 2022-04-25 22:34:41
 * @LastEditTime: 2022-04-25 22:34:41
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\default_lock.go
 */
package multilock

import (
	"multilock/algorithm"
	"multilock/lockcore"
	"multilock/locker"
)

func NewDefaultCore() *Multilock {
	return NewMultilock(algorithm.NewBubbleSort(), lockcore.NewMapLockCore(locker.NewGenerateMutex(), 10))
}
