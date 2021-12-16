package services

import (
	"memberclub/src/domain"
	"time"
)

func AddMember(member map[string]string) (err error) {
	var newMember domain.Member
	newMember.Name = member["name"]
	newMember.Email = member["email"]
	newMember.RegistationDate = time.Now().Unix()
	err = newMember.Validate()
	if err != nil {
		return
	}
	err = domain.SaveMember(&newMember)
	if err != nil {
		return
	}
	return
}

func GetAllMembers() []domain.Member {
	return domain.GetAllMembers()
}
