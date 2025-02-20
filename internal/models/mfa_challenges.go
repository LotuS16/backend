package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type MfaChallenges struct {
	db.ModelBase
	Id                  uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	FactorId            uuid.UUID   `json:"factor_id,omitempty" column:"name:factor_id;type:uuid;nullable:false"`
	CreatedAt           time.Time   `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false"`
	VerifiedAt          *time.Time  `json:"verified_at,omitempty" column:"name:verified_at;type:timestampz;nullable"`
	IpAddress           interface{} `json:"ip_address,omitempty" column:"name:ip_address;nullable:false"`
	OtpCode             *string     `json:"otp_code,omitempty" column:"name:otp_code;type:text;nullable"`
	WebAuthnSessionData interface{} `json:"web_authn_session_data,omitempty" column:"name:web_authn_session_data;type:json;nullable"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"mfa_challenges" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	MfaFactorFactor *MfaFactors `json:"mfa_factor_factor,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:factor_id"`
}
