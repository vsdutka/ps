// persnon
package models

import (
	"github.com/vsdutka/ps/shared/database"
)

// *****************************************************************************
// Person
// *****************************************************************************

// Person table contains the information for each Person
type Person struct {
	ID        uint   `form:"p_id" gorm:"primary_key"`
	FirstName string `form:"p_first_name" gorm:"not null" json:"first_name"`
	LastName  string `form:"p_last_name" gorm:"not null" json:"last_name"`
	DealerID  uint   `form:"p_dealer_id" gorm:"not null" json:"dealer_id" sql:"type:integer REFERENCES dealers(id)"`
	Photo     []byte `form:"p_photo" json:"photo"`
	Dealer    Dealer `gorm:"ForeignKey:DealerID"`
	database.ModelBase
}
