# OpenFunction function - Pubsub

- [Subscriber](#subscriber)
- [Producer](#producer)
- [Results](#results)

## Subscriber

### FUNC_CONTEXT

Prepare a context as follows, name it `function.json`. (You can refer to [OpenFunction Context Specs](https://github.com/OpenFunction/functions-framework/blob/main/docs/OpenFunction-context-specs.md) to learn more about the OpenFunction Context)

```json
{
  "name": "subscriber",
  "version": "v1",
  "requestID": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "port": "50002",
  "inputs": {
    "sub": {
      "uri": "my_topic",
      "type": "pubsub",
      "component": "msg"
    }
  },
  "outputs": {},
  "runtime": "Async",
  "prePlugins": ["plugin-custom", "plugin-example"],
  "postPlugins": ["plugin-custom", "plugin-example"]
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"subscriber","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"50003","inputs":{"sub":{"uri":"my_topic","type":"pubsub","component":"msg"}},"outputs":{},"runtime":"Async","prePlugins":["plugin-custom","plugin-example"],"postPlugins":["plugin-custom","plugin-example"]}'
```

### Run

Start the service and watch the logs.

```shell
cd subscriber/
go mod tidy
dapr run --app-id subscriber \
    --app-protocol grpc \
    --app-port 50003 \
    --dapr-grpc-port 50001 \
    --components-path ../../components \
    go run ./main.go ./plugin.go
```

## Producer

### FUNC_CONTEXT

You also need a definition of producer.

```json
{
  "name": "producer",
  "version": "v1",
  "requestID": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "port": "50004",
  "inputs": {},
  "outputs": {
    "pub": {
      "uri": "my_topic",
      "component": "msg",
      "type": "pubsub"
    }
  },
  "runtime": "Async",
  "prePlugins": ["plugin-custom", "plugin-example"],
  "postPlugins": ["plugin-custom", "plugin-example"]
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"producer","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"50004","inputs":{},"outputs":{"pub":{"uri":"my_topic","component":"msg","type":"pubsub"}},"runtime":"Async","prePlugins":["plugin-custom","plugin-example"],"postPlugins":["plugin-custom","plugin-example"]}'
export DAPR_GRPC_PORT="50002"
```

### Run

Start the service with another terminal to publish message.

```shell
cd producer/
go mod tidy
dapr run --app-id producer \
    --app-protocol grpc \
    --app-port 50004 \
    --dapr-grpc-port 50002 \
    --components-path ../../components \
    go run ./main.go ./plugin.go
```

## Results

<details>
<summary>View detailed producer logs.</summary>

```shell
== APP == I0111 09:54:53.957079  349657 framework.go:83] exec pre hooks: plugin-custom of version v1
== APP == I0111 09:54:53.957109  349657 framework.go:83] exec pre hooks: plugin-example of version v1
== APP == subscription name: pub
== APP == number of publishers: 1
== APP == publish frequency: 1s
== APP == log frequency: 3s
== APP == publish delay: 10s
== APP == I0111 09:54:53.957774  349657 framework.go:94] exec post hooks: plugin-custom of version v1
== APP == I0111 09:54:53.957801  349657 framework.go:94] exec post hooks: plugin-example of version v1
== APP == I0111 09:54:53.957926  349657 plugin-example.go:79] the sum is: 4
== APP ==          0 published,   0/sec,   0 errors
== APP ==          0 published,   0/sec,   0 errors
== APP ==          0 published,   0/sec,   0 errors
== APP ==          1 published,   0/sec,   0 errors
== APP ==          4 published,   0/sec,   0 errors
== APP ==          7 published,   0/sec,   0 errors
```
</details>

<details>
<summary>View detailed subscriber logs.</summary>

```shell
== APP == I0111 09:55:11.960645  348060 framework.go:83] exec pre hooks: plugin-custom of version v1
== APP == I0111 09:55:11.960664  348060 framework.go:83] exec pre hooks: plugin-example of version v1
== APP == 2022/01/11 09:55:11 event - Data: {"id":"p1-916f637c-7460-4db3-be3e-85893890ec7e","data":"MlpSUzV6cTRiTmszb1NLWmxhZHZKRjBlN095ZjBNWGJucWJyMGlabTE3Y0R2amR4dmZoWENTVXdkNWRMc0lQVUdBUkJ4aEFTSmNYZWtVSkkxRHdNZGhUOGplVXhnS3ZDTkJ5QlY2UkRQUzM3eWlIVWVkaVVYVld2SGZabXNFVzZMR1dkSENkeFVHUjBLaThpeTA1YlFDS1VYOWZlTjlpOVZXSlFNMTF4T3g1V3V6T1ZkMFVndEs3MGJHbEtSNTVQcE1lNXJEOUpyTkVCdFRIdEgzMUZwV21xVmlRYWxlazlBVDNOYjZZeUdLY1FlUm1xQ2Z3SnZtVml2dlRNYzJiZw==","sha":"B\ufffdV\t\ufffdW\u0001!F\u0005C\ufffd\ufffd\ufffd+\ufffd!\ufffd]\ufffd\ufffd\ufffdi\ufffde~\ufffd\ufffd\ufffd\ufffd\ufffdF","time":"2022-01-11 09:55:11.959063585 +0800 CST m=+18.007867886"}
== APP == I0111 09:55:11.960680  348060 framework.go:94] exec post hooks: plugin-custom of version v1
== APP == I0111 09:55:11.960686  348060 framework.go:94] exec post hooks: plugin-example of version v1
== APP == I0111 09:55:11.960705  348060 plugin-example.go:79] the sum is: 4
```
</details>

