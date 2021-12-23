package app

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DmytroMaslov/memberclub/src/controllers"
	"github.com/stretchr/testify/assert"
)

func Test_AddMember(t *testing.T) {
	name := "test"
	email := "test@test.com"
	uri := fmt.Sprintf("/addMember?name=%v&email=%v", name, email)
	req, err := http.NewRequest("Get", uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.AddUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_GetAllMember(t *testing.T) {
	uri := "/getAllMember"
	req, err := http.NewRequest("Get", uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetAllMember)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
