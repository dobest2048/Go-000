# 学习笔记
## 作业
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

## 分析
- 明确查询一条数据时，查询数据由结构体接收，如果查询不到数据时用指针判断（*user != nil）方式不友好，所以此时是创建错误，不需要Wrap
- 明确查询多条数据时，接收的数据结构为[]*user slice可以直接用len判断，所以无数据时不用报错（对错误进行了处理）
## 代码示例
```cmd
git clone github.com/jikebang/Go-000
cd Week02
go run main.go
```
