package reader

import (
	"bufio"
	"io"
	"log"
	"net/http"
)

func Respond(w http.ResponseWriter, r *http.Request, source *http.Response) {
	if source.StatusCode/100 != 2 {
		http.Error(w, http.StatusText(source.StatusCode), source.StatusCode)
		return
	}

	defer source.Body.Close()

	transform, err := Transform(r, source.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" : "+err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := http.ReadResponse(bufio.NewReader(transform), r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" : "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer output.Body.Close()

	copyHeader(w.Header(), output.Header)
	w.WriteHeader(output.StatusCode)

	_, err = io.Copy(w, output.Body)
	if err != nil {
		log.Println(err)
		return
	}

	err = source.Body.Close()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" : "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = output.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
