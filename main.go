package main

import "net/http"

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, 9!"))
	})
	http.ListenAndServe(":8080", server)
}
