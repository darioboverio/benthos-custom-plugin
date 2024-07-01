package main

import (
	"context"

	"github.com/redpanda-data/benthos/v4/public/service"

	// Import full suite of FOSS connect plugins
	_ "github.com/redpanda-data/connect/public/bundle/free/v4"

	// Add your plugin packages here
	_ "github.com/darioboverio/benthos-custom-plugin/bloblang"
)

func main() {
	service.RunCLI(context.Background())
}
