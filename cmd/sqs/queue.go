package main

import (
  "log"
  "fmt"

  "github.com/goamz/goamz/aws"
  "github.com/goamz/goamz/sqs"
)

func ReadMessage() sqs.Message {
  conn := sqs.New(auth, aws.EUWest)
  q, err := conn.CreateQueue("pr_helper_input")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Wainting for messages...")
  for {
    resp, err := q.ReceiveMessage(1)
    if err != nil {
      log.Fatal(err)
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
