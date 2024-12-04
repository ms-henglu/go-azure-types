package types

import (
	"encoding/json"
	"fmt"
)

var _ TypeBase = &ObjectType{}

type ObjectType struct {
	Type                 string                    `json:"$type"`
	Name                 string                    `json:"name"`
	Properties           map[string]ObjectProperty `json:"properties"`
	AdditionalProperties *TypeReference            `json:"additionalProperties"`
	Sensitive            bool                      `json:"sensitive"`
}

func (t *ObjectType) FilterConfigurableFields(body interface{}) interface{} {
	if t == nil || body == nil {
		return nil
	}
	// check body type
	bodyMap, ok := body.(map[string]interface{})
	if !ok {
		return body
	}

	res := make(map[string]interface{})
	for key, def := range t.Properties {
		if _, ok := bodyMap[key]; ok {
			if (def.IsRequired() || (!def.IsReadOnly() && !def.IsDeployTimeConstant())) && def.Type != nil && def.Type.Type != nil {
				res[key] = (*def.Type.Type).FilterConfigurableFields(bodyMap[key])
			}
		}
	}

	if t.AdditionalProperties != nil && t.AdditionalProperties.Type != nil {
		for key, value := range bodyMap {
			if _, ok := t.Properties[key]; ok {
				continue
			}
			res[key] = (*t.AdditionalProperties.Type).FilterConfigurableFields(value)
		}
	}
	return res
}

func (t *ObjectType) AsTypeBase() *TypeBase {
	typeBase := TypeBase(t)
	return &typeBase
}

type ObjectProperty struct {
	Type        *TypeReference
	Flags       []ObjectPropertyFlag
	Description *string
}

func (o *ObjectProperty) IsRequired() bool {
	for _, value := range o.Flags {
		if value == Required {
			return true
		}
	}
	return false
}

func (o *ObjectProperty) IsReadOnly() bool {
	for _, value := range o.Flags {
		if value == ReadOnly {
			return true
		}
	}
	return false
}

func (o *ObjectProperty) IsDeployTimeConstant() bool {
	for _, value := range o.Flags {
		if value == DeployTimeConstant {
			return true
		}
	}
	return false
}

func (o *ObjectProperty) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "description":
			if v != nil {
				var description string
				err := json.Unmarshal(*v, &description)
				if err != nil {
					return err
				}
				o.Description = &description
			}
		case "flags":
			if v != nil {
				var flag int
				err := json.Unmarshal(*v, &flag)
				if err != nil {
					return err
				}
				flags := make([]ObjectPropertyFlag, 0)
				for _, f := range PossibleObjectPropertyFlagValues() {
					if flag&int(f) != 0 {
						flags = append(flags, f)
					}
				}
				o.Flags = flags
			}
		case "type":
			if v != nil {
				var typeRef TypeReference
				err := json.Unmarshal(*v, &typeRef)
				if err != nil {
					return err
				}
				o.Type = &typeRef
			}
		default:
			return fmt.Errorf("unmarshalling object property, unrecognized key: %s", k)
		}
	}
	return nil
}

type ObjectPropertyFlag int

const (
	None ObjectPropertyFlag = 0

	Required ObjectPropertyFlag = 1 << 0

	ReadOnly ObjectPropertyFlag = 1 << 1

	WriteOnly ObjectPropertyFlag = 1 << 2

	DeployTimeConstant ObjectPropertyFlag = 1 << 3

	Identifier ObjectPropertyFlag = 1 << 4
)

func PossibleObjectPropertyFlagValues() []ObjectPropertyFlag {
	return []ObjectPropertyFlag{None, Required, ReadOnly, WriteOnly, DeployTimeConstant, Identifier}
}
