package main

import (
	"fmt"
	"log"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
)

func getQueue() (*sqs.Queue, error) {
	if len(auth.AccessKey) == 0 || len(auth.SecretKey) == 0 {
		return nil, fmt.Errorf("AWS credentials are not provided")
	}
	conn := sqs.New(auth, aws.EUWest)
	return conn.CreateQueue("pr_helper_input")
}

func ReadMessage() sqs.Message {
	fmt.Println("Waiting for messages...")
	q, err := getQueue()
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
	conn := sqs.New(auth, aws.EUWest)
	q, err := conn.CreateQueue("pr_helper_input")
	if err != nil {
		log.Fatal(err)
	}
	_, err = q.DeleteMessage(&m)
	if err != nil {
		log.Fatal(err)
	}
}

func SendMessage(message string, uuid string) {
	conn := sqs.New(auth, aws.EUWest)
	q, err := conn.CreateQueue("pr_helper_output")
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
