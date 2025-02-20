package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type AuditLogEntries struct {
	db.ModelBase
	InstanceId *uuid.UUID  `json:"instance_id,omitempty" column:"name:instance_id;type:uuid;nullable"`
	Id         uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	Payload    interface{} `json:"payload,omitempty" column:"name:payload;type:json;nullable"`
	CreatedAt  *time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable"`
	IpAddress  string      `json:"ip_address,omitempty" column:"name:ip_address;type:varchar;nullable:false;default:'' varying"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"audit_log_entries" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
