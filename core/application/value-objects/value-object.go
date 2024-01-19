package valueobjects

type ValueObject interface {
	Equals(obj ValueObject) bool
	ToString() string
}
