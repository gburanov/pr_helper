package pr_helper

import (
  "net/http"
  "golang.org/x/oauth2"
)

func Token() *http.Client {
  token_str := GetSettings().AuthToken
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: token_str},
  )
  return oauth2.NewClient(oauth2.NoContext, ts)
}
