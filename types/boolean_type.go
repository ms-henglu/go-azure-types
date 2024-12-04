package types

var _ TypeBase = &BooleanType{}

type BooleanType struct {
	Type string `json:"$type"`
}

func (t *BooleanType) FilterConfigurableFields(i interface{}) interface{} {
	return i
}

func (t *BooleanType) AsTypeBase() *TypeBase {
	typeBase := TypeBase(t)
	return &typeBase
}
