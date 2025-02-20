package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type OneTimeTokens struct {
	db.ModelBase
	Id        uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	UserId    uuid.UUID   `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false"`
	TokenType interface{} `json:"token_type,omitempty" column:"name:token_type;nullable:false"`
	TokenHash string      `json:"token_hash,omitempty" column:"name:token_hash;type:text;nullable:false"`
	RelatesTo string      `json:"relates_to,omitempty" column:"name:relates_to;type:text;nullable:false"`
	CreatedAt time.Time   `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable:false;default:now()"`
	UpdatedAt time.Time   `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp;nullable:false;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"one_time_tokens" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	User *Users `json:"user,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
