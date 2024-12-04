package types

var _ TypeBase = &ResourceFunctionType{}

type ResourceFunctionType struct {
	Type         string         `json:"$type"`
	Name         string         `json:"name"`
	ResourceType string         `json:"resourceType"`
	ApiVersion   string         `json:"apiVersion"`
	Input        *TypeReference `json:"input"`
	Output       *TypeReference `json:"output"`
}

func (t *ResourceFunctionType) FilterConfigurableFields(i interface{}) interface{} {
	return i
}

func (t *ResourceFunctionType) AsTypeBase() *TypeBase {
	typeBase := TypeBase(t)
	return &typeBase
}
