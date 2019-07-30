package value

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
InspectionInt : File Path Const
*/
const (
	PayloadPath      = "./_payload/"
	InspectionInt    = "inspection/int/"
	InspectionString = "inspection/string/"
	InspectionBool   = "inspection/bool/"
	Text             = ".txt"
)

var (
	inspections = []string{InspectionInt, InspectionString, InspectionBool}
	types       = map[string][]string{
		"inspection": inspections,
	}
	payloads = map[string][]string{
		InspectionInt:    {"operator"},
		InspectionString: {"tagstring", "special", "event", "sql"},
		InspectionBool:   {"bool"},
	}
)

//PayloadInterface は
type PayloadInterface interface {
	GetTypeKeys(t string) []string
	GetFileName(filetype string) []string
	GetPayload() []string
	SetInfo(key, name string)
}

//Payload は
type Payload struct {
	Division string
	Type     string
	Data     []string
}

//NewPayload は
func NewPayload(tys string) PayloadInterface {
	return &Payload{}
}

func (p *Payload) GetTypeKeys(t string) []string {
	return types[t]
}

func (p *Payload) GetFileName(filetype string) []string {
	return payloads[filetype]
}
func (p *Payload) SetInfo(key, name string) {
	p.Division = key
	p.Type = name
}
func (p *Payload) GetPayload() []string {
	f, err := os.Open(PayloadPath + p.Division + p.Type + Text)
	if err != nil {
		fmt.Println(PayloadPath + p.Division + p.Type + Text)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	return strings.Split(string(b), "\n")
}
