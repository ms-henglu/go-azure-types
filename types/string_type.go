package types

import (
	"fmt"
	"log"
	"regexp"
)

var _ TypeBase = &StringType{}

type StringType struct {
	Type      string `json:"$type"`
	MinLength *int   `json:"minLength"`
	MaxLength *int   `json:"maxLength"`
	Sensitive bool   `json:"sensitive"`
	Pattern   string `json:"pattern"`
}

func (s *StringType) Validate(body interface{}, path string) []error {
	if body == nil {
		return nil
	}
	v, ok := body.(string)
	if !ok {
		return []error{ErrorMismatch(path, "string", fmt.Sprintf("%T", body))}
	}
	if v == "" {
		// unknown values will be converted to "", skip validation for now
		// TODO: improve the validation to support unknown values
		return nil
	}
	if s.MinLength != nil && len(v) < *s.MinLength {
		return []error{ErrorCommon(path, fmt.Sprintf("string length is less than %d", *s.MinLength))}
	}
	if s.MaxLength != nil && len(v) > *s.MaxLength {
		return []error{ErrorCommon(path, fmt.Sprintf("string length is greater than %d", *s.MaxLength))}
	}
	if s.Pattern != "" {
		isMatch, err := regexp.Match(s.Pattern, []byte(v))
		if err != nil {
			log.Printf("[WARN] failed to match pattern %s: %s", s.Pattern, err)
			return nil
		}
		if !isMatch {
			return []error{ErrorCommon(path, fmt.Sprintf("string does not match pattern %s", s.Pattern))}
		}
	}
	return nil
}

func (s *StringType) FilterReadOnlyFields(i interface{}) interface{} {
	return i
}

func (s *StringType) FilterConfigurableFields(i interface{}) interface{} {
	return i
}

func (s *StringType) AsTypeBase() *TypeBase {
	typeBase := TypeBase(s)
	return &typeBase
}
