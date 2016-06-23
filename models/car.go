// car
package models

import (
	"github.com/vsdutka/ps/shared/database"
)

// *****************************************************************************
// Car
// *****************************************************************************

// Car table contains the information for each Car
type Car struct {
	ID      uint   `form:"p_id" gorm:"primary_key"`
	UserID  uint   `form:"p_user_id" gorm:"not null" json:"user_id" sql:"type:integer REFERENCES users(id)"`
	BrandID uint   `form:"p_brand_id" gorm:"not null" json:"brand_id" sql:"type:integer REFERENCES brands(id)"`
	ModelID uint   `form:"p_model_id" gorm:"not null" json:"model_id" sql:"type:integer REFERENCES models(id)"`
	Vin     string `form:"p_vin" gorm:"not null;unique;unique_index" json:"vin"`
	Year    uint   `form:"p_year" gorm:"not null" json:"year"`
	GosNum  string `form:"p_gos_num" json:"gos_num"`
	User    User   `gorm:"ForeignKey:UserID"`
	Brand   Brand  `gorm:"ForeignKey:BrandID"`
	Model   Model  `gorm:"ForeignKey:ModelID"`
	database.ModelBase
}

func (o *Car) LoadAssosiation() error {
	if err := database.Sel(&o.Brand, o.BrandID); err != nil {
		return err
	}
	if err := database.Sel(&o.Model, o.ModelID); err != nil {
		return err
	}
	if err := database.Sel(&o.User, o.UserID); err != nil {
		return err
	}
	return nil
}
