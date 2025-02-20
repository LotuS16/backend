package types

import (
	"github.com/sev-2/raiden"
)

type UserRoles struct {
	raiden.TypeBase
}

func (t *UserRoles) Name() string {
	return "_user_roles"
}

func (r *UserRoles) Format() string {
	return "user_roles[]"
}

func (r *UserRoles) Enums() []string {
	return []string{}
}

func (r *UserRoles) Attributes() []string {
	return []string{}
}

func (r *UserRoles) Comment() *string {
	return nil
}

