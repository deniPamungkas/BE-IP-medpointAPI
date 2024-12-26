package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type PublicUsers struct {
	db.ModelBase
	Id        uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:auth.uid()"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	Email     *string   `json:"email,omitempty" column:"name:email;type:text;nullable"`
	Username  *string   `json:"username,omitempty" column:"name:username;type:text;nullable"`
	Password  *string   `json:"password,omitempty" column:"name:password;type:text;nullable"`
	Role      *string   `json:"role,omitempty" column:"name:role;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"public_users" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorUsers                       []*Doctors        `json:"doctor_users,omitempty" onUpdate:"cascade" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	SpecializationsThroughDoctorsUser []*Specialization `json:"specializations_through_doctors_user,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
}
