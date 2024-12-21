package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Doctors struct {
	db.ModelBase
	Id               uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name             string     `json:"name,omitempty" column:"name:name;type:text;nullable:false"`
	Gender           string     `json:"gender,omitempty" column:"name:gender;type:text;nullable:false"`
	SpecializationId *uuid.UUID `json:"specialization_id,omitempty" column:"name:specialization_id;type:uuid;nullable;default:gen_random_uuid()"`
	UserId           *uuid.UUID `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable;default:gen_random_uuid()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctors" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
