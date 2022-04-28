/*
 * @Author: Tperam
 * @Date: 2022-04-24 23:38:55
 * @LastEditTime: 2022-04-24 23:46:50
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\locker\mutex_test.go
 */
package locker_test

import (
	"testing"

	"github.com/tperam/multilock/locker"
)

func TestMutexError(t *testing.T) {
	gm := locker.NewGenerateMutex()
	s, err := gm.New("")
	if err != nil {
		panic(err)
	}
	s.Unlock()
}
