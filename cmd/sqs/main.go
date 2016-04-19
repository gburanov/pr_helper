package main

import (
  "fmt"
  "pr_helper"
)

func ProcessUrl(url string, cb pr_helper.Callback) {
  fmt.Println("Processing url", url)
  showPr(url, cb)
}

func messageCallback(uuid string) pr_helper.Callback {
    return func(str string, args ...interface{}) {
        message := fmt.Sprintf(str, args...)
        SendMessage(message, uuid)
    }
}

func main() {
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
      ProcessUrl(message.Body, callback)
      callback("END OF MESSAGE")
    }
    DeleteMessage(message)
  }
}
