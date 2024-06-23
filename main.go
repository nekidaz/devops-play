package main

import "net/http"

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, pipiska!"))
	})
	http.ListenAndServe(":8080", server)
}
