package main

import (
	"fmt"
	"html"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "WelCome to My Host! %q\n", html.EscapeString(r.URL.Path))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("http server start err, %v\n", err)
	}
}
