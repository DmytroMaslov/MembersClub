package domain

import (
	"fmt"
	"sync"
)

var (
	atomicMembers *members
	MemberDao     MemberDaoInterface
)

type MemberDaoInterface interface {
	SaveMember(m *Member) (outM *Member, err error)
	GetAllMembers() []Member
}

type members struct {
	lock    sync.Mutex
	members map[string]*Member
}

func init() {
	atomicMembers = initMembers()
	MemberDao = &members{}

}

func initMembers() *members {
	return &members{
		members: make(map[string]*Member)}
}

func (ms *members) getMember(k string) (m *Member, exist bool) {
	m, exist = ms.members[k]
	return m, exist
}

func (ms *members) setMember(m *Member) {
	ms.members[m.Email] = m
}

func (ms *members) getAllMembers() []Member {
	var mem []Member
	for _, value := range ms.members {
		mem = append(mem, *value)
	}
	return mem
}

func (ms *members) SaveMember(m *Member) (outM *Member, err error) {
	//Lock for transaction
	ms.lock.Lock()
	defer ms.lock.Unlock()
	_, isPresent := atomicMembers.getMember(m.Email)
	if isPresent {
		err = fmt.Errorf("email %v is present in saved members", m.Email)
		return
	}
	atomicMembers.setMember(m)
	outM, _ = atomicMembers.getMember(m.Email)
	return outM, nil
}

func (ms *members) GetAllMembers() []Member {
	//Lock for transaction
	ms.lock.Lock()
	defer ms.lock.Unlock()
	return atomicMembers.getAllMembers()
}
