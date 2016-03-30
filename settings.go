package main

import (
  "bytes"
  "os"
  "io"
  "log"
  "gopkg.in/yaml.v2"
)

const SettingsFile = "settings.yml"

type Settings struct {
  AuthToken string
  RepositoryPath string
}

func readFile(filename string) []byte {
  buf := bytes.NewBuffer(nil)
  f,err := os.Open(filename)
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
  return settings
}
