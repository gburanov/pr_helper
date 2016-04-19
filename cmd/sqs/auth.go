package main

import (
  "os"

  "github.com/goamz/goamz/aws"
)

var auth = aws.Auth{
  AccessKey: os.Getenv("AWS_PRIVATE_ACCESS_KEY_ID"),
  SecretKey: os.Getenv("AWS_PRIVATE_SECRET_ACCESS_KEY"),
}
