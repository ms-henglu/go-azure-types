# go-azure-types

## Introduction

Golang's implementation of [Bicep type definitions for ARM resources](https://github.com/Azure/bicep-types-az/tree/main).

## Usage

```go

import (
  "github.com/ms-henglu/go-azure-types/types"
)

func main() {
  azureSchemaLoader := types.DefaultAzureSchemaLoader()
  
  // use customized static files
  // azureSchemaLoader := types.NewAzureSchemaLoader(embeddedFiles)
  
  // list available api-versions
  apiVersions := azureSchemaLoader.ListApiVersions("Microsoft.Compute/virtualMachines")
  
  // get the resource definition for a specific api-version
  resourceDefinition, err := azureSchemaLoader.GetResourceDefinition("Microsoft.Compute/virtualMachines", "2021-03-01")
  
  // list resource functions
  resourceFunctions, err := azureSchemaLoader.ListResourceFunctions("Microsoft.Compute/virtualMachines", "2021-03-01")

  // get the function definition for a specific function
  functionDefinition, err := azureSchemaLoader.GetResourceFunction("Microsoft.Compute/virtualMachines", "2021-03-01", "start")
}

```
