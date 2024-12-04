package types

var _ TypeBase = &UnionType{}

type UnionType struct {
	Type     string           `json:"$type"`
	Elements []*TypeReference `json:"elements"`
}

func (t *UnionType) FilterConfigurableFields(i interface{}) interface{} {
	return i
}

func (t *UnionType) AsTypeBase() *TypeBase {
	typeBase := TypeBase(t)
	return &typeBase
}
