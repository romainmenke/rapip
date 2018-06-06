package router

// TODO : about / http resp editor / proxy

import (
	"net/http"
	"strings"

	"github.com/romainmenke/rapip/proxy"
)

func New() http.Handler {

	githubGist := proxy.GithubGist()

	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Host, "gist.github") {
			githubGist.ServeHTTP(w, r)
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}))
}
