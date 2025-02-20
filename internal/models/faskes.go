package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Faskes struct {
	db.ModelBase
	Id               uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Namafaskes       *string   `json:"namafaskes,omitempty" column:"name:namafaskes;type:varchar;nullable"`
	Keteranganfaskes *string   `json:"keteranganfaskes,omitempty" column:"name:keteranganfaskes;type:varchar;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"faskes" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	BookingControllerFaskesids                   []*BookingController `json:"booking_controller_faskesids,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:faskesid"`
	ReservationsThroughBookingControllerFaskesid []*Reservation       `json:"reservations_through_booking_controller_faskesid,omitempty" join:"joinType:manyToMany;through:booking_controller;sourcePrimaryKey:id;sourceForeignKey:faskesid;targetPrimaryKey:reservationid;targetForeign:faskesid"`
	UsersThroughBookingControllerFaskesid        []*Users             `json:"users_through_booking_controller_faskesid,omitempty" join:"joinType:manyToMany;through:booking_controller;sourcePrimaryKey:id;sourceForeignKey:faskesid;targetPrimaryKey:id;targetForeign:faskesid"`
}
