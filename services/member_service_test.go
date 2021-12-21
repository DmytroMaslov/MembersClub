package services

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DmytroMaslov/memberclub/domain"
	"github.com/DmytroMaslov/memberclub/mock"
	"github.com/stretchr/testify/assert"
)

func init() {
	domain.MemberDao = mock.GetMemberDaoMock()
}

func Test_AddUser_Success(t *testing.T) {
	expectedName := "test"
	expectedEmail := "test@test.com"
	mock.SaveMemberFunc = func(m *domain.Member) (outM *domain.Member, err error) {
		return &domain.Member{
			Name:  expectedName,
			Email: expectedEmail,
		}, nil
	}
	m, err := AddMember(expectedName, expectedEmail)
	assert.Equal(t, expectedName, m.Name, "name after saving not same")
	assert.Equal(t, expectedEmail, m.Email, "email after saving not same")
	assert.True(t, time.Now().After(m.RegistationDate), "Registration date in a future")
	assert.Nil(t, err)
}

func Test_AddUser_ValidationFail(t *testing.T) {
	validName := "test"
	invalidEmail := "test"
	m, err := AddMember(validName, invalidEmail)
	assert.Nil(t, m)
	assert.NotNil(t, err)
}

func Test_AddUser_SavingError(t *testing.T) {
	mock.SaveMemberFunc = func(m *domain.Member) (outM *domain.Member, err error) {
		return nil, errors.New("Error")
	}
	name := "test"
	email := "test@test.com"
	m, err := AddMember(name, email)
	assert.Nil(t, m)
	assert.NotNil(t, err)
}

func Test_GetAllMembers(t *testing.T) {
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
	mock.GetAllMemberFunc = func() []domain.Member {
		return []domain.Member{*inputMemberFirst, *inputMemberSecond}
	}
	outMembers := GetAllMembers()
	assert.True(t, reflect.DeepEqual(inputMembers, outMembers))
}
