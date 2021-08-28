package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/codeSum27/iam/pkg/api"
	"github.com/codeSum27/iam/pkg/common"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"os"
)

func loadConfig(e *echo.Echo) error{

	e.Logger.Print("Load config file from setting.json")

	file, err := os.Open("setting.json")
	defer file.Close()

	if err != nil {
		e.Logger.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&common.Cnf)
	if err != nil {
		e.Logger.Fatal(err)
	}

	return err

}
func main() {
	api.RedisInit()

	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	iamServer := api.NewIamServer()
	// This is how you set up a basic Echo router
	e := echo.New()

	// Log all requests
	e.Use(echomiddleware.LoggerWithConfig(echomiddleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	e.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	api.RegisterHandlers(e, iamServer)
	// And we serve HTTP until the world ends.

	loadConfig(e)
	api.Init()

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}