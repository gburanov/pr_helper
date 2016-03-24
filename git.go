package main

import "os/exec"
//import "log"
import "strings"

import "github.com/fatih/color"

const repoPath = "/Users/gburanov/code/wimdu"

func fileAuthors(fileName string) []string {
  yellow := color.New(color.FgYellow)
  yellow.Println("Analyzing file ", fileName)
  command := exec.Command("git", "blame", "--line-porcelain", fileName)
  command.Dir = repoPath
  authors := []string{}

  out, err := command.Output()
  if err != nil {
    // assume this is because file does not exist before
    //log.Fatal(err)
    return authors
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
