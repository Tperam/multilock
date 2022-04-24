/*
 * @Author: Tperam
 * @Date: 2022-04-24 23:34:22
 * @LastEditTime: 2022-04-24 23:44:38
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\locker\cas.go
 */
package locker

import "sync/atomic"

type CASLock struct {
	i uint32
}

func (cas *CASLock) Lock() error {
	for !atomic.CompareAndSwapUint32(&cas.i, 0, 1) {
	}
	return nil
}
func (cas *CASLock) Unlock() error {
	if !atomic.CompareAndSwapUint32(&cas.i, 1, 0) {
		return ErrRepeatUnlocked
	}
	return nil
}

type GenerateCASLock struct{}

func NewGenerateCASLock() *GenerateCASLock {
	return &GenerateCASLock{}
}

func (gm *GenerateCASLock) New(lockname string) (Locker, error) {
	return &CASLock{}, nil
}
