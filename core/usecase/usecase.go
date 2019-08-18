package usecase

// Usecase interface for usecases
type Usecase interface {
	Execute(param... interface{}) (interface{}, error)
}