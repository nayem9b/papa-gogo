package main

import (
	"fmt"
	"net/http"
)

func main() {

	//setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	//start server
	server := http.Server{
		Addr: ":8008",
		Handler: router,
	}

	server.ListenAndServe()
	fmt.Println("Server is running on port 8080")
}