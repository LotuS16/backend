package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Poli struct {
	db.ModelBase
	Id             uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Namapoli       *string   `json:"namapoli,omitempty" column:"name:namapoli;type:varchar;nullable"`
	Keteranganpoli *string   `json:"keteranganpoli,omitempty" column:"name:keteranganpoli;type:varchar;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"poli" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
