package pr_helper

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

const SettingsFile = "settings.yml"

type Settings struct {
	AuthToken      string
	RepositoryPath string
	Organization   string
	Project        string
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

func GetSettings() *Settings {
	settings := new(Settings)
	err := yaml.Unmarshal(readFile(SettingsFile), &settings)
	if err != nil {
		log.Fatal(err)
	}

	// Do some overrides
	settings.AuthToken = os.GetEn("GITHUB_ACCESS_TOKEN")
	return settings
}
