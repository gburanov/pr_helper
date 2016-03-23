package main

import "fmt"
import "log"

import "github.com/google/go-github/github"

func main() {
  client := github.NewClient(nil)
  orgs, _, err := client.Organizations.List("willnorris", nil)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(orgs)
}
