package main

import (
	"fmt"
	"log"

	"github.com/gburanov/pr_helper/lib"
	"github.com/joho/godotenv"
)

func processURL(url string, cb pr_helper.Callback, m *pr_helper.Mutex) {
	showPr(url, cb, m)
	cb("END OF MESSAGE")
}

func messageCallback(uuid string) pr_helper.Callback {
	return func(str string, args ...interface{}) {
		message := fmt.Sprintf(str, args...)
		fmt.Println(message)
		SendMessage(message, uuid)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	m := pr_helper.NewMutex()

	for {
		message := ReadMessage()
		var uuid string
		for _, attr := range message.MessageAttribute {
			if attr.Name == "uuid" {
				uuid = attr.Value.StringValue
				break
			}
		}
		if len(uuid) == 0 {
			fmt.Println("Invalid message processed")
		} else {
			callback := messageCallback(uuid)
			go processURL(message.Body, callback, m)
		}
		DeleteMessage(message)
	}
}
