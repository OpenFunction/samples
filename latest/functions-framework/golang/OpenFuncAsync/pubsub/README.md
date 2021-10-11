# Pubsub

- [Subscriber](#subscriber)
- [Producer](#producer)
- [Results](#results)

## Subscriber

Prepare a context as follows, name it `function.json`. (You can refer to [OpenFunction Context Specs](https://github.com/OpenFunction/functions-framework/blob/main/docs/OpenFunction-context-specs.md) to learn more about the OpenFunction Context)

```json
{
  "name": "subscriber",
  "version": "v1",
  "requestID": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "port": "60011",
  "inputs": {
    "sub": {
      "uri": "my_topic",
      "type": "pubsub",
      "component": "msg"
    }
  },
  "outputs": {},
  "runtime": "OpenFuncAsync"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"subscriber","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"60011","inputs":{"sub":{"uri":"my_topic","type":"pubsub","component":"msg"}},"outputs":{},"runtime":"OpenFuncAsync"}'
```

Start the service and watch the logs.

```shell
cd subscriber/
go mod tidy
dapr run --app-id subscriber \
    --app-protocol grpc \
    --app-port 60011 \
    --components-path ../../components \
    go run ./main.go
```

## Producer

You also need a definition of producer.

```json
{
  "name": "producer",
  "version": "v1",
  "requestID": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "port": "60012",
  "inputs": {},
  "outputs": {
    "pub": {
      "uri": "my_topic",
      "component": "msg",
      "type": "pubsub"
    }
  },
  "runtime": "OpenFuncAsync"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"producer","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"60012","inputs":{},"outputs":{"msg":{"uri":"my_topic","component":"msg","type":"pubsub"}},"runtime":"OpenFuncAsync"}'
```

Start the service with another terminal to publish message.

```shell
cd producer/
go mod tidy
dapr run --app-id producer \
    --app-protocol grpc \
    --app-port 60012 \
    --components-path ../../components \
    go run ./main.go
```

## Results

<details>
<summary>View detailed producer logs.</summary>

```shell
== APP ==          0 published,   0/sec,   0 errors
== APP ==          0 published,   0/sec,   0 errors
== APP ==          0 published,   0/sec,   0 errors
== APP ==          1 published,   0/sec,   0 errors
== APP ==          4 published,   0/sec,   0 errors
== APP ==          7 published,   0/sec,   0 errors
== APP ==         10 published,   0/sec,   0 errors
== APP ==         13 published,   1/sec,   0 errors
```
</details>

<details>
<summary>View detailed subscriber logs.</summary>

```shell
== APP == 2021/06/28 10:04:18 event - Data: "{\"id\":\"p1-533d83d3-dd7c-4f1f-a822-f87b88f74d3e\",\"data\":\"QWdPTktLUjgxd1A1M096dVRDOHNWellyQjFoQ3FtM0FjeTY1Q2Q5S2NCVTRyMjhJbHlQcUVzdmxqWUJnZVB0YUlJRFRHWEFzWG5zZlQ3aGVRMUtrT21SalBHNzl4Rmx2bmNVSmNaOE11c3dmZ3plMk5ZRDF6Q0k5MmFFSVpuWUhmQ2J6aTlNSTQxajd1VURRNVJkMVNZYmhsUUs4UWRXN054Y3BDOXNHaDZTVEpZTzB5UFVJU2ZEQnZaZzJRYU5HaENDeFN6UzJPTVNYOU82QURxSnNndHB1dkIzcDVtRm1tT0haODJoMUM0UTl6blBjb3R0Qm8zbWRnQkhEbjE1Wg==\",\"sha\":\"f\\ufffd\\ufffd\\ufffdr\\ufffdE\\ufffdtu\\ufffd\\ufffd=n\\ufffd7\\u001d\\ufffd\\ufffdrB0^\\ufffd\\n1\\ufffd\\u0008\\ufffdWI\",\"time\":\"2021-06-28 10:04:13.867293349 +0800 CST m=+19.006097088\"}"
```
</details>

