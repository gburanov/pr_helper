package main

import "fmt"
import "log"

import "github.com/google/go-github/github"

func main() {
  auth_token := token()
  fmt.Println(auth_token)
  client := github.NewClient(auth_token)
  orgs, _, err := client.Organizations.List("willnorris", nil)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(orgs)
}
