package main

import (
  "log"
  "os/exec"
  "strings"
  "github.com/fatih/color"
)

var MyEmail = ""

func myEmail() string {
  if MyEmail != "" {
    return MyEmail
  }

  command := exec.Command("git", "config", "user.email")
  command.Dir = GetSettings().RepositoryPath
  out, err := command.Output()
  if err != nil {
    log.Fatal(err)
  }
  MyEmail = strings.TrimSuffix(string(out), "\n")
  return MyEmail
}

func checkFileExist(fileName string) bool {
  command := exec.Command("test", "-f", fileName)
  command.Dir = GetSettings().RepositoryPath
  retCode := command.Run()
  if retCode != nil {
    if verbose {
      red := color.New(color.FgRed)
      red.Println(fileName, "not found")
    }
    return false
  }
  return true
}

func fileAuthors(fileName string) []Author {
  if verbose {
    yellow := color.New(color.FgYellow)
    yellow.Println("Analyzing file ", fileName)
  }
  authors := []Author{}
  if checkFileExist(fileName) == false {
    return authors
  }

  command := exec.Command("git", "blame", "--line-porcelain", fileName)
  command.Dir = GetSettings().RepositoryPath

  out, err := command.Output()
  if err != nil {
    log.Fatal(err)
  }
  lines := strings.Split(string(out), "\n")

  name := ""
  for _, line := range lines {
    if strings.Contains(line, "author ") {
      name = strings.TrimPrefix(line, "author ")
    }
    if strings.Contains(line, "author-mail <") {
      email := strings.TrimSuffix(strings.TrimPrefix(line, "author-mail <"), ">")
      author := Author{ Name: name, Email: email }
      if author.available() {
        authors = append(authors, author)
      }
    }
  }
  return authors
}
