package pr_helper

import (
	"bytes"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const settingsFile = "settings.yml"

type settings struct {
	AuthToken      string
	RepositoryPath string
	Verbosity      bool
	Label          string
}

func readFile(filename string) []byte {
	buf := bytes.NewBuffer(nil)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(buf, f)
	f.Close()
	return buf.Bytes()
}

func getSettings() *settings {
	settings := new(settings)
	err := yaml.Unmarshal(readFile(settingsFile), &settings)
	if err != nil {
		log.Fatal(err)
	}
	// Do some overrides
	if len(settings.AuthToken) == 0 {
		settings.AuthToken = os.Getenv("GITHUB_ACCESS_TOKEN")
	}
	if len(settings.RepositoryPath) == 0 {
		settings.RepositoryPath = "repository"
	}
	return settings
}
