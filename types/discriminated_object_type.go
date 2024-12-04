package types

var _ TypeBase = &DiscriminatedObjectType{}

type DiscriminatedObjectType struct {
	Type           string                    `json:"$type"`
	Name           string                    `json:"name"`
	Discriminator  string                    `json:"discriminator"`
	BaseProperties map[string]ObjectProperty `json:"baseProperties"`
	Elements       map[string]*TypeReference `json:"elements"`
}

func (t *DiscriminatedObjectType) FilterConfigurableFields(body interface{}) interface{} {
	if t == nil || body == nil {
		return []error{}
	}
	// check body type
	bodyMap, ok := body.(map[string]interface{})
	if !ok {
		return nil
	}

	res := make(map[string]interface{})
	for key, def := range t.BaseProperties {
		if _, ok := bodyMap[key]; ok {
			if !def.IsReadOnly() && def.Type != nil && def.Type.Type != nil {
				res[key] = (*def.Type.Type).FilterConfigurableFields(bodyMap[key])
			}
		}
	}

	if _, ok := bodyMap[t.Discriminator]; !ok {
		return nil
	}

	if discriminator, ok := bodyMap[t.Discriminator].(string); ok {
		if t.Elements[discriminator] != nil && t.Elements[discriminator].Type != nil {
			if additionalProps := (*t.Elements[discriminator].Type).FilterConfigurableFields(body); additionalProps != nil {
				if additionalMap, ok := additionalProps.(map[string]interface{}); ok {
					for key, value := range additionalMap {
						res[key] = value
					}
					return res
				}
			}
		}
	}

	// if the discriminator's type is not in the embedded schema, add unchecked properties to res
	for key, value := range bodyMap {
		if _, ok := t.BaseProperties[key]; ok {
			continue
		}
		res[key] = value
	}
	return res
}

func (t *DiscriminatedObjectType) AsTypeBase() *TypeBase {
	typeBase := TypeBase(t)
	return &typeBase
}
