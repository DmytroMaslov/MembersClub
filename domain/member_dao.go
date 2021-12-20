package domain

import (
	"fmt"
	"sync"
)

type members struct {
	lock    sync.Mutex
	members map[string]*Member
}

var atomicMembers *members

func init() {
	atomicMembers = initMembers()
}

func initMembers() *members {
	return &members{
		members: make(map[string]*Member)}
}

func (ms *members) getMember(k string) (m *Member, st bool) {
	ms.lock.Lock()
	defer ms.lock.Unlock()
	m, st = ms.members[k]
	return m, st
}

func (ms *members) setMember(m *Member) {
	ms.lock.Lock()
	defer ms.lock.Unlock()
	ms.members[m.Email] = m
}

func (ms *members) getAllMembers() []Member {
	ms.lock.Lock()
	defer ms.lock.Unlock()
	var mem []Member
	for _, value := range ms.members {
		mem = append(mem, *value)
	}
	return mem
}

func SaveMember(m *Member) (outM *Member, err error) {
	_, isPresent := atomicMembers.getMember(m.Email)
	if isPresent {
		err = fmt.Errorf("email %v is present in saved members", m.Email)
		return
	}
	atomicMembers.setMember(m)
	outM, _ = atomicMembers.getMember(m.Email)
	return outM, nil
}

func GetAllMembers() []Member {
	return atomicMembers.getAllMembers()
}
