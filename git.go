package main

import (
  "log"
  "os/exec"
  "strings"
  "github.com/fatih/color"
)

const repoPath = "/Users/gburanov/code/wimdu"

func checkFileExist(fileName string) bool {
  command := exec.Command("test", "-f", fileName)
  command.Dir = repoPath
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

func fileAuthors(fileName string) []string {
  if verbose {
    yellow := color.New(color.FgYellow)
    yellow.Println("Analyzing file ", fileName)
  }
  authors := []string{}
  if checkFileExist(fileName) == false {
    return authors
  }

  command := exec.Command("git", "blame", "--line-porcelain", fileName)
  command.Dir = repoPath

  out, err := command.Output()
  if err != nil {
        red := color.New(color.FgRed)
        red.Println("NOT FOUND!!!2222", fileName)
    log.Fatal(err)
  }
  lines := strings.Split(string(out), "\n")
  for _, line := range lines {
    if strings.Contains(line, "author ") {
      author := strings.TrimPrefix(line, "author ")
      authors = append(authors, author)
    }
  }
  return authors
}
