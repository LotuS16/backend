package types

import (
	"github.com/sev-2/raiden"
)

type JadwalDoctors struct {
	raiden.TypeBase
}

func (t *JadwalDoctors) Name() string {
	return "_jadwal_doctors"
}

func (r *JadwalDoctors) Format() string {
	return "jadwal_doctors[]"
}

func (r *JadwalDoctors) Enums() []string {
	return []string{}
}

func (r *JadwalDoctors) Attributes() []string {
	return []string{}
}

func (r *JadwalDoctors) Comment() *string {
	return nil
}

