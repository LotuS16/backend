package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type SamlProviders struct {
	db.ModelBase
	Id               uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	SsoProviderId    uuid.UUID   `json:"sso_provider_id,omitempty" column:"name:sso_provider_id;type:uuid;nullable:false"`
	EntityId         string      `json:"entity_id,omitempty" column:"name:entity_id;type:text;nullable:false;unique"`
	MetadataXml      string      `json:"metadata_xml,omitempty" column:"name:metadata_xml;type:text;nullable:false"`
	MetadataUrl      *string     `json:"metadata_url,omitempty" column:"name:metadata_url;type:text;nullable"`
	AttributeMapping interface{} `json:"attribute_mapping,omitempty" column:"name:attribute_mapping;type:json;nullable"`
	CreatedAt        *time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable"`
	UpdatedAt        *time.Time  `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`
	NameIdFormat     *string     `json:"name_id_format,omitempty" column:"name:name_id_format;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"saml_providers" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	SsoProviderSso *SsoProviders `json:"sso_provider_sso,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:sso_provider_id"`
}
