package types

import (
	"github.com/sev-2/raiden"
)

type BookingController struct {
	raiden.TypeBase
}

func (t *BookingController) Name() string {
	return "_booking_controller"
}

func (r *BookingController) Format() string {
	return "booking_controller[]"
}

func (r *BookingController) Enums() []string {
	return []string{}
}

func (r *BookingController) Attributes() []string {
	return []string{}
}

func (r *BookingController) Comment() *string {
	return nil
}

