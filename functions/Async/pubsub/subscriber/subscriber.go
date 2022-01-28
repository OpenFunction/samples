package subscriber

import (
	"log"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
)

func Subscriber(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	log.Printf("event - Data: %s", in)
	return ctx.ReturnOnSuccess(), nil
}
