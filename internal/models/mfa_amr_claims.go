package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type MfaAmrClaims struct {
	db.ModelBase
	SessionId            uuid.UUID `json:"session_id,omitempty" column:"name:session_id;type:uuid;nullable:false"`
	CreatedAt            time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false"`
	UpdatedAt            time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable:false"`
	AuthenticationMethod string    `json:"authentication_method,omitempty" column:"name:authentication_method;type:text;nullable:false"`
	Id                   uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"mfa_amr_claims" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Session *Sessions `json:"session,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:session_id"`
}
