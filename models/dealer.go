// dealer
package models

import (
	"github.com/vsdutka/ps/shared/database"
)

// *****************************************************************************
// Dealer
// *****************************************************************************

// Dealer table contains the information for each Dealer
type Dealer struct {
	ID        uint   `form:"p_id" gorm:"primary_key"`
	Name      string `form:"p_name" gorm:"not null;unique;unique_index" json:"name"`
	Phones    string `form:"p_phones" json:"phones"`
	Address   string `form:"p_address" json:"address"`
	Email     string `form:"p_email" json:"email"`
	SecretKey string `form:"p_secret_key" json:"-"`
	database.ModelBase
}
