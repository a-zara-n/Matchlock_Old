package service

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/a-zara-n/Matchlock/src/domain/value"
	"github.com/sergi/go-diff/diffmatchpatch"
	"golang.org/x/net/html"
)

type Judgmenter interface {
	Run()
}

type judgmenter struct {
}

func NewJudgmenter() {

}

func (jud *judgmenter) Run(diffs []diffmatchpatch.Diff, payloadData value.Payload, req http.Request, inputstr string) {
	switch payloadData.Division {
	case value.InspectionInt:
	case value.InspectionString:
		switch payloadData.Type {
		case "tagstring":
			jud.xss(diffs, payloadData.GetPayload(), req, inputstr, tagstring)
		case "special":
			jud.xss(diffs, payloadData.GetPayload(), req, inputstr, special)
		case "event":
			jud.xss(diffs, payloadData.GetPayload(), req, inputstr, event)
		case "sql":
		case "command":
		}
	case value.InspectionBool:
		//fmt.Println(payload.InspectionBool)
	}
}

func (jud *judgmenter) xss(diffs []diffmatchpatch.Diff, data []string, req http.Request, inputstr string, fillter func(string, []string, http.Request, string)) {
	for _, diff := range diffs {
		if diff.Type == diffmatchpatch.DiffInsert {
			fillter(diff.Text, data, req, inputstr)
		}
	}
}

func tagstring(text string, data []string, req http.Request, inputstr string) {
	for _, v := range data {
		if strings.Contains(text, v) {
			massage := newAlertMessage(req.Method, req.URL.String(), "Existence of unescaped tag character")
			var isOut bool
			text = strings.Trim(text, " ")
			for _, str := range strings.Split(html.UnescapeString(inputstr), "&") {
				if strings.Contains(str, v) {
					isOut = true
					massage = append(massage, "		- Input  :"+str)
				}
			}
			if isOut {
				for _, mes := range massage {
					fmt.Println(mes)
				}
				fmt.Println("		- Output :", text)
			}
			break
		}
	}
}

func special(text string, data []string, req http.Request, inputstr string) {
	v := []string{`&lt;&gt;&quot;&apos;&amp;`, `&lt;&gt;&quot;&apos;`}

	if !strings.Contains(text, v[0]) || !strings.Contains(text, v[1]) {
		massage := newAlertMessage(req.Method, req.URL.String(), "HTML Special char has not been escaped")
		var isOut bool
		text = strings.Trim(text, " ")
		for _, str := range strings.Split(html.UnescapeString(inputstr), "&") {
			if strings.Contains(str, v[0]) {
				isOut = true
				massage = append(massage, "		- Input  :"+str)
			}
		}
		if isOut {
			for _, mes := range massage {
				fmt.Println(mes)
			}
			fmt.Println("		- Output :", text)
		}
	}
}

func event(text string, data []string, req http.Request, inputstr string) {
	for _, v := range data {
		r := regexp.MustCompile(`.*<.*` + v[2:] + `.*>.*`)
		if strings.Contains(strings.Trim(text, " "), v[2:]) {
			if r.MatchString(text) {
				var isOut bool
				massage := newAlertMessage(req.Method, req.URL.String(), "Event handler and double quote enabled in tag")
				text = strings.Trim(text, " ")
				for _, str := range strings.Split(html.UnescapeString(inputstr), "&") {
					if strings.Contains(str, v) && strings.Contains(text, strings.Split(str, "=")[0]) {
						isOut = true
						massage = append(massage, "		- Input  :"+str)
					}
				}
				if isOut {
					for _, mes := range massage {
						fmt.Println(mes)
					}
					fmt.Println("		- Output :", text)
				}
			}
			break
		}
	}
}

func newAlertMessage(method, url, message string) []string {
	return []string{"# [INFO] " + message, "	Method: " + method, "	URL: " + url}
}
