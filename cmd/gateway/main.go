package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

)

func main() {

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World")
	})

}
