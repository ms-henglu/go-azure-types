package types

var _ TypeBase = &IntegerType{}

type IntegerType struct {
	Type     string `json:"$type"`
	MinValue *int   `json:"minValue"`
	MaxValue *int   `json:"maxValue"`
}

func (t *IntegerType) FilterConfigurableFields(i interface{}) interface{} {
	return i
}

func (t *IntegerType) AsTypeBase() *TypeBase {
	typeBase := TypeBase(t)
	return &typeBase
}
