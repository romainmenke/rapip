package proxy

import (
	"bufio"
	"context"
	"io"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

// TODO : tailored http client

func New() http.Handler {

	router := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
		defer cancel()

		// TODO : this is wonkey
		urlStr := r.URL.String()
		parts := strings.Split(urlStr, "gist-githubusercontent-com/")
		if len(parts) != 2 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		urlPath := parts[1]

		req, err := http.NewRequest("GET", "https://"+path.Join("gist.githubusercontent.com/", urlPath), nil)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		req.WithContext(ctx)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if resp.StatusCode/100 != 2 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		realResp, err := http.ReadResponse(bufio.NewReader(resp.Body), r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		defer realResp.Body.Close()

		copyHeader(w.Header(), realResp.Header)
		w.WriteHeader(realResp.StatusCode)

		_, err = io.Copy(w, realResp.Body)
		if err != nil {
			log.Println(err)
			return
		}

		err = resp.Body.Close()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = realResp.Body.Close()
		if err != nil {
			log.Println(err)
			return
		}

	}))

	return router
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
