package command

type Command interface {
	Run()
}
type command struct{}

func NewCommand() Command {
	return &command{}
}

func (c *command) Run() {

}
