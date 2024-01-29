package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserModel struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type MailPayload struct {
	Email   string
	Content string
}

func SendMail(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//app.CacheConn.Call()
	m := MailPayload{
		Email:   "",
		Content: "Hey this is a test mail",
	}
	var b bool
	err := app.MailConn.Call("MailRPC.SendMail", m, &b)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(200)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u := UserModel{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
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

//func GetAll(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "POST" {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	m := make(map[int]string)
//
//	err := app.CacheConn.Call("CacheRPC.GetAll", 0, &m)
//	if err != nil {
//		w.WriteHeader(400)
//	} else {
//		fmt.Println(m)
//		w.WriteHeader(200)
//	}
//}
