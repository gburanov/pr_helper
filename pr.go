package pr_helper

import (
  "log"
  "strings"
  "github.com/fatih/color"
)

type PR struct {
  Repository *Repository
  Number int
}

func (pr *PR) ShowInfo() {
  pr_, _, err := pr.Repository.Client.PullRequests.
    Get(pr.Repository.Organization, pr.Repository.Project, pr.Number)
  if err != nil {
    log.Fatal(err)
  }
  red := color.New(color.FgRed, color.Bold)
  red.Println(*pr_.Title, "#", pr.Number)
}

func (pr *PR) Authors() *Authors {
  pr.ShowInfo()
  files, _, err := pr.Repository.Client.PullRequests.
    ListFiles(pr.Repository.Organization, pr.Repository.Project, pr.Number, nil)
  if err != nil {
    log.Fatal(err)
  }

  results := 0
  authors_channel := make(chan []Author)
  for _, file := range files {
    if strings.HasPrefix(*file.Filename, "phrase") {
      continue
    }
    results++
    go func(fileName string) {
      authors_channel <- fileAuthors(fileName)
    }(*file.Filename)
  }
  authors := []Author{}
  for i := 1; i <= results; i++ {
    authors = append(authors, <-authors_channel...)
  }
  return arrayToMap(authors)
}
