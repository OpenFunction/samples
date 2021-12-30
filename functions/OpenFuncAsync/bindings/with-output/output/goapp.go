package bindings

import (
    "encoding/json"
    "fmt"
    ofctx "github.com/OpenFunction/functions-framework-go/openfunction-context"
)

func OutputTarget(ctx *ofctx.OpenFunctionContext, in []byte) ofctx.RetValue {
    var msg Message
    err := json.Unmarshal(in, &msg)
    if err != nil {
        fmt.Println("error reading message from Kafka binding", err)
        return ctx.ReturnWithInternalError()
    }
    fmt.Printf("message from Kafka '%s'\n", msg)
    return ctx.ReturnWithSuccess()
}

type Message struct {
    Msg string `json:"message"`
}
