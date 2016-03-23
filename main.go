package main

import "fmt"
import "log"
import "sort"

import "github.com/google/go-github/github"
//import "github.com/fatih/color"

const organization = "wimdu"
const project = "wimdu"

func main() {
  auth_token := token()
  client := github.NewClient(auth_token)
  listPRs(client)
}

func listPRs(client *github.Client) {
  prs, _, err := client.PullRequests.List(organization, project, nil)
  if err != nil {
    log.Fatal(err)
  }
  for _, pr := range prs {
    processPr(&pr, client)
  }
}

func processPr(pr *github.PullRequest, client *github.Client) {
  fmt.Println(*pr.Title)
  files, _, err := client.PullRequests.ListFiles(organization, project, *pr.Number, nil)
  if err != nil {
    log.Fatal(err)
  }
  authors := []string{}

  for _, file := range files {
    authors = append(authors, fileAuthors(*file.Filename)...)
  }
  arrayToMap(authors)

  log.Fatal(nil)
}

func arrayToMap(authors []string) {
  ret := make(map[string] int )
  for _, author := range authors {
    ret[author] += 1
  }

  reverse := map[int][]string{}
  for k, v := range ret {
    reverse[v] = append(reverse[v], k)
  }

  var a []int
  for k := range reverse {
    a = append(a, k)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(a)))

  for _, key := range a {
    author := reverse[key]
    fmt.Println(author)
  }
}
