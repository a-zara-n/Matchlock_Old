package value

//CheckListInterface は
type CheckListInterface interface {
	Check(key string)
	Find(key string) bool
}

//CheckList は
type CheckList struct {
	List map[string]bool
}

//NewCheckList は
func NewCheckList() CheckListInterface {
	return &CheckList{
		List: map[string]bool{},
	}
}

//Check は
func (checkl *CheckList) Check(key string) {
	checkl.List[key] = true
}

//Find は
func (checkl *CheckList) Find(key string) bool {
	return checkl.List[key]
}
