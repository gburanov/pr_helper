package main

import (
  "log"
  "strings"
  "io/ioutil"
)

type Author struct {
  Name string
  Email string
}

func (author *Author) asStr() string {
  return author.Name + "<" + author.Email + ">"
}

var Blacklist = map[string]bool{}
func blacklist() map[string]bool {
  if len(Blacklist) == 0 {
    content, err := ioutil.ReadFile("blacklist")
    if err != nil {
      log.Fatal(err)
    }
    lines := strings.Split(string(content), "\n")
    for _, line := range lines {
      Blacklist[line] = true
    }
    Blacklist[myEmail()] = true
  }
  return Blacklist
}

func (author *Author) available() bool {
  return !blacklist()[author.Email]
}
