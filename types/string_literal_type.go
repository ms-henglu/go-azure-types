package types

var _ TypeBase = &StringLiteralType{}

type StringLiteralType struct {
	Type  string `json:"$type"`
	Value string `json:"value"`
}

func (t *StringLiteralType) FilterConfigurableFields(i interface{}) interface{} {
	return i
}

func (t *StringLiteralType) AsTypeBase() *TypeBase {
	typeBase := TypeBase(t)
	return &typeBase
}
