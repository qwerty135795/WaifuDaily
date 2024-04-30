package main

import (
	"Chat/auth"
	"Chat/data"
	"Chat/dto"
	"encoding/json"
	"log"
	"net/http"
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

	http.ListenAndServe(":5000", nil)
}
