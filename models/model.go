// model
package models

import (
	"github.com/vsdutka/ps/shared/database"
)

// *****************************************************************************
// Model
// *****************************************************************************

// Model table contains the information for each Model
type Model struct {
	ID      uint   `form:"p_id" gorm:"primary_key"`
	Name    string `form:"p_name" gorm:"not null;unique_index:uk_model" json:"name"`
	BrandID uint   `form:"p_brand_id" gorm:"not null;unique_index:uk_model" json:"brand_id" sql:"type:integer REFERENCES brands(id)"`
	Brand   Brand  `gorm:"ForeignKey:BrandID;;AssociationForeignKey:Name"`
	database.ModelBase
}
