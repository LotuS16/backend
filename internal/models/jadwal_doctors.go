package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type JadwalDoctors struct {
	db.ModelBase
	Id          int64      `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	Haripraktek string     `json:"haripraktek,omitempty" column:"name:haripraktek;type:varchar;nullable:false"`
	Jampraktek  *string    `json:"jampraktek,omitempty" column:"name:jampraktek;type:varchar;nullable"`
	UserId      *uuid.UUID `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable;default:auth.uid()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"jadwal_doctors" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	User *Users `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
