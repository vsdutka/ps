// ws
package models

import (
	"github.com/vsdutka/ps/shared/database"
)

// *****************************************************************************
// WorkType
// *****************************************************************************

// WorkType table contains the information for each WorkType
type WorkType struct {
	ID   uint   `form:"p_id" gorm:"primary_key"`
	Name string `form:"p_name" gorm:"not null;unique;unique_index" json:"name"`
	database.ModelBase
}
