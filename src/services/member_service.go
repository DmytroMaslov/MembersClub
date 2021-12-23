package services

import (
	"time"

	"github.com/DmytroMaslov/memberclub/src/domain"
	"github.com/DmytroMaslov/memberclub/src/validators"
)

func AddMember(name string, email string, regTime time.Time) (m *domain.Member, err error) {
	newMember := domain.Member{
		Name:            name,
		Email:           email,
		RegistationDate: regTime,
	}
	err = validators.Validate(newMember)
	if err != nil {
		return
	}
	m, err = domain.MemberDao.SaveMember(&newMember)
	if err != nil {
		return
	}
	return
}

func GetAllMembers() []domain.Member {
	return domain.MemberDao.GetAllMembers()
}
