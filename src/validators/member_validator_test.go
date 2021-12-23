package validators

import (
	"testing"
	"time"

	"github.com/DmytroMaslov/memberclub/src/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Validate_ValidMember(t *testing.T) {
	validMember := domain.Member{
		Name:            "Test",
		Email:           "test@test.com",
		RegistationDate: time.Now(),
	}
	res := Validate(validMember)
	assert.Nil(t, res)
}

func Test_Validate_InvalidName(t *testing.T) {
	expectedErrorMessage := "field Name must contains English letters, dots, spaces"
	invalidNameMember := domain.Member{
		Name:            "123",
		Email:           "test@test.com",
		RegistationDate: time.Now(),
	}
	res := Validate(invalidNameMember)
	assert.NotNil(t, res)
	assert.Equal(t, expectedErrorMessage, res.Error())
}

func Test_Validate_InvalidEmail(t *testing.T) {
	expectedErrorMessage := "field Email must contains valid email format"
	invalidEmailMember := domain.Member{
		Name:            "Test",
		Email:           "test",
		RegistationDate: time.Now(),
	}
	res := Validate(invalidEmailMember)
	assert.NotNil(t, res)
	assert.Equal(t, expectedErrorMessage, res.Error())
}

func Test_Validate_InvalidTime(t *testing.T) {
	expectedErrorMessage := "field RegistationDate must contains creative data less than now"
	invalidEmailMember := domain.Member{
		Name:            "Test",
		Email:           "test@test.com",
		RegistationDate: time.Now().Add(time.Hour),
	}
	res := Validate(invalidEmailMember)
	assert.NotNil(t, res)
	assert.Equal(t, expectedErrorMessage, res.Error())
}
