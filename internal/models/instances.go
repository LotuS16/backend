package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Instances struct {
	db.ModelBase
	Id            uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	Uuid          *uuid.UUID `json:"uuid,omitempty" column:"name:uuid;type:uuid;nullable"`
	RawBaseConfig *string    `json:"raw_base_config,omitempty" column:"name:raw_base_config;type:text;nullable"`
	CreatedAt     *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"instances" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
