package datastore

import (
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB = Database{Database: "./test.db"}

type Database struct {
	Database string
	Table    interface{}
}

func (d *Database) OpenDatabase() *gorm.DB {
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
	db := d.OpenDatabase()
	defer db.Close()
	db.AutoMigrate(d.Table)
}

func (d *Database) Insert(data interface{}) bool {
	db := d.OpenDatabase()
	defer db.Close()
	db.Create(data)
	return true
}
func (d *Database) Update(data interface{}) bool {
	if reflect.TypeOf(d.Table) != reflect.TypeOf(data) {
		return false
	}
	db := d.OpenDatabase()
	defer db.Close()
	db.Create(data)
	return true
}

func (d *Database) SelectALL(retSchema interface{}) interface{} {
	db := d.OpenDatabase()
	defer db.Close()
	db.Find(retSchema)
	return retSchema
}

func (d *Database) SelectWhere(where map[string]interface{}, retSchema interface{}) interface{} {
	db := d.OpenDatabase()
	defer db.Close()
	db.Where(where).Find(retSchema)
	return retSchema
}
