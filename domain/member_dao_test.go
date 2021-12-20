package domain

import "testing"

func Test_SaveMember(t *testing.T) {
	mem1 := Member{
		Name:  "Dima",
		Email: "123",
		//RegistationDate: 123,
	}

	mem2 := Member{
		Name:  "NeDima",
		Email: "123",
		//RegistationDate: 123,
	}

	SaveMember(&mem1)
	SaveMember(&mem2)
	t.Error(GetAllMembers())
}
