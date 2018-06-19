package router

// TODO : about / http resp editor / proxy

import (
	"net/http"
	"strings"

	"github.com/romainmenke/rapip/proxy"
	"github.com/romainmenke/rapip/store"
)

func New(kvStore store.Store) http.Handler {

	githubGist := proxy.GithubGist(kvStore)
	githubRepo := proxy.GithubRepo(kvStore)

	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.Host, "gist-github") {
			githubGist.ServeHTTP(w, r)
			return
		}

		if strings.HasPrefix(r.Host, "github") {
			githubRepo.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, "https://github.com/romainmenke/rapip", http.StatusPermanentRedirect)
	}))
}
