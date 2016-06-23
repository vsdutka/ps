// brand
package models

import (
	"github.com/vsdutka/ps/shared/database"
)

// *****************************************************************************
// Brand
// *****************************************************************************

// Brand table contains the information for each Brand
type Brand struct {
	ID   uint   `form:"p_id" gorm:"primary_key"`
	Name string `form:"p_name" gorm:"not null;unique;unique_index" json:"name"`
	database.ModelBase
}
