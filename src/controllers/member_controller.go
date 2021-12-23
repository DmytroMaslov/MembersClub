package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/DmytroMaslov/memberclub/src/log"
	"github.com/DmytroMaslov/memberclub/src/services"
)

func AddUser(resp http.ResponseWriter, req *http.Request) {
	log.Get().Info("Request method:%v, Uri:%v", req.Method, req.RequestURI)
	name := req.URL.Query().Get("name")
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("Request must contain name"))
		log.Get().Error("Responce: Status:%v", http.StatusBadRequest)
		return
	}
	email := req.URL.Query().Get("email")
	email = strings.TrimSpace(email)
	if len(email) == 0 {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("Request must contain email"))
		log.Get().Error("Responce: Status:%v", http.StatusBadRequest)
		return
	}
	reqistrationTime := time.Now().UTC()
	m, err := services.AddMember(name, email, reqistrationTime)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(fmt.Sprint(err)))
		log.Get().Error("Responce: Status:%v", http.StatusInternalServerError)
		return
	}
	jsonMember, err := json.Marshal(m)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Get().Error("Responce: Status:%v", http.StatusInternalServerError)
		return
	}
	resp.WriteHeader(http.StatusCreated)
	resp.Write(jsonMember)
	log.Get().Info("Responce: Status:%v Data:%v", http.StatusCreated, string(jsonMember))

}

func GetAllMember(resp http.ResponseWriter, req *http.Request) {
	log.Get().Info("Request method:%v, Uri:%v", req.Method, req.RequestURI)
	members := services.GetAllMembers()
	if len(members) == 0 {
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("Any saved members"))
		log.Get().Info("Responce: Status:%v Data:%v", http.StatusOK, "Any saved members")
		return
	}
	jsonMembers, err := json.Marshal(members)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Get().Error("Responce: Status:%v", http.StatusInternalServerError)
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(jsonMembers)
	log.Get().Info("Responce: Status:%v Data:%v", http.StatusOK, string(jsonMembers))
}
