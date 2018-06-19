package proxy

import (
	"context"
	"net/http"
	"path"
	"time"

	"github.com/romainmenke/rapip/reader"
	"github.com/romainmenke/rapip/store"
)

func GithubGist(kvStore store.Store) http.Handler {

	client := http.Client{
		Timeout: time.Second * 30,
		Jar:     nil, // explicit nil, this is a shared client!!
	}

	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
		defer cancel()

		req, err := http.NewRequest("GET", "https://"+path.Join("gist.githubusercontent.com/", r.URL.Path), nil)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		req.WithContext(ctx)

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		reader.Respond(ctx, w, r, resp, kvStore)
	}))
}
