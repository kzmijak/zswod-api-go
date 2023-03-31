package role

import (
	"github.com/kzmijak/zswod_api_go/modules/errors"
	bimap "github.com/kzmijak/zswod_api_go/utils/bimap"
	"golang.org/x/exp/slices"
)

const (
	ErrRoleDoesNotExist = "ErrRoleDoesNotExist: Role does not exist."
)

type Role int 

const (
	Unknown Role = iota
	Admin
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

func FromString(roleString string) (Role, error) {
	role, exists := roleStringMap.GetByValue(roleString)
	if !exists {
		return Unknown, errors.Error(ErrRoleDoesNotExist)
	}
	
	return role, nil
}

func FromId(roleId int) (Role, error) {
	exists := roleStringMap.GetLength() > roleId
	if !exists {
		return Unknown, errors.Error(ErrRoleDoesNotExist)
	}

	return Role(roleId), nil
}

func (r Role) OrHigher() []Role{
	return authority[:slices.Index(authority, r) + 1]
}