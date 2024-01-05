package validator

type validator interface {
	Validate() error
}

type allValidator interface {
	ValidateAll() error
}
