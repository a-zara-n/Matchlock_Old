package value

//HTMLElementInterface は
type HTMLElementInterface interface {
}

//HTMLElement は
type HTMLElement struct {
}

//NewHTMLElement は
func NewHTMLElement() HTMLElementInterface {

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

//AttrInterface は
type AttrInterface interface {
}

//Attr は
type Attr struct {
}

//NewAttr は
func NewAttr() AttrInterface {

	return &Attr{}
}
