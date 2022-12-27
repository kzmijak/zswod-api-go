package role

import bimap "github.com/kzmijak/zswod_api_go/utils/bimah"

type Role int 

const (
	Unknown Role = iota - 1
	Admin Role = iota - 1
	Teacher
	LegalGuardian
	Student
)

var roleStringMap = bimap.NewBiMap[Role]().
	WithMember(Admin, "Admin").
	WithMember(Teacher, "Teacher").
	WithMember(LegalGuardian, "LegalGuardian").
	WithMember(Student, "Student").
	WithMember(Unknown, "Unknown")

func (r Role) String() string {
	return roleStringMap.GetByKey(r)
}