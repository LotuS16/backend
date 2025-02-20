package controllers

import (
	"backend/internal/models"

	"github.com/sev-2/raiden"
)

type JadwalDoctorsController struct {
	raiden.ControllerBase
	Http  string `path:"/jadwal_doctors" type:"rest"`
	Model models.JadwalDoctors
}
