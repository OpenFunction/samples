package hello

import (
	"fmt"
	"net/http"
)

// HelloWorld writes "Hello, World!" to the HTTP response.
func HelloWorld(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "Hello, World!\n")
	return nil
}