package sender

import (
	"encoding/json"
	"log"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
)

func ForwardToKafka(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	var greeting []byte
	if in != nil {
		log.Printf("http - Data: %s", in)
		greeting = in
	} else {
		log.Print("http - Data: Received")
		greeting, _ = json.Marshal(map[string]string{"message": "Hello"})
	}

	_, err := ctx.Send("kafka-server", greeting)
	if err != nil {
		log.Print(err.Error())
		return ctx.ReturnOnInternalError(), err
	}

	return ctx.ReturnOnSuccess(), nil
}
