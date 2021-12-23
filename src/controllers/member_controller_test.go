package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/DmytroMaslov/memberclub/src/domain"
	"github.com/DmytroMaslov/memberclub/src/mock"
	"github.com/stretchr/testify/assert"
)

func init() {
	domain.MemberDao = mock.GetMemberDaoMock()
}

func Test_AddUser_Success(t *testing.T) {
	inputQuery := "name=Test&email=test@test.com"
	mockedMember := &domain.Member{
		Name:            "test",
		Email:           "test@test.com",
		RegistationDate: time.Now(),
	}
	jsonMockedMember, _ := json.Marshal(mockedMember)
	mock.SaveMemberFunc = func(m *domain.Member) (outM *domain.Member, err error) {
		return mockedMember, nil
	}
	mock.GetMemberDaoMock().SaveMember(mockedMember)
	request := &http.Request{
		URL: &url.URL{
			RawQuery: inputQuery,
		},
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddUser)
	handler.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, string(jsonMockedMember), rr.Body.String())
}

func Test_Add_User_SavedError(t *testing.T) {
	inputQuery := "name=Test&email=test@test.com"
	errorMessage := "email test@test.com is present in saved members"
	mock.SaveMemberFunc = func(m *domain.Member) (outM *domain.Member, err error) {
		return nil, errors.New(errorMessage)
	}
	mockedMember := &domain.Member{
		Name:            "test",
		Email:           "test@test.com",
		RegistationDate: time.Now(),
	}
	mock.GetMemberDaoMock().SaveMember(mockedMember)
	request := &http.Request{
		URL: &url.URL{
			RawQuery: inputQuery,
		},
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddUser)
	handler.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, errorMessage, rr.Body.String())
}

func Test_AddUser_WithoutName(t *testing.T) {
	inputQuery := "email=test@test.com"
	request := &http.Request{
		URL: &url.URL{
			RawQuery: inputQuery,
		},
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddUser)
	handler.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Request must contain name", rr.Body.String())
}

func Test_AddUser_WithoutEmail(t *testing.T) {
	inputQuery := "name=Test"
	request := &http.Request{
		URL: &url.URL{
			RawQuery: inputQuery,
		},
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddUser)
	handler.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Request must contain email", rr.Body.String())
}

func Test_GetAllMember_Success(t *testing.T) {
	inputMemberFirst := &domain.Member{
		Name:            "First",
		Email:           "first@test.com",
		RegistationDate: time.Now(),
	}
	inputMemberSecond := &domain.Member{
		Name:            "Second",
		Email:           "second@test.com",
		RegistationDate: time.Now(),
	}
	inputMembers := []domain.Member{*inputMemberFirst, *inputMemberSecond}
	jsonMockedMember, _ := json.Marshal(inputMembers)
	mock.GetAllMemberFunc = func() []domain.Member {
		return inputMembers
	}
	mock.GetMemberDaoMock().GetAllMembers()
	request := &http.Request{}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllMember)
	handler.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, string(jsonMockedMember), rr.Body.String())
}

func Test_GetAllMember_EmptyStorage(t *testing.T) {
	mock.GetAllMemberFunc = func() []domain.Member {
		return make([]domain.Member, 0)
	}
	mock.GetMemberDaoMock().GetAllMembers()
	request := &http.Request{}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllMember)
	handler.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Any saved members", rr.Body.String())
}
