package datastore

/*
datastore
このパッケージはツールでの簡単なデータベース関連の操作をパックするためのものです
*/
import (
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Database は各種データベース系の情報を管理する構造体です。
type Database struct {
	Database string
	Table    interface{}
}

// DB はDatabaseの保存先を仮指定する項目です。
var DB = Database{Database: "./test.db"}

// ChangeDatabase はSQLiteのデータベースの変更をするための関数です。
func (d *Database) ChangeDatabase(changeDBName string) {
	d.Database = changeDBName
}

// OpenDatabase はDatabaseに定義されているDBとのコネクトを開き、gorm.DBを返却します。
func (d *Database) OpenDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", d.Database)
	if err != nil {
		panic("DB Bad connection!")
	}
	return db
}

// InitMigration はDatabase.Tableに設定されているテーブルをinitしてくれます
func (d *Database) InitMigration() {
	db := d.OpenDatabase()
	defer db.Close()
	db.AutoMigrate(d.Table)
}

//Insert はDatabase.Tableに設定されている同じ型をdataとして第一引数に渡してDataを挿入します
func (d *Database) Insert(data interface{}) bool {
	db := d.OpenDatabase()
	defer db.Close()
	db.Create(data)
	return true
}

//Update はDatabase.Tableに設定されている同じ型で、変更したい内容をdataとして第一引数に渡して更新します。
func (d *Database) Update(data interface{}) bool {
	if reflect.TypeOf(d.Table) != reflect.TypeOf(data) {
		return false
	}
	db := d.OpenDatabase()
	defer db.Close()
	db.Create(data)
	return true
}

// SelectALL はDatabase.Tableに設定されている同じ型のデータを全て取り出します
func (d *Database) SelectALL(retSchema interface{}) interface{} {
	db := d.OpenDatabase()
	defer db.Close()
	db.Find(retSchema)
	return retSchema
}

//SelectWhere は第一引数にmapで書かれたwhereを条件式に割り当て検索することが可能です
func (d *Database) SelectWhere(where map[string]interface{}, retSchema interface{}) interface{} {
	db := d.OpenDatabase()
	defer db.Close()
	db.Where(where).Find(retSchema)
	return retSchema
}
