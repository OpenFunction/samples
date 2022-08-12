package hello

import (
	"context"
	"encoding/json"
	"net/http"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
	"github.com/OpenFunction/functions-framework-go/functions"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"k8s.io/klog/v2"
)

func init() {
	functions.HTTP("Hello", hello,
		functions.WithFunctionPath("/hello/{name}"),
		functions.WithFunctionMethods("GET", "POST"),
	)

	functions.CloudEvent("Foo", foo,
		functions.WithFunctionPath("/foo/{name}"),
	)

	functions.OpenFunction("Bar", bar,
		functions.WithFunctionPath("/bar/{name}"),
		functions.WithFunctionMethods("GET", "POST"),
	)
}

func hello(w http.ResponseWriter, r *http.Request) {
	vars := ofctx.VarsFromCtx(r.Context())
	response := map[string]string{
		"hello": vars["name"],
	}
	responseBytes, _ := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.Write(responseBytes)
}

func foo(ctx context.Context, ce cloudevents.Event) error {
	vars := ofctx.VarsFromCtx(ctx)
	response := map[string]string{
		string(ce.Data()): vars["name"],
	}
	responseBytes, _ := json.Marshal(response)
	klog.Infof("cloudevent - Data: %s", string(responseBytes))
	return nil
}

func bar(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	vars := ofctx.VarsFromCtx(ctx.GetNativeContext())
	response := map[string]string{
		string(in): vars["name"],
	}
	responseBytes, _ := json.Marshal(response)
	return ctx.ReturnOnSuccess().WithData(responseBytes), nil
}
