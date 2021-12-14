package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job 表示任务
type Job struct {
	id       int
	randomno int
}

// Result 表示结果
type Result struct {
	job         Job
	sumofdigits int
	worker      int
}

// 分别创建用于接收作业和写入结果的缓冲信道
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

//func main() {
//	workPoolDemo()
//}

func workPoolDemo() {
	startTime := time.Now()

	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()

	// 计算消耗时间
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// 创建工作协程
func worker(wg *sync.WaitGroup, workerId int) {
	//fmt.Println("开始工作")
	for job := range jobs {
		// 计算并返回结果
		output := Result{job, digits(job.randomno), workerId}
		results <- output
	}
	wg.Done()
	//fmt.Println("工作结束")
}

// 创建工作池
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i)
		fmt.Println("创建worker协程")
	}
	wg.Wait()
	close(results)
}

// 分配任务
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

// 读取任务结果
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, worker id %d,input random no %d , sum of digits %d\n", result.job.id, result.worker, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
