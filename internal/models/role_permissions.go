package models

import (
	"backend/internal/types"
	"github.com/sev-2/raiden/pkg/db"
)

type RolePermissions struct {
	db.ModelBase
	Id         int64               `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	Role       types.AppRole       `json:"role,omitempty" column:"name:role;type:app_role;nullable:false"`
	Permission types.AppPermission `json:"permission,omitempty" column:"name:permission;type:app_permission;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"role_permissions" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
