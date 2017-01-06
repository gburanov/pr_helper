package sqs_lib

import (
	"fmt"
	"log"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
	"github.com/joho/godotenv"
)

type Queue struct {
	Name  string
	Queue *sqs.Queue
}

const inputQueue string = "pr_helper_input"
const outputQueue string = "pr_helper_output"

func getQueue(name string) (*Queue, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if len(auth.AccessKey) == 0 || len(auth.SecretKey) == 0 {
		return nil, fmt.Errorf("AWS credentials are not provided")
	}
	conn := sqs.New(auth, aws.EUWest)
	queue, err := conn.CreateQueue(name)
	if err != nil {
		return nil, err
	}
	return &Queue{Name: name, Queue: queue}, nil
}

func InputQueue() (*Queue, error) {
	return getQueue(inputQueue)
}

func OutputQueue() (*Queue, error) {
	return getQueue(outputQueue)
}

func (queue *Queue) opposite() (*Queue, error) {
	if queue.Name == inputQueue {
		return getQueue(outputQueue)
	}
	return getQueue(inputQueue)
}

func (queue *Queue) SendMessage(message string, uuid string) error {
	m := map[string]string{}
	m["uuid"] = uuid
	_, err := queue.Queue.SendMessageWithAttributes(message, m)
	return err
}
