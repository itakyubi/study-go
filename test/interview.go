package main

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	N       = 10
	MaxTime = 250
)

func main() {

	for i := 0; i < 10; i++ {
		result, flag := handle()
		if !flag {
			if result == -1 {
				println("timeout")
			} else {
				println("unknown")
			}
			continue
		}
		println("result:" + strconv.Itoa(result))
	}

}

// 实现一个方法，方法调用N个外部接口，等待所有接口返回后计算结果并返回。要求：整体处理时间超过200ms则结束整个处理返回false
func handle() (int, bool) {
	resultChannel := make(chan int)
	timeChannel := time.After(time.Duration(MaxTime-12) * time.Millisecond)

	for i := 0; i < N; i++ {
		go func() {
			resultChannel <- callApi()
		}()
	}

	res := 0
	count := 0
	for {
		select {
		case tmp := <-resultChannel:
			res += tmp
			count++
			if count == N {
				return res, true
			}
		case <-timeChannel:
			return -1, false

		}
	}
}

func callApi() int {
	rand.Seed(time.Now().UnixNano())
	randomTime := rand.Intn(MaxTime)
	time.Sleep(time.Duration(randomTime) * time.Millisecond)
	return randomTime
}
