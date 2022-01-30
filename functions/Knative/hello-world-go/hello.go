package hello

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
	return nil
}
