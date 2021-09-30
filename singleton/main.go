package main

import (
	"fmt"
	"sync"
)

var (
	//singleton *People
	singleton *People = &People{Name: "zhb", Age: 18}
	//once      sync.Once
	//mu sync.Mutex
)

type People struct {
	Name   string
	Age    int
	Weight string
	Hight  string
}

func main() {
	//once.Do实现单例例子
	/* var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Printf("当前是第%d个协程", i)
			})
		}(i)
	}
	wg.Wait() */

	//懒汉非线程安全实现单例
	/* var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			if singleton == nil {
				singleton = &People{Age: i}
				fmt.Printf("当前协程是第%d协程，本人%d岁\n", i, singleton.Age)
			} else {
				fmt.Printf("我是第%d协程,我得不到初始化\n", i)
			}
		}(i)
	}
	wg.Wait() */

	//懒汉加双重锁
	/* var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func(i int) {
			defer wg.Done()
			if singleton == nil {
				mu.Lock()
				defer mu.Unlock()
				if singleton == nil {
					singleton = &People{Age: i}
					fmt.Printf("当前协程是第%d协程，本人%d岁\n", i, singleton.Age)
				}
			} else {
				fmt.Printf("我是第%d协程,我得不到初始化\n", i)

			}
		}(i)
	}
	wg.Wait() */

	//饿汉,类似已初始化的全局变量
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println(singleton, n)
		}(i)
	}
	wg.Wait()
}
