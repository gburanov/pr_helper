package main

import (
  "bytes"
  "os"
  "io"
  "log"
  "net/http"
  "golang.org/x/oauth2"
)

const filename = "./auth_token"

func token() *http.Client {
  buf := bytes.NewBuffer(nil)
  f,err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  io.Copy(buf, f)
  f.Close()
  token_str := string(buf.Bytes())
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: token_str},
  )
  return oauth2.NewClient(oauth2.NoContext, ts)
}
