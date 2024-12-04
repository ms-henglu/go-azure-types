package types

import (
	"fmt"
)

var _ TypeBase = &StringLiteralType{}

type StringLiteralType struct {
	Type  string `json:"$type"`
	Value string `json:"value"`
}

func (t *StringLiteralType) Validate(body interface{}, path string) []error {
	if t == nil || body == nil {
		return []error{}
	}
	errors := make([]error, 0)
	if stringValue, ok := body.(string); ok {
		if stringValue != t.Value {
			errors = append(errors, ErrorMismatch(path, t.Value, stringValue))
		}
	} else {
		errors = append(errors, ErrorMismatch(path, "string", fmt.Sprintf("%T", body)))
	}
	return errors
}

func (t *StringLiteralType) FilterReadOnlyFields(i interface{}) interface{} {
	return i
}

func (t *StringLiteralType) FilterConfigurableFields(i interface{}) interface{} {
	return i
}

func (t *StringLiteralType) AsTypeBase() *TypeBase {
	typeBase := TypeBase(t)
	return &typeBase
}
