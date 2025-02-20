package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type RefreshTokens struct {
	db.ModelBase
	InstanceId *uuid.UUID `json:"instance_id,omitempty" column:"name:instance_id;type:uuid;nullable"`
	Id         int64      `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;nullable:false;default:nextval('auth.refresh_tokens_id_seq')"`
	Token      *string    `json:"token,omitempty" column:"name:token;type:varchar;nullable;unique"`
	UserId     *string    `json:"user_id,omitempty" column:"name:user_id;type:varchar;nullable"`
	Revoked    *bool      `json:"revoked,omitempty" column:"name:revoked;type:boolean;nullable"`
	CreatedAt  *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`
	Parent     *string    `json:"parent,omitempty" column:"name:parent;type:varchar;nullable"`
	SessionId  *uuid.UUID `json:"session_id,omitempty" column:"name:session_id;type:uuid;nullable"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"refresh_tokens" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Session *Sessions `json:"session,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:session_id"`
}
