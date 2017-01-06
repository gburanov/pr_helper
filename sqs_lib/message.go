package sqs_lib

import (
	"fmt"
	"log"

	"github.com/gburanov/pr_helper"
	"github.com/goamz/goamz/sqs"
)

type Message struct {
	uuid    string
	message *sqs.Message
	queue   *Queue
}

func (m *Message) Body() string {
	return m.message.Body
}

func (m *Message) Response() pr_helper.Callback {
	return func(str string, args ...interface{}) {
		message := fmt.Sprintf(str, args...)
		fmt.Println(message)
		outputQueue, err := m.queue.opposite()
		if err != nil {
			log.Fatal(err)
		}
		outputQueue.SendMessage(message, m.uuid)
	}
}

func (m *Message) Delete() error {
	_, err := m.queue.Queue.DeleteMessage(m.message)
	return err
}

func (m *Message) init() {
	var uuid string
	for _, attr := range m.message.MessageAttribute {
		if attr.Name == "uuid" {
			uuid = attr.Value.StringValue
			break
		}
	}
	if len(uuid) == 0 {
		fmt.Println("Invalid message processed")
	}
	m.uuid = uuid
}

func (queue *Queue) ReadMessage() *Message {
	fmt.Println("Waiting for messages...")
	for {
		resp, err := queue.Queue.ReceiveMessage(1)
		if err != nil {
			log.Print(err)
			continue
		}
		if len(resp.Messages) != 0 {
			m := &Message{message: &resp.Messages[0], queue: queue}
			m.init()
			return m
		}
	}
}
