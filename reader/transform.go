package reader

import (
	"io"
	"net/http"
)

func Transform(req *http.Request, reader io.Reader) (io.Reader, error) {
	// TODO : add support for some level of dynamic responses.

	return reader, nil
}
