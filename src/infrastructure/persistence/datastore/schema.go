package datastore

import "github.com/jinzhu/gorm"

type HistorySchema struct {
	gorm.Model
	Identifier string
	IsEdit     bool
}

//RequestInfoSchema はリクエストの各種情報を保存するschemaです
type RequestInfoSchema struct {
	gorm.Model
	Identifier string
	IsEdit     bool
	Host       string
	Method     string
	URL        string
	Path       string
	Proto      string
}

//RequestHeaderSchema はヘッダー情報を保存するschemaです
type RequestHeaderSchema struct {
	gorm.Model
	Identifier string
	IsEdit     bool
	Name       string
	Value      string
}

//RequestDataSchema はQuery情報を保存するschemaです
type RequestDataSchema struct {
	gorm.Model
	Identifier string
	IsEdit     bool
	Parents    string
	Name       string
	Value      string
	Style      string //JSON / FORM
	Type       string // ex. int string float
}

//ResponseInfoSchema はレスポンスの各種情報を保存するschemaです
type ResponseInfoSchema struct {
	gorm.Model
	Identifier string
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
}

//ResponseHeaderSchema はヘッダー情報を保存するschemaです
type ResponseHeaderSchema struct {
	gorm.Model
	Identifier string
	Name       string
	Value      string
}

//ResponseBodySchema はレスポンスの各種情報を保存するschemaです
type ResponseBodySchema struct {
	gorm.Model
	Identifier string
	Body       string
	Encodetype string
	Length     int64
}
