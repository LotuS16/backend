package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Identities struct {
	db.ModelBase
	ProviderId   string      `json:"provider_id,omitempty" column:"name:provider_id;type:text;nullable:false"`
	UserId       uuid.UUID   `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false"`
	IdentityData interface{} `json:"identity_data,omitempty" column:"name:identity_data;type:json;nullable:false"`
	Provider     string      `json:"provider,omitempty" column:"name:provider;type:text;nullable:false"`
	LastSignInAt *time.Time  `json:"last_sign_in_at,omitempty" column:"name:last_sign_in_at;type:timestampz;nullable"`
	CreatedAt    *time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable"`
	UpdatedAt    *time.Time  `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`
	Email        *string     `json:"email,omitempty" column:"name:email;type:text;nullable;default:lower((identity_data ->> 'email'))"`
	Id           uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"identities" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	User *Users `json:"user,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
