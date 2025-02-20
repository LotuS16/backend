package types

import (
	"github.com/sev-2/raiden"
)

type AppPermission struct {
	raiden.TypeBase
}

func (t *AppPermission) Name() string {
	return "app_permission"
}

func (r *AppPermission) Format() string {
	return "app_permission"
}

func (r *AppPermission) Enums() []string {
	return []string{"doctor.select","doctor.delete","doctor.insert","doctor.update","jadwal_doctors.update","jadwal_doctors.delete","jadwal_doctors.insert","jadwal_doctors.select"}
}

func (r *AppPermission) Attributes() []string {
	return []string{}
}

func (r *AppPermission) Comment() *string {
	return nil
}

