package main

import (
	"fmt"
)

// 无缓冲通道 会阻塞 会死锁 会报错 fatal error: all goroutines are asleep - deadlock!
//func main() {
//	//无缓冲的通道只有在有接收的情况下才可以发送数据
//	ch := make(chan int)
//	ch <- 10
//	fmt.Println("发送成功")
//}

func recv(c chan int) {
	ret := <-c
	//因为是同步的，所以这里会阻塞，直到ch有数据写入 才会继续往下执行
	//也是因为这个原因，此次答应可能无法输出就结束了
	fmt.Println("接收成功", ret)
}

func main() {

	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10    // 把10发送到通道中
	fmt.Println("发送成功")
}
