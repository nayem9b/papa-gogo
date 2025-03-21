package main

import (
	"fmt"
	"hello/internal/http/handlers/student"
	"net/http"
)

func main() {

	//setup router
	router := http.NewServeMux()
	// router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World!"))
	// })

	router.HandleFunc("GET /api/students",student.New())

	router.HandleFunc("POST /api/students",student.New())
	//start server
	server := http.Server{
		Addr: ":8008",
		Handler: router,
	}

	server.ListenAndServe()
	fmt.Println("Server is running on port 8080")
}