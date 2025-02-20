package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Sessions struct {
	db.ModelBase
	Id          uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	UserId      uuid.UUID   `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false"`
	CreatedAt   *time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable"`
	UpdatedAt   *time.Time  `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`
	FactorId    *uuid.UUID  `json:"factor_id,omitempty" column:"name:factor_id;type:uuid;nullable"`
	Aal         interface{} `json:"aal,omitempty" column:"name:aal;nullable"`
	NotAfter    *time.Time  `json:"not_after,omitempty" column:"name:not_after;type:timestampz;nullable"`
	RefreshedAt *time.Time  `json:"refreshed_at,omitempty" column:"name:refreshed_at;type:timestamp;nullable"`
	UserAgent   *string     `json:"user_agent,omitempty" column:"name:user_agent;type:text;nullable"`
	Ip          interface{} `json:"ip,omitempty" column:"name:ip;nullable"`
	Tag         *string     `json:"tag,omitempty" column:"name:tag;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"sessions" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	MfaAmrClaimSessions  []*MfaAmrClaims  `json:"mfa_amr_claim_sessions,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:session_id"`
	RefreshTokenSessions []*RefreshTokens `json:"refresh_token_sessions,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:session_id"`
	User                 *Users           `json:"user,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
