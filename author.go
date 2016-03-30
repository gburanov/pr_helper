package main

type Author struct {
  Name string
  Email string
}

func (author *Author) asStr() string {
  return author.Name + "<" + author.Email + ">"
}

func (author *Author) available() bool {
  return true
}
