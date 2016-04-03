package pr_helper

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
)

type Author struct {
	Name  string
	Email string
}

func (author *Author) AsStr() string {
	return author.Name + " <" + author.Email + ">"
}

var Blacklist = map[string]bool{}
var Whitelist = map[string]bool{}

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
	}
	return Blacklist
}

func whitelist() map[string]bool {
	if len(Whitelist) == 0 {
		content, err := ioutil.ReadFile("whitelist")
		if err != nil {
			log.Fatal(err)
		}
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			Whitelist[line] = true
		}
	}
	return Whitelist
}

func (author *Author) Check() bool {
	if !blacklist()[author.Email] && !whitelist()[author.Email] {
		fmt.Println(author.Email)
		return false
	}
	return true
}

func (author *Author) AtWimdu() bool {
	return !blacklist()[author.Email]
}

func (author *Author) filtered() bool {
	return !author.AtWimdu() || author.Email == myEmail()
}
