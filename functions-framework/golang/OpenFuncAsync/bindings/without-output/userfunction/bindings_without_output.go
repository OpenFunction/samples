package userfunction

import (
	ofctx "github.com/OpenFunction/functions-framework-go/openfunction-context"
	"log"
)

func BindingsNoOutput(ctx *ofctx.OpenFunctionContext, in []byte) int {
	if in != nil {
		log.Printf("binding - Data: %s", in)
	} else {
		log.Print("binding - Data: Received")
	}
	return 200
}
