package payload

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
	PayloadPath      = "./payload/"
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
		InspectionString: {"tagstring", "special", "urlstring", "script", "event", "sql", "javascript", "command", "float", "bigNumbers"},
		InspectionBool:   {"bool"},
	}
)

type Payload struct {
	Division string
	Type     string
	Data     []string
}

func (p *Payload) GetTypeKeys(t string) []string {
	return types[t]
}

func (p *Payload) GetFileName(filetype string) []string {
	return payloads[filetype]
}

func (p *Payload) GetPayload(key string, name string) []string {
	f, err := os.Open(PayloadPath + key + name + Text)
	if err != nil {
		fmt.Println(PayloadPath + key + name + Text)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	return strings.Split(string(b), "\n")
}
