package attacker

import (
	"github.com/sergi/go-diff/diffmatchpatch"
)

func lineDiff(src1, src2 string) []diffmatchpatch.Diff {
	dmp := diffmatchpatch.New()
	a, b, c := dmp.DiffLinesToChars(src1, src2)
	return dmp.DiffCharsToLines(dmp.DiffMain(a, b, false), c)
}
