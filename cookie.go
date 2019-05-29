package main

import "github.com/jinzhu/gorm"

//Cookie represents the schema when saving to DB.
type cookie struct {
	gorm.Model
	Domain   string
	Path     string
	Name     string
	Value    string
	Expires  string
	Httponry bool
	Secure   bool
	SameSite string
}

func (c cookie) SetCookie(name string, value string) {
	c.Name = name
	c.Value = value
	db.Table = RequestData{}
	db.Insert(c)
}

func (c cookie) GetCookie(domain string, path string) []cookie {
	db.Table = cookie{}
	cdb := db.OpenDatabase()
	var cookies []cookie
	cdb.Select("*").
		Where("domain = ? AND path LIKE ? ", domain, path+"%").
		Find(&cookies)
	return cookies
}
