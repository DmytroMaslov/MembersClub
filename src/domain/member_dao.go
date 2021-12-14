package domain

import "fmt"

var members map[string]*Member

func SaveMember(m *Member) (err error) {
	if members == nil {
		members = initMembers()
	}
	_, isPresent := members[m.Email]
	if isPresent {
		err = fmt.Errorf("key %v is present in saved members", m.Email)
		return
	}
	members[m.Email] = m
	return
}

func GetMemberByEmail(email string) *Member {
	return members[email]
}

func GetAllMembers() map[string]*Member {
	return members
}

func initMembers() map[string]*Member {
	return make(map[string]*Member)
}
