/*
 * @Author: Tperam
 * @Date: 2022-04-24 22:57:59
 * @LastEditTime: 2022-04-24 22:59:06
 * @LastEditors: Tperam
 * @Description:
 * @FilePath: \multilock\algorithm\bubblesort.go
 */
package algorithm

// N^2 ，通常此处不会出现过多的锁
// 此算法去重
func Bubblesort(lockName []string) []string {
	for i := 0; i < len(lockName); i++ {
		for j := i + 1; j < len(lockName); j++ {
			if lockName[i] == lockName[j] {
				lockName[j] = lockName[len(lockName)-1]
				lockName = lockName[:len(lockName)-1]
				j = i + 1
			}
			if lockName[i] > lockName[j] {
				// 交换
				lockName[j], lockName[i] = lockName[i], lockName[j]
			}

		}
	}
	return lockName
}
