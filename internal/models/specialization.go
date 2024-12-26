package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Specialization struct {
	db.ModelBase
	Id             uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Specialization string    `json:"specialization,omitempty" column:"name:specialization;type:text;nullable:false"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"specialization" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorSpecializations                   []*Doctors     `json:"doctor_specializations,omitempty" onUpdate:"cascade" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:specialization_id"`
	PublicUsersThroughDoctorsSpecialization []*PublicUsers `json:"public_users_through_doctors_specialization,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:specialization_id;targetPrimaryKey:id;targetForeign:specialization_id"`
}
