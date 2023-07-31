package main

import (
	"fmt"
	"time"
)

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				arr[i]++
			}
		}(i)
	}
	time.Sleep(time.Minute) // 休眠1分钟
	fmt.Print("arr: ", arr)
}
