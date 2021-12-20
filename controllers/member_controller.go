package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/DmytroMaslov/memberclub/log"
	"github.com/DmytroMaslov/memberclub/services"
)

func AddUser(resp http.ResponseWriter, req *http.Request) {
	log.Get().Info("Request:", req.Method, req.RequestURI)
	name := req.URL.Query().Get("name")
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("Request must contain name"))
		log.Get().Error("Responce %v", http.StatusBadRequest)
		return
	}
	email := req.URL.Query().Get("email")
	email = strings.TrimSpace(email)
	if len(email) == 0 {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("Request must contain email"))
		log.Get().Error("Responce %v", http.StatusBadRequest)
		return
	}
	m, err := services.AddMember(name, email)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(fmt.Sprint(err)))
		log.Get().Error("Responce", http.StatusInternalServerError)
		return
	}
	jsonMember, err := json.Marshal(m)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Get().Error("Responce", http.StatusInternalServerError)
		return
	}
	resp.WriteHeader(http.StatusCreated)
	resp.Write(jsonMember)
	log.Get().Info("Responce:", http.StatusCreated, string(jsonMember))

}

func GetAllMember(resp http.ResponseWriter, req *http.Request) {
	log.Get().Info("Request:", req.Method, req.RequestURI)
	members := services.GetAllMembers()
	if len(members) == 0 {
		resp.WriteHeader(http.StatusNoContent)
		log.Get().Info("Responce:", http.StatusNoContent)
		return
	}
	jsonMembers, err := json.Marshal(members)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Get().Error("Responce:", http.StatusInternalServerError)
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(jsonMembers)
	log.Get().Info("Responce:", http.StatusOK, string(jsonMembers))
}
