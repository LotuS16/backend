// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"backend/internal/types"
	"github.com/sev-2/raiden/pkg/resource"
)

func RegisterTypes() {
	resource.RegisterTypes(
		&types.AppPermission{},
		&types.AppRole{},
		&types.AppPermission{},
		&types.AppRole{},
		&types.BookingController{},
		&types.Doctor{},
		&types.Faskes{},
		&types.JadwalDoctors{},
		&types.Poli{},
		&types.Reservation{},
		&types.RolePermissions{},
		&types.UserRoles{},
	)
}
