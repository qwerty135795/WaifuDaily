package main

import (
	"Chat/auth"
	"Chat/data"
	"Chat/dto"
	. "Chat/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	data.InitDatabase()
	http.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
		}
		loginReq := &struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}{}
		err := json.NewDecoder(req.Body).Decode(&loginReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		token, err := auth.Login(loginReq.Name, loginReq.Password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}
		json.NewEncoder(w).Encode(token)
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, req *http.Request) {
		var regDto dto.RegisterDto
		err := json.NewDecoder(req.Body).Decode(&regDto)
		if err != nil {
			log.Fatal(err)
		}
		res, err := auth.Register(regDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		responseStruct := struct {
			Id int
		}{Id: res}
		json.NewEncoder(w).Encode(responseStruct)
	})
	http.HandleFunc(`/messages`, SendMessage)
	http.HandleFunc(`/dialog`, GetDialog)
	defer data.Close()
	http.ListenAndServe(":5000", nil)
}

func SendMessage(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var message Message
	err := json.NewDecoder(req.Body).Decode(&message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := data.NewMessage(message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Id: %d", id)
}

func GetDialog(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	qry := req.URL.Query()
	userId, receiverId := qry.Get("userId"), qry.Get("receiverId")
	if userId == "" || receiverId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	model := struct {
		UserId     int
		ReceiverId int
	}{}
	if Id, err := strconv.Atoi(userId); err != nil {
		log.Println(http.StatusBadRequest)
		return
	} else {
		model.UserId = Id
	}
	if Id, err := strconv.Atoi(receiverId); err != nil {
		log.Println(err)
		return
	} else {
		model.ReceiverId = Id
	}
	messages, err := data.GetDialog(model.UserId, model.ReceiverId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	js, err := json.Marshal(messages)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprint(w, string(js))
}
