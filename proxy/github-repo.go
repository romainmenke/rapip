package proxy

import (
	"context"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/romainmenke/rapip/reader"
	"github.com/romainmenke/rapip/store"
)

func GithubRepo(kvStore store.Store) http.Handler {

	client := http.Client{
		Timeout: time.Second * 30,
		Jar:     nil, // explicit nil, this is a shared client!!
	}

	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
		defer cancel()

		p := r.URL.Path
		pathComponents := strings.Split(p, "/")
		if len(pathComponents) >= 3 {
			ctx = store.ContextWithPath(ctx, strings.Join(pathComponents[:3], "/"))
		}

		req, err := http.NewRequest("GET", "https://"+path.Join("raw.githubusercontent.com/", p, r.Method), nil)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError)+" : "+err.Error(), http.StatusInternalServerError)
			return
		}

		req.WithContext(ctx)

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError)+" : "+err.Error(), http.StatusInternalServerError)
			return
		}

		reader.Respond(ctx, w, r, resp, kvStore)
	}))
}
