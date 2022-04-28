/*
 * @Author: Tperam
 * @Date: 2022-04-24 23:43:36
 * @LastEditTime: 2022-04-28 17:40:26
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\locker\cas_benchmark_test.go
 */
package locker_test

import (
	"testing"

	"github.com/tperam/multilock/locker"
)

func BenchmarkCASLock(b *testing.B) {
	gcas := locker.NewGenerateCASLock()
	lock, err := gcas.New("")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock()
	}
}
