package main

import (
	"fmt"
	"log"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
)

const inputQueue string = "pr_helper_input"
const outputQueue string = "pr_helper_output"

func getQueue(name string) (*sqs.Queue, error) {
	if len(auth.AccessKey) == 0 || len(auth.SecretKey) == 0 {
		return nil, fmt.Errorf("AWS credentials are not provided")
	}
	conn := sqs.New(auth, aws.EUWest)
	return conn.CreateQueue(name)
}

func ReadMessage() sqs.Message {
	fmt.Println("Waiting for messages...")
	q, err := getQueue(inputQueue)
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, err := q.ReceiveMessage(1)
		if err != nil {
			log.Print(err)
			continue
		}
		if len(resp.Messages) != 0 {
			return resp.Messages[0]
		}
	}
}

func DeleteMessage(m sqs.Message) {
	q, err := getQueue(inputQueue)
	if err != nil {
		log.Fatal(err)
	}
	_, err = q.DeleteMessage(&m)
	if err != nil {
		log.Fatal(err)
	}
}

func SendMessage(message string, uuid string) {
	q, err := getQueue(outputQueue)
	if err != nil {
		log.Fatal(err)
	}
	m := map[string]string{}
	m["uuid"] = uuid
	_, err = q.SendMessageWithAttributes(message, m)
	if err != nil {
		log.Fatal(err)
	}
}
