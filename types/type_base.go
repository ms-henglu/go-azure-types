package types

type TypeBase interface {
	AsTypeBase() *TypeBase

	FilterConfigurableFields(interface{}) interface{}

	FilterReadOnlyFields(interface{}) interface{}

	Validate(interface{}, string) []error
}
