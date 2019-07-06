package config

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"github.com/jinzhu/gorm"
)

//DatabaseConfig はDatabaseに利用する各種設定を管理するMethodを定義します
type DatabaseConfig interface {
	OpenDB(connect string) *gorm.DB
	SetDBMS(dbms string)
	SetDBName(name string)
	SetUser(user string)
	SetPass(pass string)
	SetPROTOCOL(protocol string)
	GetDBMS() string
	GetConnect() string
}

//databaseConfig はDatabaseで利用する情報を定義します
type databaseConfig struct {
	DBMS     string
	USER     string
	PASS     string
	PROTOCOL string
	DBNAME   string
}

//NewDatabaseConfig はDataBaseのConfigを取得できます
func NewDatabaseConfig() DatabaseConfig {
	buf, err := ioutil.ReadFile("./_config/config_development.yml")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	confyaml := map[string]interface{}{}
	yaml.Unmarshal(buf, confyaml)
	if err != nil {
		fmt.Println(err)
	}
	dbconf := confyaml["DATABASE"].(map[interface{}]interface{})
	return &databaseConfig{
		DBMS:     dbconf["DBMS"].(string),
		USER:     dbconf["USER"].(string),
		PASS:     dbconf["PASS"].(string),
		PROTOCOL: dbconf["PROTOCOL"].(string),
		DBNAME:   dbconf["DBNAME"].(string),
	}
}

func (dbc *databaseConfig) SetDBMS(dbms string)         { dbc.DBMS = dbms }
func (dbc *databaseConfig) SetDBName(name string)       { dbc.DBNAME = name }
func (dbc *databaseConfig) SetUser(user string)         { dbc.USER = user }
func (dbc *databaseConfig) SetPass(pass string)         { dbc.PASS = pass }
func (dbc *databaseConfig) SetPROTOCOL(protocol string) { dbc.PROTOCOL = protocol }
func (dbc *databaseConfig) GetDBMS() string             { return dbc.DBMS }
func (dbc *databaseConfig) GetConnect() string {
	switch dbc.DBMS {
	case "sqlite3":
		return dbc.DBNAME
	}
	return ""
}
func (dbc *databaseConfig) OpenDB(connect string) *gorm.DB {
	db, err := gorm.Open(dbc.DBMS, connect)
	if err != nil {
		panic("DB Bad connection!")
	}
	return db
}
