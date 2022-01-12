package userfunction

import (
    "context"
    "fmt"
    cloudevents "github.com/cloudevents/sdk-go/v2"
)

func HelloWorld(ctx context.Context, ce cloudevents.Event) error {
    fmt.Println(string(ce.Data()))
    return nil
}