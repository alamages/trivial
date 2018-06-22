package main

import (
	"net/http"
	"os"

	"github.com/alamages/trivial/hello"
)

func main() {
	http.HandleFunc("/", hello.Handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
