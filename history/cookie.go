package history

import (
	"github.com/a-zara-n/Matchlock/datastore"
	"github.com/jinzhu/gorm"
)

//Cookie represents the schema when saving to DB.
type Cookie struct {
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

//SetCookie はDBにcookieを挿入します
func (c Cookie) SetCookie(name string, value string) {
	c.Name = name
	c.Value = value
	datastore.DB.Insert(c)
}

//GetCookie はcookieを取得できます
func (c Cookie) GetCookie(domain string, path string) []Cookie {
	cdb := datastore.DB.OpenDatabase()
	var cookies []Cookie
	cdb.Select("*").
		Where("domain = ? AND path LIKE ? ", domain, path+"%").
		Find(&cookies)
	return cookies
}
