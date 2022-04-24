# multilock



本项目为多重分布式锁，用于防止处理事务时可能发生的死锁问题。

#### 使用方式

外部调用方式:

```go
m := multilock.NewMultilock(10)

err := m.ExecFunc(f,"lockname 1","lockname 2","lockname 3")
if err != nil {
    ...
}
```





#### 实现思路 & 评价

当前逻辑实现过于紧密，不宜与修改拆分，后续将按一下路径拆分实现。

1. 算法部分
   - 防止死锁的核心就是让锁按顺序进行，此处我们将剥离算法，单独实现，让结果可复用。
2. 本地锁部分
   - 本地锁部分主要解决的问题是tcp之间的相互链接。
3. 外部锁部分
   - 外部锁，主要用于防止不同进程对同一片数据竞争导致的问题。