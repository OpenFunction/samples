package hello

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OpenFunction/functions-framework-go/functions"
)

func init() {
	functions.HTTP("Foo", Foo, functions.WithFunctionPath("/foo"))
	functions.HTTP("Bar", Bar, functions.WithFunctionPath("/bar"))

}

func Foo(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"hello": "foo!",
	}
	responseBytes, _ := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.Write(responseBytes)
}

func Bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, bar!\n")
}
