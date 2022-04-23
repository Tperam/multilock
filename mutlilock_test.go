/*
 * @Author: Tperam
 * @Date: 2022-04-23 22:16:07
 * @LastEditTime: 2022-04-23 22:46:02
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\mutlilock_test.go
 */
package multilock_test

import (
	"fmt"
	"multilock"
	"sync"
	"testing"
	"time"
)

func TestMultilock(t *testing.T) {
	m := multilock.NewMultilock(10)
	f := func() (interface{}, error) {
		fmt.Println("")
		fmt.Println("开始睡眠...")
		time.Sleep(1 * time.Second)
		fmt.Println("结束睡眠...")
		return nil, nil
	}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		m.Do(f, "user5", "user21", "user44", "user3", "user2", "user01", "user1")
		wg.Done()
	}()
	go func() {
		m.Do(f, "user1", "user2", "user4", "user3", "user2", "user0", "user5")
		wg.Done()
	}()
	wg.Wait()

}
