package reader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/romainmenke/rapip/store"
)

const templatePrefix = `TEMPLATE
`

func Transform(ctx context.Context, r *http.Request, reader io.Reader, kvStore store.Store) (io.Reader, error) {
	storePath, _ := store.Path(ctx)

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if len(b) < 9 || string(b[:9]) != templatePrefix {
		return bytes.NewBuffer(b), nil
	}

	b = b[9:]

	keyValues := map[string]map[string]interface{}{
		"request": map[string]interface{}{},
		"headers": map[string]interface{}{},
		"cookies": map[string]interface{}{},
		"store":   kvStore.Get(ctx, storePath),
	}

	err = r.ParseForm()
	if err != nil {
		return nil, err
	}

	formValues := r.Form
	for k, vv := range formValues {
		if len(vv) == 1 {
			keyValues["request"][k] = vv[0]
		} else {
			keyValues["request"][k] = vv
		}
	}

	for _, c := range r.Cookies() {
		keyValues["cookies"][c.Name] = c.Value
	}

	for k, vv := range r.Header {
		if len(vv) == 1 {
			keyValues["headers"][k] = vv[0]
		} else {
			keyValues["headers"][k] = vv
		}
	}

	t := template.New("")
	t.Funcs(map[string]interface{}{
		"set": setter(ctx, kvStore, storePath),
	})
	t, err = t.Parse(string(b))
	if err != nil {
		return nil, err
	}

	outBuf := bytes.NewBuffer(nil)
	err = t.Execute(outBuf, keyValues)
	if err != nil {
		return nil, err
	}

	return outBuf, nil
}

func setter(ctx context.Context, kvStore store.Store, path string) func(key string, value interface{}) {
	return func(key string, value interface{}) {
		if path == "" {
			return
		}

		err := kvStore.Put(ctx, path, key, value)
		if err != nil {
			fmt.Println(err)
		}
	}
}
