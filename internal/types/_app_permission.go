package types

import (
	"github.com/sev-2/raiden"
)

type AppPermission struct {
	raiden.TypeBase
}

func (t *AppPermission) Name() string {
	return "_app_permission"
}

func (r *AppPermission) Format() string {
	return "app_permission[]"
}

func (r *AppPermission) Enums() []string {
	return []string{}
}

func (r *AppPermission) Attributes() []string {
	return []string{}
}

func (r *AppPermission) Comment() *string {
	return nil
}

