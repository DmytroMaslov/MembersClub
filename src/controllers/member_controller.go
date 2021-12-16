package controllers

import (
	"encoding/json"
	"fmt"
	"memberclub/src/services"
	"net/http"
	"strings"
)

func AddUser(resp http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		resp.WriteHeader(http.StatusBadRequest)
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
		return
	}
	resp.WriteHeader(http.StatusOK)
}

func GetAllMember(resp http.ResponseWriter, req *http.Request) {
	members := services.GetAllMembers()
	jsonMembers, err := json.Marshal(members)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(jsonMembers)
}
