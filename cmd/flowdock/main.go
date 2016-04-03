package main

import (
  "fmt"
  "log"

  "github.com/gburanov/go-flowdock/flowdock"
)

func main() {
  client := flowdock.NewClientWithToken(nil, GetSettings().AuthToken)
  flows, _, err := client.Organizations.All()
	if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Hi")
  fmt.Println(flows)
}
