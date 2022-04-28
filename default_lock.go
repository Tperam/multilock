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
	"github.com/tperam/multilock/algorithm"
	"github.com/tperam/multilock/lockcore"
	"github.com/tperam/multilock/locker"
)

func NewDefaultCore() *Multilock {
	return NewMultilock(algorithm.NewBubbleSort(), lockcore.NewMapLockCore(locker.NewGenerateMutex(), 10))
}
