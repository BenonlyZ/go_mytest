package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("can't receive")
			break
		}

		fmt.Println("Received back from client :" + reply)

		msg := "Reveived:" + reply

		fmt.Println("Sending to client:" + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("can't send")
			break
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":8099", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
