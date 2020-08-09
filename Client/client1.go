package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	_, err = conn.Write([]byte("hahahahah"))

	for {

	}

}
