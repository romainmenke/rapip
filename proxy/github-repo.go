package proxy

import (
	"context"
	"net/http"
	"path"
	"time"

	"github.com/romainmenke/rapip/reader"
)

func GithubRepo() http.Handler {

	client := http.Client{
		Timeout: time.Second * 30,
		Jar:     nil, // explicit nil, this is a shared client!!
	}

	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
		defer cancel()

		req, err := http.NewRequest("GET", "https://"+path.Join("raw.githubusercontent.com/", r.URL.Path, r.Method), nil)
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

		reader.Respond(w, r, resp)
	}))
}
