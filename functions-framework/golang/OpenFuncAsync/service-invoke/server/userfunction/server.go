package userfunction

import (
	ofctx "github.com/OpenFunction/functions-framework-go/openfunction-context"
	"log"
)

func Server(ctx *ofctx.OpenFunctionContext, in []byte) int {
	if in != nil {
		log.Printf("invoke - Data: %s", in)
	} else {
		log.Print("invoke - Data: Received")
	}
	return 200
}
