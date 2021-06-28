# Service invocation

- [Server](#server)
- [Client](#client)
- [Results](#results)

## Server

Prepare a context as follows, name it `function.json`. (You can refer to [OpenFunction Context Specs](https://github.com/OpenFunction/functions-framework/blob/main/docs/OpenFunction-context-specs.md) to learn more about the OpenFunction Context)

```json
{
  "name": "server",
  "version": "v1",
  "requestID": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "port": "50001",
  "input": {
    "name": "echo",
    "uri": "print",
    "params": {
      "type": "invoke"
    }
  },
  "outputs": {},
  "runtime": "OpenFuncAsync"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"server","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"50001","input":{"name":"echo","uri":"print","params":{"type":"invoke"}},"outputs":{},"runtime":"OpenFuncAsync"}'
```

Start the service and watch the logs.
```shell
cd server/
go mod tidy
dapr run --app-id server \
    --app-protocol grpc \
    --app-port 50001 \
    go run ./main.go
```

## Client

You also need a definition of client.

```json
{
  "name": "client",
  "version": "v1",
  "requestID": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "port": "50002",
  "input": {},
  "outputs": {
    "server": {
      "uri": "print",
      "params": {
        "method": "post",
        "type": "invoke"
      }
    }
  },
  "runtime": "OpenFuncAsync"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"client","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"50002","input":{},"outputs":{"server":{"uri":"print","params":{"method":"post","type":"invoke"}}},"runtime":"OpenFuncAsync"}'
```

Start the client to post request.

```shell
cd client/
go mod tidy
dapr run --app-id client \
    --app-protocol grpc \
    go run ./main.go
```

## Results

<details>
<summary>View detailed logs.</summary>

```shell
== APP == 2021/06/28 11:50:19 invoke - Data: hello
```
</details>

