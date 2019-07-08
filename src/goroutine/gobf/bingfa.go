package gobf

import (
	"fmt"
	"sync"
	"time"
)

func Hello()  {
	fmt.Println("Hello world goroutine")
}

func Numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d\n ", i)
	}
}
func Alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c\n ", i)
	}
}

func Hello2(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}


func AlcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	fmt.Println(sum)
	squareop <- sum
}

func CalcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	fmt.Println(sum)
	cubeop <- sum
}

func Producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)  //关闭信道
}


func Digits(number int, chn chan int){
	for number != 0{
		digit := number % 10
		chn <- digit
		number /= 10
	}
	close(chn)
}

func CalcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go Digits(number, dch)
	for digit := range dch{
		sum += digit * digit
	}
	squareop <- sum
}

func CalcCubeses(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go Digits(number, dch)
	for digit := range dch{
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func Write(ch chan int)  {
	for i :=0;i<5;i++{
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func Prosess(i int, wg *sync.WaitGroup ){
	fmt.Println("started Goroutine ", i)
	time.Sleep(1*time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()  //减少计数器
}

func Server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func Server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}