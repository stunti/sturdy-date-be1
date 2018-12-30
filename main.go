package main

import (
	"fmt"

	"github.com/thepauleh/goserverless/serverless"
)

func main() {

	// Create a new CloudFormation template
	template := serverless.NewTemplate("myService")

	template.Service = &serverless.Service {
		Name: "myService",
	}

	template.Provider = &serverless.Provider{
		Name:             "aws",
		Runtime:          "nodejs6.10",
		MemorySize:       512,
		Timeout:          10,
		VersionFunctions: false,
	}

	// An example function
	template.Functions["users"] = serverless.Function {
		Handler:             "service.o",
		Name:                "${self:provider.stage}-users",
		Description:         "Description of what the lambda function does",
		Runtime:             "go1.x",
		MemorySize:          128,
		Timeout:             30,
		Events: []serverless.Events{
			serverless.Events{
				HTTPEvent: &serverless.HTTPEvent {
					Path:   "users/create",
					Method: "post",
				},
			},
		},
	}

	y, err := template.YAML()
	if err != nil {
		fmt.Printf("Failed to generate YAML: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(y))
	}
}
