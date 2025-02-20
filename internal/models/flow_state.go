package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type FlowState struct {
	db.ModelBase
	Id                   uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	UserId               *uuid.UUID  `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable"`
	AuthCode             string      `json:"auth_code,omitempty" column:"name:auth_code;type:text;nullable:false"`
	CodeChallengeMethod  interface{} `json:"code_challenge_method,omitempty" column:"name:code_challenge_method;nullable:false"`
	CodeChallenge        string      `json:"code_challenge,omitempty" column:"name:code_challenge;type:text;nullable:false"`
	ProviderType         string      `json:"provider_type,omitempty" column:"name:provider_type;type:text;nullable:false"`
	ProviderAccessToken  *string     `json:"provider_access_token,omitempty" column:"name:provider_access_token;type:text;nullable"`
	ProviderRefreshToken *string     `json:"provider_refresh_token,omitempty" column:"name:provider_refresh_token;type:text;nullable"`
	CreatedAt            *time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable"`
	UpdatedAt            *time.Time  `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`
	AuthenticationMethod string      `json:"authentication_method,omitempty" column:"name:authentication_method;type:text;nullable:false"`
	AuthCodeIssuedAt     *time.Time  `json:"auth_code_issued_at,omitempty" column:"name:auth_code_issued_at;type:timestampz;nullable"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"flow_state" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	SamlRelayStateFlows                         []*SamlRelayStates `json:"saml_relay_state_flows,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:flow_state_id"`
	SsoProvidersThroughSamlRelayStatesFlowState []*SsoProviders    `json:"sso_providers_through_saml_relay_states_flow_state,omitempty" join:"joinType:manyToMany;through:saml_relay_states;sourcePrimaryKey:id;sourceForeignKey:flow_state_id;targetPrimaryKey:id;targetForeign:flow_state_id"`
}
