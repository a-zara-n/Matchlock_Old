package decid

import (
	"fmt"

	"github.com/WestEast1st/Matchlock/scanner/attacker/payload"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func Decider(diff diffmatchpatch.Diff, payloadData payload.Payload) {
	if diff.Type == diffmatchpatch.DiffInsert {
		switch payloadData.Division {
		case payload.InspectionInt:
			fmt.Println(payload.InspectionInt)
		case payload.InspectionString:
			fmt.Println(payload.InspectionString)
		case payload.InspectionBool:
			fmt.Println(payload.InspectionBool)
		}
	}
}

func decidInspectionInt() {

}
func decidInspectionString() {

}
func decidInspectionBool() {

}
