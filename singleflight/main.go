package main

import (
	"errors"
	"log"
	"sync"

	"golang.org/x/sync/singleflight"
)

//实验证明singleflight只能对同一key的并发请求做处理;多个不同key则需要分别处理
var g singleflight.Group
var errorNotExist = errors.New("not exist")

func main() {
	var wg sync.WaitGroup
	wg.Add(20)

	//模拟10个并发
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data1, err := getData("zhb")
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data1, "我是data1")
		}()
		go func() {
			defer wg.Done()
			data2, err := getData("Ben")
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data2, "我是data2")
		}()

	}
	wg.Wait()

}

//获取数据
func getData(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errorNotExist {
		//模拟从db中获取数据
		v, err, _ := g.Do(key, func() (interface{}, error) {
			return getDataFromDB(key)
		})

		if err != nil {
			log.Println(err)
			return "", err
		}

		data = v.(string)
		//TOOD: set cache
	} else if err != nil {
		return "", err
	}
	return data, nil
}

//模拟从cache中获取值，cache中无该值
func getDataFromCache(key string) (string, error) {
	return "", errorNotExist
}

//模拟从数据库中获取值
func getDataFromDB(key string) (string, error) {
	log.Printf("get %s from database", key)
	return "data", nil
}
