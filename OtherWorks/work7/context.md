# Context源码阅读报告
## 什么是Context
Context专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消。
## 源码
```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```
## 方法
### 初始化
初始化`context`的两种方法，没有任何区别，只是写法不同。
```go
// 创建两个空的context
var ctx1 = context.TODO();
var ctx2 = context.Background();
```
### Deadline
可用于设置上下文的过期时间， 当多次设置过期时间时，只会取最早到期的过期时间。

***Deadline() (deadline time.Time, ok bool)***
- `deadline` 上下文的过期时间
- `ok` 是否存在过期时间，是为true

***WithTimeout(parent Context, timeout time.Duration) (child Context,cancel CancelFunc)***
- `parent` 传入context
- `timeout` 设置多长时间后过期
- `child` 传出context
- `cancel` 关闭函数，当context达到过期时间后自动执行，但也可以手动执行`cancel()`，此时无论context是否到期，都会关闭

另外，如果不需要定时关闭，也可以通过`WithCancel()`函数只获得一个关闭函数。

***WithCancel(parent Context) (ctx Context, cancel CancelFunc)***
- `parent` 传入context
- `ctx` 传出context
- `cancel` 关闭函数

```go
fmt.Println("start:", time.Now())
// 父context的过期时间为5秒后
father, cancel1 := context.WithTimeout(context.TODO(), time.Second*5)
defer cancel1()
// 子context的过期时间为10秒后
son, cancel2 := context.WithTimeout(father, time.Second*10)
defer cancel2()
// 打印子context过期时间，为5秒后
deadline, ok := son.Deadline()
fmt.Printf("son deadline:%v ok:%v\n", deadline, ok)
```
运行结果：
```
start: 2023-07-06 13:25:38.4958247 +0800 CST m=+0.002554401
son deadline:2023-07-06 13:25:43.5119692 +0800 CST m=+5.018698901 ok:true
```
### Done
返回一个只读channel，当context被关闭时，会向`Done()`中写入空struct
```go
fmt.Println("start:", time.Now())
ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
defer cancel()
select {
case <-ctx.Done():
    fmt.Println("ctx done, time:", time.Now())
}
```
运行结果：
```
start: 2023-07-06 14:17:34.9937486 +0800 CST m=+0.003599201
ctx done, time: 2023-07-06 14:17:40.0176379 +0800 CST m=+5.027488501
```
### Err
返回context关闭的原因

```go
// 通过deadline关闭的context
ctx_timeout, _ := context.WithTimeout(context.TODO(), time.Second*1)
select {
case <-ctx_timeout.Done():
    fmt.Println("ctx timeout, err:", ctx_timeout.Err())
}
// 通过cancel关闭的context
ctx_cancel, cancel := context.WithCancel(context.TODO())
go func() {
    time.Sleep(time.Second * 1)
    cancel()
}()
select {
case <-ctx_cancel.Done():
    fmt.Println("ctx cancel, err:", ctx_cancel.Err())
}
```
运行结果：
```
ctx timeout, err: context deadline exceeded
ctx cancel, err: context canceled
```
### Value
value中保存的是一个`map[any]any`，用于在上下文中存取数据。

相比显式传递所有参数， 通过将多个变量封装在context中， 可以传递任意数量的变量，且在上下文中传递。

***Value(key any) (value any)***
- `key` 传入任意类型key
- `value` 传出对应value，如果key不存在则panic

***WithValue(parent Context, key any, value any) (child Context)***
- `parent` 传入context
- `key` 传入任意类型key
- `value` 传入任意类型value
- `child` 传出context

```go
func f1(ctx context.Context) context.Context {
    return context.WithValue(ctx, "f1", 1)
}
func f2(ctx context.Context) context.Context {
    fmt.Println("father:",ctx.Value("f1"))
    return context.WithValue(ctx, 2, "f2")
}
func main() {
    grandpa := context.TODO()
    father := f1(grandpa)
    son := f2(father)
    fmt.Println("son:",son.Value("f1"))
    fmt.Println("son:",son.Value(2))
} 
```
输出结果：
```
father:1
son:1
son:f2
```
## 使用Context的注意事项
- 推荐以参数的方式显示传递Context 
- 以Context作为参数的函数方法，应该把Context作为第一个参数。 
- 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
- Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数 
- Context是线程安全的，可以放心的在多个goroutine中传递