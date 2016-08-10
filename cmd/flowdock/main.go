package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"crypto/rand"
	"encoding/base64"

	"github.com/gburanov/go-flowdock/flowdock"
	"github.com/gburanov/pr_helper/lib"
)

const organization = "wimdu"
const project = "wimdu"

func main() {
	client := flowdock.NewClientWithToken(nil, GetSettings().AuthToken)
	channel, _, err := client.Messages.Stream(GetSettings().StreamToken, "wimdu", "pr-helper-flow")
	if err != nil {
		log.Fatal(err)
	}
	for {
		message := <-channel
		if *message.Event == "message" {
			if message.UUID != nil && strings.HasPrefix(*message.UUID, "robot") {
				continue
			}
			response(message.Content().String(), client.Messages)
		}
	}
}

func randString() string {
	size := 32 // change the length of the generated random string here
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		log.Fatal(err)
	}
	rs := base64.URLEncoding.EncodeToString(rb)
	return rs
}

func sendMessage(message string, service *flowdock.MessagesService) {
	fmt.Println(message)
	fm := flowdock.MessagesCreateOptions{
		FlowID:  GetSettings().Flow,
		Event:   "message",
		Content: message,
		UUID:    "robot" + randString(),
	}
	_, _, err := service.Create(&fm)
	if err != nil {
		log.Fatal(err)
	}
}

func displayPR(pr pr_helper.PR, service *flowdock.MessagesService) {
	sendMessage(pr.Topic(), service)
	sendMessage(pr.Url(), service)
	authors := *pr.Authors()

	left, total := authors.GetLinesStat()
	percent := float32(left) / float32(total)

	str := fmt.Sprintf("%d lines out of %d lines unmaintained", left, total)
	sendMessage(str, service)
	if (total > 100 && percent > 0.7) || (percent > 0.9) {
		sendMessage("WARNING! DEEP LEGACY", service)
	}

	authorsStr := "Suggested authors for review "
	for author := range *pr_helper.FilterTop(5, &authors) {
		authorsStr += author.Name
		authorsStr += ", "
	}
	sendMessage(authorsStr, service)
}

func response(command string, service *flowdock.MessagesService) {
	num, err := strconv.Atoi(command)
	if err != nil {
		message := "Cannot find PR " + command
		sendMessage(message, service)
		return
	}
	sendMessage("Analysis in progress...", service)
	repo := pr_helper.NewRepository(organization, project, pr_helper.Token())
	pr, err := repo.GetPR(num)
	if err != nil {
		sendMessage(err.Error(), service)
		return
	}
	displayPR(pr, service)
}
