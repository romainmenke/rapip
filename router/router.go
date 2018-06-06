package router

// TODO : about / http resp editor / proxy

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/romainmenke/rapip/proxy"
)

func New() http.Handler {

	githubGist := proxy.GithubGist()

	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Host)
		if strings.HasPrefix(r.Host, "gist.github") {
			githubGist.ServeHTTP(w, r)
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}))
}
