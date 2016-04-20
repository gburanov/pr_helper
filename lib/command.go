package pr_helper

import (
  "os/exec"
  "fmt"
  "github.com/fatih/color"
)

func (repo *Repository) ExecuteSilently(name string, arg ...string) error {
  command := exec.Command(name, arg...)
  _, error := command.Output()
  return error
}

func (repo *Repository) ExecuteCommand(name string, arg ...string) error {
  return repo.ExecuteCommandInDir(repo.LocalPath(), name, arg...)
}

func (repo *Repository) ExecuteCommandInDir(dir string, name string, arg ...string) error {
  fmt.Println(name, arg)

  command := exec.Command(name, arg...)
  if dir != "" {
	   command.Dir = dir
   }
	out, error := command.Output()
	if error != nil {
    red := color.New(color.FgRed)
    red.Println(error)
    red.Println(out)
    return error
  }
  return nil
}
