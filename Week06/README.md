# 学习笔记

# 作业
参考 Hystrix 实现一个滑动窗口计数器。

答题：work文件夹

单元测试如下：
```cassandraql
go test -v
=== RUN   TestSlidingWindow
    TestSlidingWindow: sliding_window_test.go:22: 1秒的bucket长度 12
    TestSlidingWindow: sliding_window_test.go:24: 21 65
    TestSlidingWindow: sliding_window_test.go:24: 86 171
    TestSlidingWindow: sliding_window_test.go:24: 76 178
    TestSlidingWindow: sliding_window_test.go:24: 95 186
    TestSlidingWindow: sliding_window_test.go:26: ==============
    TestSlidingWindow: sliding_window_test.go:28: 21 65
    TestSlidingWindow: sliding_window_test.go:28: 138 289
    TestSlidingWindow: sliding_window_test.go:28: 147 308
--- PASS: TestSlidingWindow (11.34s)
PASS
ok      github.com/jikebang/Go-000/week06/work  11.829s
```