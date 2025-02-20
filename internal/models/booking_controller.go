package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type BookingController struct {
	db.ModelBase
	Boid          uuid.UUID  `json:"boid,omitempty" column:"name:boid;type:uuid;primaryKey;nullable:false"`
	Userid        *uuid.UUID `json:"userid,omitempty" column:"name:userid;type:uuid;nullable"`
	Doctoid       *uuid.UUID `json:"doctoid,omitempty" column:"name:doctoid;type:uuid;nullable"`
	Faskesid      *uuid.UUID `json:"faskesid,omitempty" column:"name:faskesid;type:uuid;nullable"`
	Reservationid *uuid.UUID `json:"reservationid,omitempty" column:"name:reservationid;type:uuid;nullable"`
	Bookingdate   *time.Time `json:"bookingdate,omitempty" column:"name:bookingdate;type:date;nullable"`
	Bookingtime   *time.Time `json:"bookingtime,omitempty" column:"name:bookingtime;type:time;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"booking_controller" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	FaskeFaskesid            *Faskes      `json:"faske_faskesid,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:faskesid"`
	ReservationReservationid *Reservation `json:"reservation_reservationid,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:reservationid;foreignKey:reservationid"`
	UserUserid               *Users       `json:"user_userid,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:userid"`
}
