package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type SchemaMigrations struct {
	db.ModelBase
	Version string `json:"version,omitempty" column:"name:version;type:varchar;primaryKey;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"schema_migrations" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
