package types

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"

	typesEmbed "github.com/ms-henglu/go-azure-types/embed"
)

func DefaultAzureSchemaLoader() *AzureSchemaLoader {
	return &AzureSchemaLoader{
		staticFiles: typesEmbed.StaticFiles,
		mutex:       sync.Mutex{},
	}
}

func NewAzureSchemaLoader(staticFiles embed.FS) *AzureSchemaLoader {
	return &AzureSchemaLoader{
		staticFiles: staticFiles,
		mutex:       sync.Mutex{},
	}
}

type AzureSchemaLoader struct {
	schema      *Schema
	mutex       sync.Mutex
	staticFiles embed.FS
}

func (r *AzureSchemaLoader) GetSchema() *Schema {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.schema == nil {
		data, err := r.staticFiles.ReadFile("generated/index.json")
		if err != nil {
			log.Printf("[ERROR] failed to load schema index: %+v", err)
			return nil
		}
		err = json.Unmarshal(data, &r.schema)
		if err != nil {
			log.Printf("[ERROR] failed to unmarshal schema index: %+v", err)
			return nil
		}

		for _, resource := range r.schema.Resources {
			for _, definition := range resource.Definitions {
				definition.Location.StaticFiles = &r.staticFiles
			}
		}
		for _, function := range r.schema.Functions {
			for _, definition := range function.Definitions {
				definition.Location.StaticFiles = &r.staticFiles
			}
		}
	}
	return r.schema
}

func (r *AzureSchemaLoader) ListApiVersions(resourceType string) []string {
	azureSchema := r.GetSchema()
	if azureSchema == nil {
		return []string{}
	}
	res := make([]string, 0)
	for key, value := range azureSchema.Resources {
		if strings.EqualFold(key, resourceType) {
			for _, v := range value.Definitions {
				res = append(res, v.ApiVersion)
			}
		}
	}

	// TODO: remove the below codes when Resources RP 2024-07-01 is available
	if strings.EqualFold(resourceType, "Microsoft.Resources/resourceGroups") {
		temp := make([]string, 0)
		for _, v := range res {
			if v != "2024-07-01" {
				temp = append(temp, v)
			}
		}
		res = temp
	}

	sort.Strings(res)
	return res
}

func (r *AzureSchemaLoader) GetResourceDefinition(resourceType, apiVersion string) (*ResourceType, error) {
	azureSchema := r.GetSchema()
	if azureSchema == nil {
		return nil, fmt.Errorf("failed to load azure schema index")
	}
	for key, value := range azureSchema.Resources {
		if strings.EqualFold(key, resourceType) {
			for _, v := range value.Definitions {
				if v.ApiVersion == apiVersion {
					return v.GetDefinition()
				}
			}
		}
	}
	return nil, fmt.Errorf("failed to find resource type %s api-version %s in azure schema index", resourceType, apiVersion)
}

func (r *AzureSchemaLoader) ListResourceFunctions(resourceType, apiVersion string) ([]*FunctionDefinition, error) {
	res := make([]*FunctionDefinition, 0)
	azureSchema := r.GetSchema()
	if azureSchema == nil {
		return nil, fmt.Errorf("failed to load azure schema index")
	}
	for key, value := range azureSchema.Functions {
		if strings.EqualFold(key, resourceType) {
			for _, v := range value.Definitions {
				if v.ApiVersion == apiVersion {
					res = append(res, v)
				}
			}
		}
	}

	return res, nil
}

func (r *AzureSchemaLoader) GetResourceFunction(resourceType, apiVersion, name string) (*ResourceFunctionType, error) {
	azureSchema := r.GetSchema()
	if azureSchema == nil {
		return nil, fmt.Errorf("failed to load azure schema index")
	}
	for key, value := range azureSchema.Functions {
		if strings.EqualFold(key, resourceType) {
			for _, v := range value.Definitions {
				if v.ApiVersion == apiVersion {
					def, err := v.GetDefinition()
					if err == nil && def.Name == name {
						return def, nil
					}
				}
			}
		}
	}
	return nil, nil
}
