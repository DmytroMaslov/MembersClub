package services

import (
	"time"

	"github.com/DmytroMaslov/memberclub/domain"
	"github.com/DmytroMaslov/memberclub/validators"
)

func AddMember(name string, email string) (m *domain.Member, err error) {
	newMember := domain.Member{
		Name:            name,
		Email:           email,
		RegistationDate: time.Now().UTC(),
	}
	err = validators.Validate(newMember)
	if err != nil {
		return
	}
	m, err = domain.SaveMember(&newMember)
	if err != nil {
		return
	}
	return
}

func GetAllMembers() []domain.Member {
	return domain.GetAllMembers()
}
