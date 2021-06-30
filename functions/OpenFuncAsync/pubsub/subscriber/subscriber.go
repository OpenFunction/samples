package subscriber

import (
	ofctx "github.com/OpenFunction/functions-framework-go/openfunction-context"
	"log"
)

func Subscriber(ctx *ofctx.OpenFunctionContext, in []byte) int {
	log.Printf("event - Data: %s", in)
	return 200
}
