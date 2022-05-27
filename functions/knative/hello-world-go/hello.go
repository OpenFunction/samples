package hello

import (
	"fmt"
	"net/http"

	"github.com/OpenFunction/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloWorld", HelloWorld)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}
