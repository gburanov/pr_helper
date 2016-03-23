package main

import "fmt"
import "os/exec"
import "log"
import "strings"

func fileAuthors(fileName string) []string {
  fmt.Println("Analyzing file ", fileName)
  command := exec.Command("git", "blame", "--line-porcelain", fileName)
  command.Dir = "/Users/gburanov/code/wimdu"

  out, err := command.Output()
  if err != nil {
    log.Fatal(err)
  }
  authors := []string{}

  lines := strings.Split(string(out), "\n")
  for _, line := range lines {
    if strings.Contains(line, "author ") {
      author := strings.TrimPrefix(line, "author ")
      fmt.Println(author)
    }
  }
  log.Fatal(nil)

  return authors
}
