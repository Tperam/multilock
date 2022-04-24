/*
 * @Author: Tperam
 * @Date: 2022-04-24 23:07:55
 * @LastEditTime: 2022-04-24 23:40:38
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\locker\mutex.go
 */
package locker

import "sync"

type Mutex struct {
	sync.Mutex
}

func (l *Mutex) Lock() error {
	l.Mutex.Lock()
	return nil
}
func (l *Mutex) Unlock() error {
	l.Mutex.Unlock()
	return nil
}

type GenerateMutex struct{}

func NewGenerateMutex() *GenerateMutex {
	return &GenerateMutex{}
}

func (gm *GenerateMutex) New(lockname string) (Locker, error) {
	return &Mutex{}, nil
}
