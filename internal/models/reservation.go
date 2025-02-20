package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Reservation struct {
	db.ModelBase
	Reservationid      uuid.UUID  `json:"reservationid,omitempty" column:"name:reservationid;type:uuid;primaryKey;nullable:false"`
	Userid             *uuid.UUID `json:"userid,omitempty" column:"name:userid;type:uuid;nullable"`
	Dokterid           *uuid.UUID `json:"dokterid,omitempty" column:"name:dokterid;type:uuid;nullable"`
	Faskesid           *uuid.UUID `json:"faskesid,omitempty" column:"name:faskesid;type:uuid;nullable"`
	Poliid             *uuid.UUID `json:"poliid,omitempty" column:"name:poliid;type:uuid;nullable"`
	Jadwaldoctorid     *uuid.UUID `json:"jadwaldoctorid,omitempty" column:"name:jadwaldoctorid;type:uuid;nullable"`
	Reservationday     *time.Time `json:"reservationday,omitempty" column:"name:reservationday;type:timestamp;nullable"`
	Reservationtime    *time.Time `json:"reservationtime,omitempty" column:"name:reservationtime;type:time;nullable"`
	Billingreservation *uuid.UUID `json:"billingreservation,omitempty" column:"name:billingreservation;type:uuid;nullable"`
	Billingupdate      *string    `json:"billingupdate,omitempty" column:"name:billingupdate;type:varchar;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"reservation" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	BookingControllerReservationids             []*BookingController `json:"booking_controller_reservationids,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:reservationid;foreignKey:reservationid"`
	FaskesThroughBookingControllerReservationid []*Faskes            `json:"faskes_through_booking_controller_reservationid,omitempty" join:"joinType:manyToMany;through:booking_controller;sourcePrimaryKey:reservationid;sourceForeignKey:reservationid;targetPrimaryKey:id;targetForeign:reservationid"`
	UsersThroughBookingControllerReservationid  []*Users             `json:"users_through_booking_controller_reservationid,omitempty" join:"joinType:manyToMany;through:booking_controller;sourcePrimaryKey:reservationid;sourceForeignKey:reservationid;targetPrimaryKey:id;targetForeign:reservationid"`
}
