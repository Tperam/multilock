/*
 * @Author: Tperam
 * @Date: 2022-04-24 23:46:53
 * @LastEditTime: 2022-04-24 23:47:07
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\locker\mutex_benchmark_test.go
 */
/*
 * @Author: Tperam
 * @Date: 2022-04-24 23:43:36
 * @LastEditTime: 2022-04-24 23:45:49
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\locker\cas_benchmark_test.go
 */
package locker_test

import (
	"multilock/locker"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	gcas := locker.NewGenerateMutex()
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
