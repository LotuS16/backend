package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Users struct {
	db.ModelBase
	InstanceId               *uuid.UUID  `json:"instance_id,omitempty" column:"name:instance_id;type:uuid;nullable"`
	Id                       uuid.UUID   `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false"`
	Aud                      *string     `json:"aud,omitempty" column:"name:aud;type:varchar;nullable"`
	Role                     *string     `json:"role,omitempty" column:"name:role;type:varchar;nullable"`
	Email                    *string     `json:"email,omitempty" column:"name:email;type:varchar;nullable"`
	EncryptedPassword        *string     `json:"encrypted_password,omitempty" column:"name:encrypted_password;type:varchar;nullable"`
	EmailConfirmedAt         *time.Time  `json:"email_confirmed_at,omitempty" column:"name:email_confirmed_at;type:timestampz;nullable"`
	InvitedAt                *time.Time  `json:"invited_at,omitempty" column:"name:invited_at;type:timestampz;nullable"`
	ConfirmationToken        *string     `json:"confirmation_token,omitempty" column:"name:confirmation_token;type:varchar;nullable"`
	ConfirmationSentAt       *time.Time  `json:"confirmation_sent_at,omitempty" column:"name:confirmation_sent_at;type:timestampz;nullable"`
	RecoveryToken            *string     `json:"recovery_token,omitempty" column:"name:recovery_token;type:varchar;nullable"`
	RecoverySentAt           *time.Time  `json:"recovery_sent_at,omitempty" column:"name:recovery_sent_at;type:timestampz;nullable"`
	EmailChangeTokenNew      *string     `json:"email_change_token_new,omitempty" column:"name:email_change_token_new;type:varchar;nullable"`
	EmailChange              *string     `json:"email_change,omitempty" column:"name:email_change;type:varchar;nullable"`
	EmailChangeSentAt        *time.Time  `json:"email_change_sent_at,omitempty" column:"name:email_change_sent_at;type:timestampz;nullable"`
	LastSignInAt             *time.Time  `json:"last_sign_in_at,omitempty" column:"name:last_sign_in_at;type:timestampz;nullable"`
	RawAppMetaData           interface{} `json:"raw_app_meta_data,omitempty" column:"name:raw_app_meta_data;type:json;nullable"`
	RawUserMetaData          interface{} `json:"raw_user_meta_data,omitempty" column:"name:raw_user_meta_data;type:json;nullable"`
	IsSuperAdmin             *bool       `json:"is_super_admin,omitempty" column:"name:is_super_admin;type:boolean;nullable"`
	CreatedAt                *time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable"`
	UpdatedAt                *time.Time  `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable"`
	Phone                    *string     `json:"phone,omitempty" column:"name:phone;type:text;nullable;default:NULL varying;unique"`
	PhoneConfirmedAt         *time.Time  `json:"phone_confirmed_at,omitempty" column:"name:phone_confirmed_at;type:timestampz;nullable"`
	PhoneChange              *string     `json:"phone_change,omitempty" column:"name:phone_change;type:text;nullable;default:'' varying"`
	PhoneChangeToken         *string     `json:"phone_change_token,omitempty" column:"name:phone_change_token;type:varchar;nullable;default:'' varying"`
	PhoneChangeSentAt        *time.Time  `json:"phone_change_sent_at,omitempty" column:"name:phone_change_sent_at;type:timestampz;nullable"`
	ConfirmedAt              *time.Time  `json:"confirmed_at,omitempty" column:"name:confirmed_at;type:timestampz;nullable;default:LEAST(email_confirmed_at, phone_confirmed_at)"`
	EmailChangeTokenCurrent  *string     `json:"email_change_token_current,omitempty" column:"name:email_change_token_current;type:varchar;nullable;default:'' varying"`
	EmailChangeConfirmStatus *int16      `json:"email_change_confirm_status,omitempty" column:"name:email_change_confirm_status;type:smallint;nullable;default:0"`
	BannedUntil              *time.Time  `json:"banned_until,omitempty" column:"name:banned_until;type:timestampz;nullable"`
	ReauthenticationToken    *string     `json:"reauthentication_token,omitempty" column:"name:reauthentication_token;type:varchar;nullable;default:'' varying"`
	ReauthenticationSentAt   *time.Time  `json:"reauthentication_sent_at,omitempty" column:"name:reauthentication_sent_at;type:timestampz;nullable"`
	IsSsoUser                bool        `json:"is_sso_user,omitempty" column:"name:is_sso_user;type:boolean;nullable:false;default:false"`
	DeletedAt                *time.Time  `json:"deleted_at,omitempty" column:"name:deleted_at;type:timestampz;nullable"`
	IsAnonymous              bool        `json:"is_anonymous,omitempty" column:"name:is_anonymous;type:boolean;nullable:false;default:false"`

	// Table information
	Metadata string `json:"-" schema:"auth" tableName:"users" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	BookingControllerUserids                   []*BookingController `json:"booking_controller_userids,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:userid"`
	DoctorUsers                                []*Doctor            `json:"doctor_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	FaskesThroughBookingControllerUserid       []*Faskes            `json:"faskes_through_booking_controller_userid,omitempty" join:"joinType:manyToMany;through:booking_controller;sourcePrimaryKey:id;sourceForeignKey:userid;targetPrimaryKey:id;targetForeign:userid"`
	IdentityUsers                              []*Identities        `json:"identity_users,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	JadwalDoctorUsers                          []*JadwalDoctors     `json:"jadwal_doctor_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	MfaFactorUsers                             []*MfaFactors        `json:"mfa_factor_users,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	OneTimeTokenUsers                          []*OneTimeTokens     `json:"one_time_token_users,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	ReservationsThroughBookingControllerUserid []*Reservation       `json:"reservations_through_booking_controller_userid,omitempty" join:"joinType:manyToMany;through:booking_controller;sourcePrimaryKey:id;sourceForeignKey:userid;targetPrimaryKey:reservationid;targetForeign:userid"`
	SessionUsers                               []*Sessions          `json:"session_users,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	UserRoleUsers                              []*UserRoles         `json:"user_role_users,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
}
