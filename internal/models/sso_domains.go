package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type SsoDomains struct {
	db.ModelBase
	Id            uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	SsoProviderId uuid.UUID  `json:"sso_provider_id,omitempty" column:"name:sso_provider_id;type:uuid;nullable:false"`
	Domain        string     `json:"domain,omitempty" column:"name:domain;type:text;nullable:false"`
	CreatedAt     *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"sso_domains" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	SsoProviderSso *SsoProviders `json:"sso_provider_sso,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:sso_provider_id"`
}
