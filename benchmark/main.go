package main

/* func printer(str string) {
	lock.Lock()
	for _, val := range str {
		fmt.Printf("%c", val)
		//time.Sleep(time.Millisecond * 100)
	}
	lock.Unlock()
} */

func main() {
	myChan := make(chan int, 10)
	go ChanProducer(myChan)
	go ChanProducer2(myChan)
	go ChanConsumer(myChan)
	for {

	}

}
