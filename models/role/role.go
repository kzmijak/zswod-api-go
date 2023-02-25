package role

import (
	bimap "github.com/kzmijak/zswod_api_go/utils/bimap"
	"golang.org/x/exp/slices"
)

type Role int 

const (
	Unknown Role = iota - 1
	Admin Role = iota - 1
	Teacher
	LegalGuardian
	Student
)

var authority = []Role{
	Admin, Teacher, LegalGuardian, Student, Unknown,
}

var roleStringMap = bimap.NewBiMap[Role]().
	WithMember(Admin, "Admin").
	WithMember(Teacher, "Teacher").
	WithMember(LegalGuardian, "LegalGuardian").
	WithMember(Student, "Student").
	WithMember(Unknown, "Unknown")

func (r Role) String() string {
	return roleStringMap.GetByKey(r)
}

func FromString(role string) (Role, bool) {
	return roleStringMap.GetByValue(role)
}

func FromId(id int) (role Role, exists bool) {
	role = authority[id]
	exists = len(authority) > id
	return
}

func (r Role) OrHigher() []Role{
	return authority[:slices.Index(authority, r) + 1]
}