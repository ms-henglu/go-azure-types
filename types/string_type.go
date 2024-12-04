package types

var _ TypeBase = &StringType{}

type StringType struct {
	Type      string `json:"$type"`
	MinLength *int   `json:"minLength"`
	MaxLength *int   `json:"maxLength"`
	Sensitive bool   `json:"sensitive"`
	Pattern   string `json:"pattern"`
}

func (s *StringType) FilterConfigurableFields(i interface{}) interface{} {
	return i
}

func (s *StringType) AsTypeBase() *TypeBase {
	typeBase := TypeBase(s)
	return &typeBase
}
