package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type SamlRelayStates struct {
	db.ModelBase
	Id            uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	SsoProviderId uuid.UUID  `json:"sso_provider_id,omitempty" column:"name:sso_provider_id;type:uuid;nullable:false"`
	RequestId     string     `json:"request_id,omitempty" column:"name:request_id;type:text;nullable:false"`
	ForEmail      *string    `json:"for_email,omitempty" column:"name:for_email;type:text;nullable"`
	RedirectTo    *string    `json:"redirect_to,omitempty" column:"name:redirect_to;type:text;nullable"`
	CreatedAt     *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`
	FlowStateId   *uuid.UUID `json:"flow_state_id,omitempty" column:"name:flow_state_id;type:uuid;nullable"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"saml_relay_states" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	FlowStateFlow  *FlowState    `json:"flow_state_flow,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:flow_state_id"`
	SsoProviderSso *SsoProviders `json:"sso_provider_sso,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:sso_provider_id"`
}
