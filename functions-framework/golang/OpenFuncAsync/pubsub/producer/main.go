package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/OpenFunction/functions-framework-go/functionframeworks"
	ofctx "github.com/OpenFunction/functions-framework-go/openfunction-context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"log"
	"main.go/userfunction"
	"net/http"
)

func register(fn interface{}) error {
	ctx := context.Background()
	if fnHTTP, ok := fn.(func(http.ResponseWriter, *http.Request)); ok {
		if err := functionframeworks.RegisterHTTPFunction(ctx, fnHTTP); err != nil {
			return fmt.Errorf("Function failed to register: %v\n", err)
		}
	} else if fnCloudEvent, ok := fn.(func(context.Context, cloudevents.Event) error); ok {
		if err := functionframeworks.RegisterCloudEventFunction(ctx, fnCloudEvent); err != nil {
			return fmt.Errorf("Function failed to register: %v\n", err)
		}
	} else if fnOpenFunction, ok := fn.(func(*ofctx.OpenFunctionContext, []byte) ofctx.RetValue); ok {
		if err := functionframeworks.RegisterOpenFunction(ctx, fnOpenFunction); err != nil {
			return fmt.Errorf("Function failed to register: %v\n", err)
		}
	} else {
		err := errors.New("unrecognized function")
		return fmt.Errorf("Function failed to register: %v\n", err)
	}
	return nil
}

func main() {
	if err := register(userfunction.Producer); err != nil {
		log.Fatalf("Failed to register: %v\n", err)
	}

	if err := functionframeworks.Start(); err != nil {
		log.Fatalf("Failed to start: %v\n", err)
	}
}
