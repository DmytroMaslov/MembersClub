package controllers

import (
	"encoding/json"
	"fmt"
	logger "memberclub/src/log"
	"memberclub/src/services"
	"net/http"
	"strings"
)

func AddUser(resp http.ResponseWriter, req *http.Request) {
	logger.Info(fmt.Sprintf("Request: %v %v", req.Method, req.RequestURI))
	name := req.URL.Query().Get("name")
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		resp.WriteHeader(http.StatusBadRequest)
		logger.Info(fmt.Sprintf("Request %v", http.StatusBadRequest))
		return
	}
	email := req.URL.Query().Get("email")
	email = strings.TrimSpace(email)
	if len(email) == 0 {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	member := map[string]string{
		"name":  name,
		"email": email,
	}
	err := services.AddMember(member)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(fmt.Sprint(err)))
		logger.Info(fmt.Sprintf("Responce:%v"))
		return
	}
	resp.WriteHeader(http.StatusOK)
}

func GetAllMember(resp http.ResponseWriter, req *http.Request) {
	logger.Info("Incoming request")
	members := services.GetAllMembers()
	jsonMembers, err := json.Marshal(members)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(jsonMembers)
}
