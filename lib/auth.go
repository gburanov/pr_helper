package pr_helper

import (
	"net/http"

	"golang.org/x/oauth2"
)

func token() *http.Client {
	tokenStr := getSettings().AuthToken
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tokenStr},
	)
	return oauth2.NewClient(oauth2.NoContext, ts)
}
