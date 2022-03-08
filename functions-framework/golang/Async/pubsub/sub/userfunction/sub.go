package userfunction

import (
	"encoding/json"
	"log"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
)

func Subscriber(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	msg := map[string]string{}
	json.Unmarshal(in, &msg)
	log.Printf("event - Data: %s", string(in))
	log.Printf("event - Data: %s", msg)
	return ctx.ReturnOnSuccess(), nil
}
