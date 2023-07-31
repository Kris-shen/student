package main

func main() {
	////timer基本使用
	//timer1 := time.NewTimer(2 * time.Second)
	//t1 := time.Now()
	//fmt.Println("t1:", t1)
	//t2 := <-timer1.C
	//fmt.Println("t2:", t2)

	//// 2.验证timer只能响应1次
	//timer2 := time.NewTimer(time.Second)
	//for true {
	//	<-timer2.C
	//	fmt.Println("时间到")
	//}
	// 3.timer实现延时的功能
	//(1)
	//time.Sleep(time.Second)
	//(2)
	//timer3 := time.NewTimer(2 * time.Second)
	//<-timer3.C
	//fmt.Println("2秒到")
	//(3)
	//	<-time.After(2 * time.Second)
	//	fmt.Println("2秒到")
	// 4.停止定时器
	//timer4 := time.NewTimer(3 * time.Second)
	//go func() {
	//	<-timer4.C
	//	fmt.Println("定时器执行了")
	//}()
	//b := timer4.Stop()
	//if b {
	//	fmt.Println("timer4已经关闭")
	//}

}
