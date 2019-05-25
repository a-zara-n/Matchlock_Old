package datastore

import (
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	Database string
	Table    interface{}
}

func (d *Database) openDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", d.Database)
	d.checkConnectError(err)
	return db
}

func (d *Database) checkConnectError(err error) {
	if err != nil {
		panic("DB Bad connection!")
	}
}

func (d *Database) InitMigration() {
	db := d.openDatabase()
	defer db.Close()
	db.AutoMigrate(d.Table)
}

func (d *Database) Insert(data interface{}) bool {
	db := d.openDatabase()
	defer db.Close()
	db.Create(data)
	return true
}
func (d *Database) Update(data interface{}) bool {
	if reflect.TypeOf(d.Table) != reflect.TypeOf(data) {
		return false
	}
	db := d.openDatabase()
	defer db.Close()
	db.Create(data)
	return true
}
