// database
package database

import (
	"log"
	"time"

	_ "github.com/gin-gonic/gin/binding"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ModelBase struct {
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

var (
	DB *gorm.DB
)

// Connect to the database
func Open(path string) {
	var err error

	if DB, err = gorm.Open("sqlite3", path); err != nil {
		log.Println("Database connect error", err)
	}
	DB.Exec("PRAGMA foreign_keys = ON")
}
func Close() {
	_ = DB.Close()
}

func Ins(value interface{}) error {
	return DB.Create(value).Error
}

func Upd(value interface{}) error {
	return DB.Save(value).Error
}

func Del(value interface{}) error {
	return DB.Delete(value).Error
}

func Sel(value interface{}, id uint) error {
	return DB.First(value, id).Error
}

func List(value interface{}, where ...interface{}) error {
	return DB.Find(value, where...).Error
}
