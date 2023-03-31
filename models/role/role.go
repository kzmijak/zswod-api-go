package role

import (
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/utils/array"
	"golang.org/x/exp/slices"
)

const (
	ErrRoleDoesNotExist = "ErrRoleDoesNotExist: Role does not exist."
)

type Role string

const (
	Unknown Role = "Unknown"
	Admin Role = "Admin"
	Teacher Role = "Teacher"
	LegalGuardian Role = "LegalGuardian"
	Student Role = "Student"
)

var authority = []Role{
	Admin, Teacher, LegalGuardian, Student,
}

func (r Role) String() string {
	return string(r)
}

func FromString(roleString string) (Role, error) {
	role := Role(roleString)
	exists := array.Includes(authority, role)
	if !exists {
		return Unknown, errors.Error(ErrRoleDoesNotExist)
	}
	
	return role, nil
}

func (r Role) OrHigher() []Role{
	return authority[:slices.Index(authority, r) + 1]
}