package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type SsoProviders struct {
	db.ModelBase
	Id         uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	ResourceId *string    `json:"resource_id,omitempty" column:"name:resource_id;type:text;nullable"`
	CreatedAt  *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"sso_providers" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	FlowStatesThroughSamlRelayStatesSsoProvider []*FlowState       `json:"flow_states_through_saml_relay_states_sso_provider,omitempty" join:"joinType:manyToMany;through:saml_relay_states;sourcePrimaryKey:id;sourceForeignKey:sso_provider_id;targetPrimaryKey:id;targetForeign:sso_provider_id"`
	SamlProviderSsos                            []*SamlProviders   `json:"saml_provider_ssos,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:sso_provider_id"`
	SamlRelayStateSsos                          []*SamlRelayStates `json:"saml_relay_state_ssos,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:sso_provider_id"`
	SsoDomainSsos                               []*SsoDomains      `json:"sso_domain_ssos,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:sso_provider_id"`
}
