package main

import (
	"log"

	"github.com/gburanov/pr_helper"
	"github.com/gburanov/pr_helper/sqs_lib"
)

func processURL(url string, cb pr_helper.Callback, m *pr_helper.Mutex) {
	showPr(url, cb, m)
	cb("END OF MESSAGE")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	m := pr_helper.NewMutex()

	inputQueue, err := sqs_lib.InputQueue()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	for {
		message := inputQueue.ReadMessage()
		go processURL(message.Body(), message.Response(), m)
		message.Delete()
	}
}
