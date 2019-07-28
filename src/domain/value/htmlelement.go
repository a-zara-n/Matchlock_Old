package value

import "golang.org/x/net/html"

//FormElementInterface は
type FormElementInterface interface {
}

//HTMLElement は
type HTMLElement struct {
	Tag       string
	Attr      []html.Attribute
	ParamData []HTMLElement
}

//NewFormElement は
func NewFormElement() FormElementInterface {

	return &HTMLElement{}
}

//LinkStringInterface は
type LinkStringInterface interface {
}

//LinkString は
type LinkString struct {
}

//NewLinkString は
func NewLinkString() LinkStringInterface {

	return &LinkString{}
}
