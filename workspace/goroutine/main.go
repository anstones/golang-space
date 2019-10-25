package main

import (
	"fmt"
	"goroutine/gobf"
	"sync"
)

func main()  {
	//go gobf.Numbers()
	//go gobf.Alphabets()
	//time.Sleep(3 * time.Second)
	//fmt.Println("main function")

	//done := make(chan bool) // 定义信道
	//fmt.Println("Main going to call hello go goroutine")
	//go gobf.Hello2(done)
	//<-done //此时处于阻塞状态，需要往信道内抛数据 才会结束阻塞
	//fmt.Println("Main received data\n")

	//number := 589
	//a := make(chan int)
	//b := make(chan  int)
	//go gobf.AlcSquares(number, a)
	//go gobf.CalcCubes(number, b)
	//x, y := <-a,<-b
	//fmt.Println("Final output", x+y)


	//ch := make(chan int)
	//go gobf.Producer(ch)
	//for {
	//	v,ok := <-ch
	//	if ok == false{
	//		break
	//	}
	//	fmt.Println(v)
	//}
	//for v := range ch{  //for range 遍历信道，直至信道关闭
	//	fmt.Println(v)
	//}

	//number1 := 789
	//c := make(chan int)
	//d := make(chan  int)
	//go gobf.CalcSquares(number1, c)
	//go gobf.CalcCubeses(number1, d)
	//m, n := <-c,<-d
	//fmt.Println("Final output",m+n)

	//ch := make(chan string,2) // 定义缓冲信道
	//ch <- "james"
	//ch <- "AD"
	////ch <- "IK" // 会造成死锁
	//fmt.Println(<- ch)
	//fmt.Println(<- ch)

	//chn := make(chan int,2)
	//go gobf.Write(chn)
	//time.Sleep(2*time.Second)
	//for v := range chn{
	//	fmt.Println("read value", v,"from ch")  //定义了缓冲为2的信道，当往信道填充2个值后，信道阻塞，直至2个值被消费
	//	time.Sleep(2*time.Second)
	//}


	chna := make(chan string, 3)
	chna <- "naveen"
	chna <- "paul"
	fmt.Println("capacity is", cap(chna))  //容量，即定义信道时信道的大小
	fmt.Println("length is", len(chna))   //长度，信道内存在的元素个数
	fmt.Println("read value", <-chna)
	fmt.Println("new length is", len(chna))

	var wg sync.WaitGroup
	for i := 0;i<3;i++{
		wg.Add(1)  // 增加计数器
		go gobf.Prosess(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")


	output1 := make(chan string)
	output2 := make(chan string)
	go gobf.Server1(output1)
	go gobf.Server2(output2)

	select {  //只会选择一个case 执行
	case s1:= <-output1:
		fmt.Println(s1)
	case s2:= <- output2:
		fmt.Println(s2)
	}

}
