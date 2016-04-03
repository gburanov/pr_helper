package main

import (
  "fmt"
  "log"
  "strconv"

  "github.com/gburanov/go-flowdock/flowdock"

  "pr_helper"
)

const organization = "wimdu"
const project = "wimdu"

func main() {
  client := flowdock.NewClientWithToken(nil, GetSettings().AuthToken)
  channel, _, err := client.Messages.Stream(GetSettings().StreamToken,"wimdu", "pr-helper-flow")
	if err != nil {
    log.Fatal(err)
  }
  for {
    message := <- channel
    if *message.Event == "message" {
      response(message.Content().String(), client.Messages)
    }
  }
}

func sendMessage(message string, service *flowdock.MessagesService) {
  fmt.Println(message)
  fm := flowdock.MessagesCreateOptions{
    FlowID: GetSettings().Flow,
    Event: "message",
    Content: message,
  }
    _, _, err := service.Create(&fm)
    if err != nil {
      log.Fatal(err)
    }
}

func displayPR(pr *pr_helper.PR, service *flowdock.MessagesService) {
  sendMessage(pr.ShowInfo(), service)
  for author, _ := range *pr.Authors() {
    sendMessage(author.AsStr(), service)
  }
}

func response(command string, service *flowdock.MessagesService) {
  num, err := strconv.Atoi(command)
  if err != nil {
    return
  }

  repo := pr_helper.NewRepository(organization, project, pr_helper.Token())
  displayPR(repo.GetPR(num), service)
}
