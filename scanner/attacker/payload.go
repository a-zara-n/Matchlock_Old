package attacker

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
	payloads    = map[string][]string{
		InspectionInt:    {"operator", "zero", "float", "bigNumbers"},
		InspectionString: {"tagstring", "special", "urlstring", "script", "event", "sql", "javascript", "command"},
		InspectionBool:   {"bool"},
	}
)

type payload struct {
	payload [][]string
}
