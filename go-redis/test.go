package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

func main() {

	var client goredis.Client
	client.Set("c", []byte("haha"))
	val, _ := client.Get("c")
	fmt.Println(string(val))
	//	client.Del("a")

	vals := []string{"a1", "b1", "c1", "d1", "e1"}
	for _, v := range vals {
		client.Rpush("d", []byte(v))
	}
	dbvals, _ := client.Lrange("d", 0, 4)
	for i, v := range dbvals {
		println(i, ":", string(v))
	}
	//client.Del("b")
}
