package pr_helper

import (
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
	"strings"
	"strconv"
	"time"
)

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

func checkFileExist(repo *Repository, fileName string) bool {
	error := repo.ExecuteSilently("test", "-f", fileName)
	if error != nil {
		if GetSettings().Verbosity {
			red := color.New(color.FgRed)
			red.Println(fileName, "not found")
		}
		return false
	}
	return true
}

func fileStatistics(repo *Repository, fileName string) []Stat {
	if GetSettings().Verbosity {
		yellow := color.New(color.FgYellow)
		yellow.Println("Analyzing file ", fileName)
	}
	stats := []Stat{}
	if checkFileExist(repo, fileName) == false {
		return stats
	}

	command := exec.Command("git", "blame", "--line-porcelain", fileName)
	command.Dir = repo.LocalPath()

	out, err := command.Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(out), "\n")

	name := ""
	email := ""
	for _, line := range lines {
		if strings.Contains(line, "author ") {
			name = strings.TrimPrefix(line, "author ")
		}
		if strings.Contains(line, "author-mail <") {
			email = strings.TrimSuffix(strings.TrimPrefix(line, "author-mail <"), ">")
		}
		if (strings.Contains(line, "author-time ")) {
			time_as_str := strings.TrimPrefix(line, "author-time ")
			time_as_int, err := strconv.ParseInt(time_as_str, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			author := Author{Name: name, Email: email}
			stat := Stat{Author: author, Time: time.Unix(time_as_int, 0)}
			stats = append(stats, stat)
		}
	}
	return stats
}
