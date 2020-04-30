package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(writter http.ResponseWriter, request *http.Request) {
		writter.Write([]byte("Hello world"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
