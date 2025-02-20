package types

import (
	"github.com/sev-2/raiden"
)

type AppRole struct {
	raiden.TypeBase
}

func (t *AppRole) Name() string {
	return "_app_role"
}

func (r *AppRole) Format() string {
	return "app_role[]"
}

func (r *AppRole) Enums() []string {
	return []string{}
}

func (r *AppRole) Attributes() []string {
	return []string{}
}

func (r *AppRole) Comment() *string {
	return nil
}

