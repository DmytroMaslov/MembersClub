package domain

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_SaveMember_success(t *testing.T) {
	inputMember := &Member{
		Name:            "Test",
		Email:           "test@test.com",
		RegistationDate: time.Now(),
	}
	outMember, err := MemberDao.SaveMember(inputMember)
	assert.Equal(t, inputMember, outMember, "input and output Member not equal")
	assert.Nil(t, err, "error not nil")
	memberInMemory := atomicMembers.members[inputMember.Email]
	assert.Equal(t, inputMember, memberInMemory, "input member and member in memory not same")
}

func Test_SaveMember_fail(t *testing.T) {
	inputMember := &Member{
		Name:            "Test",
		Email:           "test@test.com",
		RegistationDate: time.Now(),
	}
	MemberDao.SaveMember(inputMember)
	_, err := MemberDao.SaveMember(inputMember)
	assert.NotNil(t, err, "Expect error if try save Member with existing email")
}

func Test_GetAllMembers(t *testing.T) {
	inputMemberFirst := &Member{
		Name:            "First",
		Email:           "first@test.com",
		RegistationDate: time.Now(),
	}
	inputMemberSecond := &Member{
		Name:            "Second",
		Email:           "second@test.com",
		RegistationDate: time.Now(),
	}
	inputMembers := []Member{*inputMemberFirst, *inputMemberSecond}
	atomicMembers = nil
	atomicMembers = initMembers()
	MemberDao.SaveMember(inputMemberFirst)
	MemberDao.SaveMember(inputMemberSecond)
	outMembers := MemberDao.GetAllMembers()
	assert.True(t, reflect.DeepEqual(inputMembers, outMembers))
}
