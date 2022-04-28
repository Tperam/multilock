/*
 * @Author: Tperam
 * @Date: 2022-04-26 22:21:16
 * @LastEditTime: 2022-04-26 22:31:15
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\mutlilock_benchmark_test.go
 */
/*
 * @Author: Tperam
 * @Date: 2022-04-23 22:16:07
 * @LastEditTime: 2022-04-26 22:21:04
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\mutlilock_test.go
 */
package multilock_test

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/tperam/multilock"
	"github.com/tperam/multilock/algorithm"
	"github.com/tperam/multilock/lockcore"
	"github.com/tperam/multilock/locker"
)

func BenchmarkMultilock(b *testing.B) {
	// one alloc
	m := multilock.NewMultilock(algorithm.NewBubbleSort(), lockcore.NewMapLockCore(locker.NewGenerateMutex(), 15))
	f := func() (interface{}, error) {
		// t.Log("")
		// t.Log("开始睡眠...")
		time.Sleep(5 * time.Millisecond)
		// t.Log("结束睡眠...")
		return nil, nil
	}
	for i := 0; i < b.N; i++ {

		wg := &sync.WaitGroup{} // one alloc
		wg.Add(1)
		for range [1]struct{}{} {

			go func() {

				locknames := make([]string, 1)
				for i := range locknames {
					// one alloc
					random := rand.Int()
					locknames[i] = strconv.Itoa(random)
				}

				m.Do(f, locknames...)
				wg.Done()
			}()
		}
		wg.Wait()
	}

}

func BenchmarkMultilockCas(b *testing.B) {
	m := multilock.NewMultilock(algorithm.NewBubbleSort(), lockcore.NewMapLockCore(locker.NewGenerateCASLock(), 15))
	f := func() (interface{}, error) {
		// t.Log("")
		// t.Log("开始睡眠...")
		time.Sleep(5 * time.Millisecond)
		// t.Log("结束睡眠...")
		return nil, nil
	}
	for i := 0; i < b.N; i++ {

		wg := &sync.WaitGroup{}
		wg.Add(1)
		for range [1]struct{}{} {

			go func() {
				locknames := make([]string, 5)
				for i := range locknames {
					random := rand.Int()
					locknames[i] = strconv.Itoa(random)
				}

				m.Do(f, locknames...)
				wg.Done()
			}()
		}
		wg.Wait()
	}

}

func BenchmarkLock(b *testing.B) {
	// one alloc
	m := multilock.NewMultilock(algorithm.NewBubbleSort(), lockcore.NewMapLockCore(locker.NewGenerateMutex(), 15))
	f := func() (interface{}, error) {
		// t.Log("")
		// t.Log("开始睡眠...")
		time.Sleep(5 * time.Millisecond)
		// t.Log("结束睡眠...")
		return nil, nil
	}
	for i := 0; i < b.N; i++ {

		// one alloc
		wg := &sync.WaitGroup{}
		wg.Add(1)
		for range [1]struct{}{} {

			go func() {
				m.Do(f, "1")
				wg.Done()
			}()
		}
		wg.Wait()
	}

}
