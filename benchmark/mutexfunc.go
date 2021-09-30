package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用锁实现生产者消费者机制
var lock = sync.Mutex{}
var buff = [10]int{}

func MutexProducer() {
	lock.Lock()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Println("生产者生产了", num)
		buff[i] = num
		time.Sleep(time.Microsecond * 300)
	}
	lock.Unlock()
}

func MutexConsumer() {
	lock.Lock()
	for i := 0; i < 10; i++ {
		fmt.Println("消费者消费到了", buff[i])
	}
	lock.Unlock()
}
