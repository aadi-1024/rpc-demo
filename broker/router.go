package main

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/send-mail", http.HandlerFunc(SendMail))
	mux.Handle("/add-user", http.HandlerFunc(AddUser))
	mux.Handle("/delete-user", http.HandlerFunc(DeleteUser))

	return mux
}
