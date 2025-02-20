package controllers

import (
	"backend/internal/models"

	"github.com/sev-2/raiden"
)

type BooksController struct {
	raiden.ControllerBase
	Http  string `path:"/doctor" type:"rest"`
	Model models.Doctor
}
