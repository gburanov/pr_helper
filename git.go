package main

import (
  "log"
  "os/exec"
  "strings"
  "github.com/fatih/color"
)

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
  for _, line := range lines {
    if strings.Contains(line, "author ") {
      author := Author{ Name: strings.TrimPrefix(line, "author ") }
      authors = append(authors, author)
    }
  }
  return authors
}
