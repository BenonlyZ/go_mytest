package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ChanProducer(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		ch <- num
		fmt.Println("生产者生产了", num)
		time.Sleep(time.Millisecond * 300)
	}
}
func ChanProducer2(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		ch <- num
		fmt.Println("生产者生产了", num)
		time.Sleep(time.Millisecond * 300)
	}
}

func ChanConsumer(ch chan int) {
	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("消费者消费到了", num)
	}
}
