package types

import (
	"github.com/sev-2/raiden"
)

type RolePermissions struct {
	raiden.TypeBase
}

func (t *RolePermissions) Name() string {
	return "_role_permissions"
}

func (r *RolePermissions) Format() string {
	return "role_permissions[]"
}

func (r *RolePermissions) Enums() []string {
	return []string{}
}

func (r *RolePermissions) Attributes() []string {
	return []string{}
}

func (r *RolePermissions) Comment() *string {
	return nil
}

