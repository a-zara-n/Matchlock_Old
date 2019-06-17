package shared

import "testing"

func Test_randomString(t *testing.T) {
	var isEquals = []bool{}
	for _, input := range []int{1, 3, 5, 7, 10} {
		isEquals = append(isEquals, randomString(input) == randomString(input))
	}
	problemConfirmation(t, isEquals)
}

func Test_GetSha1(t *testing.T) {
	var isEquals = []bool{}
	for _, input := range []string{"Alice", "alpha", "Bob", "blabo"} {
		isEquals = append(isEquals, GetSha1(input) == GetSha1(input))
	}
	problemConfirmation(t, isEquals)
}

func problemConfirmation(t *testing.T, isEquals []bool) {
	var c int
	for _, isEqual := range isEquals {
		if isEqual == true {
			c++
		}
	}
	if c > 3 {
		t.Error("The same value occupies half")
	} else {
		t.Log("Pass!")
	}
}
