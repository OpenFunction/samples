# Bindings

- [Bindings without output](#bindings-without-output)
- [Bindings with output](#bindings-with-output)
- [Results](#results)
  + [Without Output](#without-output)
  + [With Output](#with-output)

## Bindings without output

This input source will be executed every 2s (Refer to [cron.yaml](../config/cron.yaml)).

Prepare a context as follows, name it `function.json`. (You can refer to [OpenFunction Context Specs](https://github.com/OpenFunction/functions-framework/blob/main/docs/OpenFunction-context-specs.md) to learn more about the OpenFunction Context)

```json
{
  "name": "bindings",
  "version": "v1",
  "requestID": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "port": "50001",
  "inputs": {
    "cron": {
      "uri": "cron_input",
      "type": "bindings",
      "component": "cron_input"
    }
  },
  "outputs": {},
  "runtime": "OpenFuncAsync"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"bindings","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"50001","inputs":{"cron":{"uri":"cron_input","type":"bindings","component":"cron_input"}},"outputs":{},"runtime":"OpenFuncAsync"}'
```

Start the service and watch the logs.

```shell
cd without-output/
go mod tidy
dapr run --app-id bindings_grpc \
    --app-protocol grpc \
    --app-port 50001 \
    --components-path ../../components \
    go run ./main.go
```

## Bindings with output

We need to prepare an output target first.

```shell
cd with-output/
go mod tidy
dapr run --app-id output \
    --app-protocol http \
    --app-port 7489 \
    --dapr-http-port 7490 \
    go run ./output/main.go
```

This will generate two available targets, one for access through Dapr's proxy address and another for direct access through the app serving address.

> Simple test with execution `curl -X POST -H "ContentType: application/json" -d '{"Hello": "World"}' <urlPath>`
>
> `urlPath` refer to follows.

```
via Dapr: http://localhost:7490/v1.0/invoke/output_demo/method/echo
via App: http://localhost:7489/echo
```

In this example, the proxy address of Dapr will be used as the target of output.

>Here we have defined only one output, which will be called `item` in the following
>
>`app-id` is "echo" derived from the key of `item`
>
>Dapr component params are in `item.params`. Refer to [Dapr components reference](https://docs.dapr.io/reference/components-reference/).

```json
{
  "name": "bindings",
  "version": "v1",
  "requestID": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "port": "50001",
  "inputs": {
    "cron": {
      "uri": "cron_input",
      "type": "bindings",
      "component": "cron_input"
    }
  },
  "outputs": {
    "echo": {
      "uri": "echo",
      "operation": "create",
      "component": "echo",
      "metadata": {
        "path": "echo",
        "Content-Type": "application/json; charset=utf-8"
      },
      "type": "bindings"
    }
  },
  "runtime": "OpenFuncAsync"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"bindings","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"50001","inputs":{"cron":{"uri":"cron_input","type":"bindings","component":"cron_input"}},"outputs":{"echo":{"uri":"echo","operation":"create","component":"echo","metadata":{"path":"echo","Content-Type":"application/json; charset=utf-8"},"type":"bindings"}},"runtime":"OpenFuncAsync"}'
```

Start the service and watch the logs.

```shell
cd with-output/
dapr run --app-id bindings_grpc \
    --app-protocol grpc \
    --app-port 50001 \
    --components-path ../../components \
    go run ./main.go
```


## Results

### Without Output

The logs of user function is ...

<details>
<summary>View detailed logs.</summary>

```shell
== APP == 2021/06/28 10:43:58 binding - Data: Received
== APP == 2021/06/28 10:44:00 binding - Data: Received
```

</details>

### With Output

The logs of user function is ...

<details>
<summary>View detailed logs.</summary>

```shell
== APP == 2021/06/28 10:39:43 binding - Data: Received
== APP == 2021/06/28 10:39:45 binding - Data: Received
```

</details>

And the logs of output target app is ...

<details>
<summary>View detailed logs.</summary>

```shell
== APP == 2021/06/28 10:39:45 Receive a message:
== APP == 2021/06/28 10:39:45 Hello
```

</details>
