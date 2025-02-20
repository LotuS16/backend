package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type MfaFactors struct {
	db.ModelBase
	Id                 uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	UserId             uuid.UUID   `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false"`
	FriendlyName       *string     `json:"friendly_name,omitempty" column:"name:friendly_name;type:text;nullable"`
	FactorType         interface{} `json:"factor_type,omitempty" column:"name:factor_type;nullable:false"`
	Status             interface{} `json:"status,omitempty" column:"name:status;nullable:false"`
	CreatedAt          time.Time   `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false"`
	UpdatedAt          time.Time   `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable:false"`
	Secret             *string     `json:"secret,omitempty" column:"name:secret;type:text;nullable"`
	Phone              *string     `json:"phone,omitempty" column:"name:phone;type:text;nullable"`
	LastChallengedAt   *time.Time  `json:"last_challenged_at,omitempty" column:"name:last_challenged_at;type:timestampz;nullable;unique"`
	WebAuthnCredential interface{} `json:"web_authn_credential,omitempty" column:"name:web_authn_credential;type:json;nullable"`
	WebAuthnAaguid     *uuid.UUID  `json:"web_authn_aaguid,omitempty" column:"name:web_authn_aaguid;type:uuid;nullable"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"mfa_factors" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	MfaChallengeFactors []*MfaChallenges `json:"mfa_challenge_factors,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:factor_id"`
	User                *Users           `json:"user,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
