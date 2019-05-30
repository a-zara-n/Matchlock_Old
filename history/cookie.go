package history

import "github.com/jinzhu/gorm"

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

func (c Cookie) SetCookie(name string, value string) {
	c.Name = name
	c.Value = value
	db.Table = RequestData{}
	db.Insert(c)
}

func (c Cookie) GetCookie(domain string, path string) []Cookie {
	db.Table = Cookie{}
	cdb := db.OpenDatabase()
	var cookies []Cookie
	cdb.Select("*").
		Where("domain = ? AND path LIKE ? ", domain, path+"%").
		Find(&cookies)
	return cookies
}
