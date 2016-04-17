package pr_helper

import (
	"github.com/fatih/color"
	"log"
	"os"
	"fmt"
	"os/exec"
	"strings"
)

func CreateRepository() {
	err := os.MkdirAll(GetSettings().RepositoryPath, 0777)
	if err != nil {
		log.Fatal(err)
	}
	path := fmt.Sprintf("https://%s@github.com/%s/%s.git",
		GetSettings().AuthToken, GetSettings().Organization, GetSettings().Project)
	fmt.Println("Clonning", path)
	command := exec.Command("git", "clone", path, ".")
	command.Dir = GetSettings().RepositoryPath
	err := command.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

func GetRepositoryPath() string {
	path := GetSettings().RepositoryPath
	exist, err := exists(path)
	if err != nil {
		log.Fatal(err)
	}
	if exist {
		return path
	}
	CreateRepository()
	return path
}

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

func checkFileExist(fileName string) bool {
	command := exec.Command("test", "-f", fileName)
	command.Dir = GetRepositoryPath()
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

func fileAuthors(fileName string) []Author {
	if GetSettings().Verbosity {
		yellow := color.New(color.FgYellow)
		yellow.Println("Analyzing file ", fileName)
	}
	authors := []Author{}
	if checkFileExist(fileName) == false {
		return authors
	}

	command := exec.Command("git", "blame", "--line-porcelain", fileName)
	command.Dir = GetRepositoryPath()

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
