package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (result string, err error) {
	resp, errl := http.Get(url)
	if errl != nil {
		err = errl
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err=", err)
			break
		}
		result += string(buf[:n])
	}
	return
}

func WritePage(j int, page chan<- int) {
	url := "https://tieba.baidu.com/f?kw=绝地求生&ie=utf-8&pn=" + strconv.Itoa((j-1)*50)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err=", err)
		return
	}
	fileName := strconv.Itoa(j) + ".html"
	f, errl := os.Create(fileName)
	if errl != nil {
		fmt.Println("os.Create errl=", errl)
		return
	}else{
		fmt.Println("创建文件成功！")
	}
	f.WriteString(result)
	f.Close()
	page <- j
}

func DoWork(start, end int) {
	fmt.Printf("正在爬取%d到%d的页面\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go WritePage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页（>=1）：")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>=起始页）：")
	fmt.Scan(&end)

	DoWork(start, end)

}
