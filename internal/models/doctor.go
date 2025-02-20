package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Doctor struct {
	db.ModelBase
	Id         int64      `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	Namadoctor string     `json:"namadoctor,omitempty" column:"name:namadoctor;type:varchar;nullable:false"`
	Spesialis  *string    `json:"spesialis,omitempty" column:"name:spesialis;type:varchar;nullable"`
	UserId     *uuid.UUID `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable;default:auth.uid()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctor" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	User *Users `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
