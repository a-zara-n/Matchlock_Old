package entity

import "net/url"

//RequestInfo は
type RequestInfo struct {
	Host   string
	Method string
	URL    *url.URL
	Path   string
	Proto  string
}
