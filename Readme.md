# multilock



本项目为多重锁，用于防止处理事务时可能发生的死锁问题。



外部调用方式:

```go
m := multilock.NewMultilock(10)

err := m.ExecFunc(f,"lockname 1","lockname 2","lockname 3")
if err != nil {
    ...
}
```

