package bindings

import (
	"encoding/json"
	ofctx "github.com/OpenFunction/functions-framework-go/openfunction-context"
	"log"
)

func BindingsOutput(ctx *ofctx.OpenFunctionContext, in []byte) ofctx.RetValue {
	var greeting []byte
	if in != nil {
		log.Printf("binding - Data: %s", in)
		greeting = in
	} else {
		log.Print("binding - Data: Received")
		greeting, _ = json.Marshal(map[string]string{"message": "Hello"})
	}

	_, err := ctx.Send("sample", greeting)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return ctx.ReturnWithInternalError()
	}
	return ctx.ReturnWithSuccess()
}
