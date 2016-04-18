package pr_helper

import (
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
	"strings"
)

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

/*
var MyEmail = ""
func myEmail() string {
	if MyEmail != "" {
		return MyEmail
	}

	command := exec.Command("git", "config", "user.email")
	command.Dir = GetRepositoryPath()
	out, err := command.Output()
	if err != nil {
		log.Fatal(err)
	}
	MyEmail = strings.TrimSuffix(string(out), "\n")
	return MyEmail
}
*/

func checkFileExist(repo *Repository, fileName string) bool {
	command := exec.Command("test", "-f", fileName)
	command.Dir = repo.LocalPath()
	retCode := command.Run()
	if retCode != nil {
		if GetSettings().Verbosity {
			red := color.New(color.FgRed)
			red.Println(fileName, "not found")
		}
		return false
	}
	return true
}

func fileAuthors(repo *Repository, fileName string) []Author {
	if GetSettings().Verbosity {
		yellow := color.New(color.FgYellow)
		yellow.Println("Analyzing file ", fileName)
	}
	authors := []Author{}
	if checkFileExist(repo, fileName) == false {
		return authors
	}

	command := exec.Command("git", "blame", "--line-porcelain", fileName)
	command.Dir = repo.LocalPath()

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
			author := Author{Name: name, Email: email}
			authors = append(authors, author)
		}
	}
	return authors
}
