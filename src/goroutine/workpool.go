package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id int
	randomno  int //用于计算其每位数之和
}

type Result struct {
	job Job
	sumofdigits int //计算结果
}

var jobs = make(chan Job,10)
var results = make(chan Result,10)

func digits(number int ) int { // 用于计算，返回结果
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(1 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup)  {
	for job := range jobs{
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(numworker int){
	var wg sync.WaitGroup
	for i:=0;i<numworker;i++{
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results) //关闭results信道
}

//把作业分配给工作者
func allocate(numjobs int){
	for i:=0;i<numjobs;i++{
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main()  {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}