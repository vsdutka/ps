// user
package models

import (
	"fmt"
	"time"

	//	"encoding/json"

	//	"database/sql"

	//	"github.com/jinzhu/gorm"
	"github.com/ttacon/libphonenumber"
	"github.com/vsdutka/ps/shared/database"
	"github.com/vsdutka/ps/utils"
	"gopkg.in/errgo.v1"
)

// *****************************************************************************
// User
// *****************************************************************************

// User table contains the information for each user
type User struct {
	ID           uint      `gorm:"primary_key" form:"p_id"`
	FirstName    string    `form:"p_first_name"`
	LastName     string    `form:"p_last_name"`
	Email        string    `form:"p_email"`
	Phone        string    `form:"p_phone" gorm:"not null;unique;unique_index"`
	PinCode      string    `form:"p_pin_code" json:"-"`
	PinTimestamp time.Time `json:"-"`
	PinConfirmed bool      `json:"-"`
	SecretKey    string    `json:"-"`
	database.ModelBase
}

func (u *User) UserCreate() error {
	u.PinCode = utils.Pin()
	u.PinTimestamp = time.Now().Add(30 * time.Second)
	u.PinConfirmed = false
	u.SecretKey = utils.SecureRandomAlphaString(40)

	phone, err := phoneNormalize(u.Phone)
	if err != nil {
		return err
	}
	u.Phone = phone

	if err := database.DB.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) UserUpdate() error {
	if (u.ID == 0) && (u.Phone == "") {
		return errgo.New("отсутствуют данные для по идентификатору пользователя и номеру телефона")
	}
	var old User

	if u.ID != 0 {
		if err := database.DB.Where("id = ?", u.ID).First(&old).Error; err != nil {
			return err
		}
	} else {
		if u.Phone != "" {
			var err error
			u.Phone, err = phoneNormalize(u.Phone)
			if err != nil {
				return err
			}
			if err = database.DB.Where("phone = ?", u.Phone).First(&old).Error; err != nil {
				return err
			}
		}
	}
	if u.PinCode != "" {
		//		fmt.Println(old.PinTimestamp)
		if old.PinConfirmed {
			return errgo.New("Пользователь уже подтвержден")
		}
		if old.PinTimestamp.Before(time.Now()) {
			return errgo.New("Истекло время на подтверждение регистрации")
		}
		if u.PinCode != old.PinCode {
			return errgo.New("код подтверждения не подходит")
		}

		u.PinCode = ""
		u.PinConfirmed = true
		u.PinTimestamp = time.Time{}

		u.ID = old.ID
		u.Phone = old.Phone
		u.FirstName = old.FirstName
		u.LastName = old.LastName
		u.Email = old.Email
		u.SecretKey = old.SecretKey
	}

	if err := database.DB.Save(&u).Error; err != nil {
		return err
	}
	return nil
}
func (u *User) UserDel() error {
	if (u.ID == 0) && (u.Phone == "") {
		return errgo.New("отсутствуют данные для по идентификатору пользователя и номеру телефона")
	}

	if err := database.DB.Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) UserGet() error {
	if (u.ID == 0) && (u.Phone == "") {
		return errgo.New("отсутствуют данные для по идентификатору пользователя и номеру телефона")
	}

	if u.ID != 0 {
		if err := database.DB.Where("id = ?", u.ID).First(&u).Error; err != nil {
			return err
		}
	} else {
		if u.Phone != "" {
			var err error
			u.Phone, err = phoneNormalize(u.Phone)
			if err != nil {
				return err
			}
			if err = database.DB.Where("phone = ?", u.Phone).First(&u).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func UserList() ([]User, error) {
	var ul []User
	if err := database.DB.Find(&ul).Error; err != nil {
		return ul, err
	}
	return ul, nil
}

//func (u *User) UserID() string {
//	return fmt.Sprintf("%d", u.ID)
//}

//func UserTableCreate() error {

//	_, err := database.ExecDDL(`
//create table if not exists users (
//	id integer primary key autoincrement,
//	first_name varchar(240),
//	last_name varchar(240),
//	email varchar(40),
//	phone varchar(40) unique,
//	created_at date,
//	updated_at date,
//	pin_code varchar(10),
//	pin_timestamp date,
//	pin_confirmed boolean,
//	secret_key varchar(40)
//)`)
//	return err
//}

//func UserByID(id string) (User, error) {
//	o := User{}
//	if err := database.Get(&o, "select * from users where id=$1", id); err != nil {
//		return o, err
//	}
//	return o, nil
//}

//func UserByPhone(phone string) (User, error) {
//	o := User{}
//	var err error
//	phone, err = phoneNormalize(phone)
//	if err != nil {
//		return o, err
//	}
//	if err := database.Get(&o, "select * from users where phone=$1", phone); err != nil {
//		return o, err
//	}
//	return o, nil
//}

//// UserCreate creates user
//func UserUpdate(id, firstName, lastName, email, phone, pinCode string) (User, error) {
//	o := User{}

//	if phone != "" {
//		var err error
//		phone, err = phoneNormalize(phone)
//		if err != nil {
//			return o, err
//		}
//	}
//	found, err := func() (bool, error) {
//		var err error
//		if id != "" {
//			if o, err = UserByID(id); err != nil {
//				return false, err
//			}
//			return true, nil
//		}
//		if phone != "" {
//			if o, err = UserByPhone(phone); err != nil {
//				return false, err
//			}
//			return true, nil
//		}
//		return false, nil
//	}()

//	if found {
//		o.FirstName = firstName
//		o.LastName = lastName
//		o.Email = email
//		o.Phone = phone
//		if pinCode != "" {
//			if o.PinConfirmed {
//				return o, errgo.New("Пользователь уже подтвержден")
//			}
//			if o.PinTimestamp.Before(time.Now()) {
//				return o, errgo.New("Истекло время на подтверждение регистрации")
//			}
//			if o.PinCode != pinCode {
//				return o, errgo.New("код подтверждения не подходит")
//			}

//			o.PinCode = ""
//			o.PinConfirmed = true
//			o.PinTimestamp = time.Time{}
//		}
//		o.UpdatedAt = time.Now()

//		_, err := database.Exec("update users set first_name=:first_name, last_name=:last_name, email=:email, phone=:phone, created_at=:created_at, updated_at=:updated_at, pin_code=:pin_code, pin_timestamp=:pin_timestamp, pin_confirmed=:pin_confirmed, secret_key=:secret_key where id = :id", o)
//		if err != nil {
//			return o, err
//		}
//		return o, nil
//	}
//	now := time.Now()
//	o.FirstName = firstName
//	o.LastName = lastName
//	o.Email = email
//	o.Phone = phone
//	o.CreatedAt = now
//	o.UpdatedAt = now
//	o.PinCode = utils.Pin()
//	o.PinTimestamp = now.Add(30 * time.Minute)
//	o.PinConfirmed = false
//	o.SecretKey = utils.SecureRandomAlphaString(40)

//	res, err := database.Exec("insert into users(first_name, last_name, email, phone, created_at, updated_at, pin_code, pin_timestamp, pin_confirmed, secret_key) values(:first_name, :last_name, :email, :phone, :created_at, :updated_at, :pin_code, :pin_timestamp, :pin_confirmed, :secret_key)", o)
//	if err != nil {
//		return o, err
//	}
//	o.ID, err = res.LastInsertId()
//	if err != nil {
//		return o, err
//	}
//	if err = utils.SmsSend(o.Phone, fmt.Sprintf("Pin code : %s", o.PinCode)); err != nil {
//		return o, err
//	}
//	return o, nil
//}

//func Users() ([]User, error) {
//	users := []User{}
//	if err := database.Select(&users, "select * from users order by 1"); err != nil {
//		return users, err
//	}
//	return users, nil
//}

func phoneNormalize(phone string) (string, error) {
	num, err := libphonenumber.Parse(phone, "RU")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("+%d%d", *num.CountryCode, *num.NationalNumber), nil
}
