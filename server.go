package main

import (
	"fmt"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("request map:", r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ";"))
	}

	fmt.Fprintln(w, "i am glad to server you.")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		fmt.Println("ListenAndServe", err)
	}
}
