package decid

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
	"strings"

	"github.com/a-zara-n/Matchlock/scanner/attacker/payload"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func Decider(diffs []diffmatchpatch.Diff, payloadData payload.Payload, req http.Request, inputstr string) {
	//fmt.Println("hoge")
	switch payloadData.Division {
	case payload.InspectionInt:
		//fmt.Println(payload.InspectionInt)
	case payload.InspectionString:
		switch payloadData.Type {
		case "tagstring":
			for _, diff := range diffs {
				if diff.Type == diffmatchpatch.DiffInsert {
					tagstring(diff.Text, payloadData.Data, req, inputstr)
				}
			}
		case "special":
			for _, diff := range diffs {
				if diff.Type == diffmatchpatch.DiffInsert {
					special(diff.Text, req, inputstr)
				}
			}
		case "event":
			for _, diff := range diffs {
				if diff.Type == diffmatchpatch.DiffInsert {
					event(diff.Text, payloadData.Data, req, inputstr)
				}
			}
		case "sql":
			sql(diffs, payloadData.Data, req, inputstr)
		case "command":
		}
	case payload.InspectionBool:
		//fmt.Println(payload.InspectionBool)
	}
}

func tagstring(text string, data []string, req http.Request, inputstr string) {
	for _, v := range data {
		if strings.Contains(text, v) {
			massage := []string{
				"# [INFO] Existence of unescaped tag character",
				"	Method" + req.Method,
				"	URL:" + req.URL.String(),
			}
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

func special(text string, req http.Request, inputstr string) {
	v := []string{`&lt;&gt;&quot;&apos;&amp;`, `&lt;&gt;&quot;&apos;`}

	if !strings.Contains(text, v[0]) || !strings.Contains(text, v[1]) {
		massage := []string{
			"# [INFO] HTML Special char has not been escaped",
			"	Method" + req.Method,
			"	URL:" + req.URL.String(),
		}
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
				massage := []string{
					"# [INFO] Event handler and double quote enabled in tag",
					"	Method: " + req.Method,
					"	URL: " + req.URL.String(),
				}
				var isOut bool
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
func sql(diffs []diffmatchpatch.Diff, data []string, req http.Request, inputstr string) {
	count_map := map[diffmatchpatch.Operation]int{
		diffmatchpatch.DiffInsert: 0,
		diffmatchpatch.DiffEqual:  0,
		diffmatchpatch.DiffDelete: 0,
	}
	c := 0
	for _, diff := range diffs {
		count_map[diff.Type]++
		c++
	}
}
