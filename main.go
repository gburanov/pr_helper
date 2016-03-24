package main

import "github.com/google/go-github/github"

const organization = "wimdu"
const project = "wimdu"

func main() {
  auth_token := token()
  client := github.NewClient(auth_token)
  listPRs(client)
}
