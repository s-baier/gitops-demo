package main

import (
	"fmt"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Hello(name string) string {
	if name == "" {
		return "Hello, stranger! Are you lost?!"
	}

	message := fmt.Sprintf("Hello %v, welcome to the GitOps demo!", name)
	return message
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	message := Hello(name)

	fmt.Fprintln(w, message)
}

func main() {
	port := getEnv("PORT", "8080")
	http.HandleFunc("/", helloHandler)

	fmt.Printf("Started, serving at %v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
