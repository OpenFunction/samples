package userfunction

import (
	"log"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
)

func BindingsOutput(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	var greeting []byte
	if in != nil {
		log.Printf("binding - Data: %s", in)
		greeting = in
	} else {
		log.Print("binding - Data: Received")
		greeting = []byte("Hello")
	}

	_, err := ctx.Send("echo", greeting)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return ctx.ReturnOnInternalError(), err
	}
	return ctx.ReturnOnSuccess(), nil
}
