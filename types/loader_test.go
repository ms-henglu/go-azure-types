package types

import (
	"testing"

	"github.com/ms-henglu/go-azure-types/embed"
)

func Test_DefaultAzureSchemaLoader(t *testing.T) {
	azureTypes := DefaultAzureSchemaLoader()
	if azureTypes.GetSchema() == nil {
		t.Errorf("failed to load azure schema")
	}
}

func Test_NewAzureSchemaLoader(t *testing.T) {
	azureTypes := NewAzureSchemaLoader(embed.StaticFiles)
	if azureTypes.GetSchema() == nil {
		t.Errorf("failed to load azure schema")
	}
}

func Test_ListApiVersions(t *testing.T) {
	azureTypes := DefaultAzureSchemaLoader()
	case1 := "Microsoft.MachineLearningServices/workspaces/computes"
	if len(azureTypes.ListApiVersions(case1)) == 0 {
		t.Errorf("expect multiple api-version but got 0 for Microsoft.MachineLearningServices/workspaces/computes")
	}

	case2 := "Microsoft.MachineLearningServices/workspaces/computes0"
	if len(azureTypes.ListApiVersions(case2)) != 0 {
		t.Errorf("expect 0 api-version but got multiple for Microsoft.MachineLearningServices/workspaces/computes0")
	}
}

func Test_GetResourceDefinition(t *testing.T) {
	azureTypes := DefaultAzureSchemaLoader()
	case1 := "Microsoft.MachineLearningServices/workspaces/computes"
	versions := azureTypes.ListApiVersions(case1)
	for _, v := range versions {
		def, err := azureTypes.GetResourceDefinition(case1, v)
		if err != nil {
			t.Error(err)
		}
		if def == nil {
			t.Errorf("failed to load resource definition for %s api-version %s", case1, v)
		}
	}
}

func Test_ListResourceFunctions(t *testing.T) {
	azureTypes := DefaultAzureSchemaLoader()
	case1 := "Microsoft.MachineLearningServices/workspaces/computes"
	functions, err := azureTypes.ListResourceFunctions(case1, "2021-01-01")
	if err != nil {
		t.Fatal(err)
	}
	if len(functions) == 0 {
		t.Errorf("expect multiple functions but got 0 for Microsoft.MachineLearningServices/workspaces/computes")
	}
}

func Test_GetResourceFunction(t *testing.T) {
	azureTypes := DefaultAzureSchemaLoader()
	case1 := "Microsoft.MachineLearningServices/workspaces/computes"
	function, err := azureTypes.GetResourceFunction(case1, "2021-01-01", "listNodes")
	if err != nil {
		t.Fatal(err)
	}
	if function == nil {
		t.Errorf("expect a valid function but got nil for Microsoft.MachineLearningServices/workspaces/computes@2021-01-01 listNodes")
	}
}

func Test_AllBicepTypes(t *testing.T) {
	azureTypes := DefaultAzureSchemaLoader()
	schema := azureTypes.GetSchema()
	if schema == nil {
		t.Fatal("failed to load azure schema")
	}
	if len(schema.Resources) == 0 {
		t.Fatal("expect resources are not empty")
	}
	for resourceName, res := range schema.Resources {
		if len(resourceName) == 0 {
			t.Fatal("expect resource name is not empty")
		}
		if res == nil {
			t.Fatalf("expect resource definition is not nil, resource name: %s", resourceName)
		}
		if len(res.Definitions) == 0 {
			t.Fatalf("expect resource definitions are not empty, resource name: %s", resourceName)
		}
		for _, definition := range res.Definitions {
			_, err := definition.GetDefinition()
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
