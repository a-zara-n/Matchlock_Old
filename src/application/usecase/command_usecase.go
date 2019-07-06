package usecase

type CommandUsecase interface{}
type commandUsecase struct{}

func NewCommandUsecase() CommandUsecase {
	return &commandUsecase{}
}
