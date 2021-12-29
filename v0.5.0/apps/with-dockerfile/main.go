// From https://github.com/homeport/gonut/tree/master/assets/sample-apps/golang

package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

func main() {
	port := 8080
	if strValue, ok := os.LookupEnv("PORT"); ok {
		if intValue, err := strconv.Atoi(strValue); err == nil {
			port = intValue
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! I am using %s by the way.", runtime.Version())
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
