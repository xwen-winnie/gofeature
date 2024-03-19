package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
	//test1()
	test2()
}

// context 包来处理超时和取消的情况
// 带有超时的上下文, 结合select 语句来监视超时和上下文的完成情况
func test1() {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()
	select {
	// time.After 创建一个计时器，在指定的时间间隔之后向返回的通道发送一个值
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func test2() {
	tooSlow := fmt.Errorf("too slow!")
	ctx, cancel := context.WithTimeoutCause(context.Background(), 1*time.Second, tooSlow)
	time.Sleep(2 * time.Second) //阻塞程序的执行
	cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		//context.Cause() 函数获取上下文被取消的原因
		fmt.Println(context.Cause(ctx))
	}
}

// 回调的 demo 代码片段
func WithFirstCancel(ctx1, ctx2 context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx1)
	// 注册匿名回调函数
	stopf := context.AfterFunc(ctx2, func() { cancel() })
	// 多 Context 合并取消
	return ctx, func() {
		cancel()
		stopf()
	}
}
