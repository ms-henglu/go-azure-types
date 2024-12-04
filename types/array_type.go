package types

var _ TypeBase = &ArrayType{}

type ArrayType struct {
	Type      string         `json:"$type"`
	ItemType  *TypeReference `json:"itemType"`
	MinLength *int           `json:"minLength"`
	MaxLength *int           `json:"maxLength"`
}

func (t *ArrayType) Validate(i interface{}, s string) []error {
	//TODO implement me
	panic("implement me")
}

func (t *ArrayType) FilterReadOnlyFields(i interface{}) interface{} {
	if t == nil || i == nil {
		return nil
	}
	var itemType *TypeBase
	if t.ItemType != nil {
		itemType = t.ItemType.Type
	}
	// check body type
	bodyArray, ok := i.([]interface{})
	if !ok {
		return nil
	}
	if itemType == nil {
		return bodyArray
	}

	res := make([]interface{}, 0)
	for _, value := range bodyArray {
		res = append(res, (*itemType).FilterReadOnlyFields(value))
	}
	return res
}

func (t *ArrayType) FilterConfigurableFields(i interface{}) interface{} {
	if t == nil || i == nil {
		return nil
	}
	var itemType *TypeBase
	if t.ItemType != nil {
		itemType = t.ItemType.Type
	}
	// check body type
	bodyArray, ok := i.([]interface{})
	if !ok {
		return nil
	}
	res := make([]interface{}, 0)
	for _, value := range bodyArray {
		if itemType != nil {
			res = append(res, (*itemType).FilterConfigurableFields(value))
		}
	}
	return res
}

func (t *ArrayType) AsTypeBase() *TypeBase {
	typeBase := TypeBase(t)
	return &typeBase
}
