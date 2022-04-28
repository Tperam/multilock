/*
 * @Author: Tperam
 * @Date: 2022-04-23 22:16:07
 * @LastEditTime: 2022-04-25 22:48:50
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\mutlilock_test.go
 */
package multilock_test

import (
	"sync"
	"testing"
	"time"

	"github.com/tperam/multilock"
	"github.com/tperam/multilock/algorithm"
	"github.com/tperam/multilock/lockcore"
	"github.com/tperam/multilock/locker"
)

func TestMultilock(t *testing.T) {
	m := multilock.NewMultilock(algorithm.NewBubbleSort(), lockcore.NewMapLockCore(locker.NewGenerateMutex(), 15))
	f := func() (interface{}, error) {
		// t.Log("")
		// t.Log("开始睡眠...")
		time.Sleep(1 * time.Nanosecond)
		// t.Log("结束睡眠...")
		return nil, nil
	}
	wg := &sync.WaitGroup{}
	wg.Add(200)
	for range [100]struct{}{} {

		go func() {
			m.Do(f, "user5", "user21", "user44", "user3", "user2", "user01", "user1")
			wg.Done()
		}()
		go func() {
			m.Do(f, "user1", "user2", "user4", "user3", "user2", "user0", "user5")
			wg.Done()
		}()
	}
	wg.Wait()

}

func TestMultilockCAS(t *testing.T) {
	m := multilock.NewMultilock(algorithm.NewBubbleSort(), lockcore.NewMapLockCore(locker.NewGenerateCASLock(), 15))
	f := func() (interface{}, error) {
		// t.Log("")
		// t.Log("开始睡眠...")
		time.Sleep(1 * time.Nanosecond)
		// t.Log("结束睡眠...")
		return nil, nil
	}
	wg := &sync.WaitGroup{}
	wg.Add(200)
	for range [100]struct{}{} {

		go func() {
			m.Do(f, "user5", "user21", "user44", "user3", "user2", "user01", "user1")
			wg.Done()
		}()
		go func() {
			m.Do(f, "user1", "user2", "user4", "user3", "user2", "user0", "user5")
			wg.Done()
		}()
	}
	wg.Wait()

}
