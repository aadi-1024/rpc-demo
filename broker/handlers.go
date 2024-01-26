package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type UserModel struct {
	Id    int
	Email string
}

func SendMail(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//app.CacheConn.Call()
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u := UserModel{
		Id:    rand.Int(),
		Email: "test@email.com",
	}
	var s string
	err := app.CacheConn.Call("CacheRPC.AddUser", u, &s)
	if err != nil {
		http.Error(w, err.Error(), 400)
	} else {
		w.Write([]byte(s))
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m := make(map[int]string)

	err := app.CacheConn.Call("CacheRPC.GetAll", 0, &m)
	if err != nil {
		w.WriteHeader(400)
	} else {
		fmt.Println(m)
		w.WriteHeader(200)
	}
}
