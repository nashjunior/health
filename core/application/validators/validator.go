package validators

type ErrorField struct {
	Name  string
	Error string
}

type IValidator interface {
	Errors() *[]ErrorField
	Validate() bool
}
