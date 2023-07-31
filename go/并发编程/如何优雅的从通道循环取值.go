package main

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	//开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- 1
		}
		close(ch1)
	}()
	//开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	for i := range ch2 {
		println(i)
	}

}
