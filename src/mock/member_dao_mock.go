package mock

import "github.com/DmytroMaslov/memberclub/src/domain"

type membersDaoMock struct {
}

var (
	memberDaoMock    membersDaoMock
	SaveMemberFunc   func(m *domain.Member) (outM *domain.Member, err error)
	GetAllMemberFunc func() []domain.Member
)

func GetMemberDaoMock() *membersDaoMock {
	return &memberDaoMock
}

func (ms *membersDaoMock) SaveMember(m *domain.Member) (outM *domain.Member, err error) {
	return SaveMemberFunc(m)
}

func (ms *membersDaoMock) GetAllMembers() []domain.Member {
	return GetAllMemberFunc()
}
