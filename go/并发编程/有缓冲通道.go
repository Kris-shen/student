package main

import "fmt"

// 只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。
// 就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。
//func main() {
//	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
//	ch <- 10
//	fmt.Println("发送成功")
//}

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for true {
		if data, ok := <-c; ok {
			println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}
