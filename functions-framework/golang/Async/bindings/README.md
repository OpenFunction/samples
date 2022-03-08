# OpenFunction function - Bindings

- [Bindings without output](#bindings-without-output)
- [Bindings with output](#bindings-with-output)
- [Results](#results)
  + [Without Output](#without-output)
  + [With Output](#with-output)

## Bindings without output

### FUNC_CONTEXT

This input source will be executed every 2s (Refer to [cron.yaml](../config/cron.yaml)).

Prepare a context as follows, name it `function.json`. (You can refer to [OpenFunction Context Specs](https://github.com/OpenFunction/functions-framework/blob/main/docs/OpenFunction-context-specs.md) to learn more about the OpenFunction Context)

```json
{
  "name": "bindings",
  "version": "v1",
  "requestID": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "port": "50002",
  "inputs": {
    "cron": {
      "uri": "cron_input",
      "componentType": "bindings.cron",
      "componentName": "cron_input"
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
export FUNC_CONTEXT='{"name":"bindings","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"50002","inputs":{"cron":{"uri":"cron_input","componentType":"bindings.cron","componentName":"cron_input"}},"outputs":{},"runtime":"Async","prePlugins":["plugin-custom","plugin-example"],"postPlugins":["plugin-custom","plugin-example"]}'
export CONTEXT_MODE='self-host'
```

### Run

Start the service and watch the logs.

```shell
cd without-output/
go mod tidy
dapr run --app-id bindings_grpc \
    --app-protocol grpc \
    --app-port 50002 \
    --dapr-grpc-port 50001 \
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

### FUNC_CONTEXT

In this example, the proxy address of Dapr will be used as the target of output.

>We define the output target in `$.outputs`. This is a map:
>
>- The key is used to hold the `app-id` in dapr, in this case "echo"
>- The value is used to store the dapr component spec
>
>Refer to [Dapr components reference](https://docs.dapr.io/reference/components-reference/).

```json
{
  "name": "bindings",
  "version": "v1",
  "requestID": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "port": "50002",
  "inputs": {
    "cron": {
      "uri": "cron_input",
      "componentType": "bindings.cron",
      "componentName": "cron_input"
    }
  },
  "outputs": {
    "echo": {
      "uri": "echo",
      "operation": "create",
      "componentName": "echo",
      "metadata": {
        "path": "echo",
        "Content-Type": "application/json; charset=utf-8"
      },
      "componentType": "bindings.http"
    }
  },
  "runtime": "Async",
  "prePlugins": ["plugin-custom", "plugin-example"],
  "postPlugins": ["plugin-custom", "plugin-example"]
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"bindings","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"50002","inputs":{"cron":{"uri":"cron_input","componentType":"bindings.cron","componentName":"cron_input"}},"outputs":{"echo":{"uri":"echo","operation":"create","componentName":"echo","metadata":{"path":"echo","Content-Type":"application/json; charset=utf-8"},"componentType":"bindings.http"}},"runtime":"Async","prePlugins":["plugin-custom","plugin-example"],"postPlugins":["plugin-custom","plugin-example"]}'
export CONTEXT_MODE='self-host'
```

### Run

Start the service and watch the logs.

```shell
cd with-output/
dapr run --app-id bindings_grpc \
    --app-protocol grpc \
    --app-port 50002 \
    --dapr-grpc-port 50001 \
    --components-path ../../components \
    go run ./main.go
```


## Results

### Without Output

The logs of user function is ...

<details>
<summary>View detailed logs.</summary>

```shell
== APP == I0111 09:44:05.001313  329924 framework.go:83] exec pre hooks: plugin-custom of version v1
== APP == I0111 09:44:05.001334  329924 framework.go:83] exec pre hooks: plugin-example of version v1
== APP == 2022/01/11 09:44:05 binding - Data: Received
== APP == I0111 09:44:05.001343  329924 framework.go:94] exec post hooks: plugin-custom of version v1
== APP == I0111 09:44:05.001347  329924 framework.go:94] exec post hooks: plugin-example of version v1
== APP == I0111 09:44:05.001364  329924 plugin-example.go:79] the sum is: 4
```

</details>

### With Output

The logs of user function is ...

<details>
<summary>View detailed logs.</summary>

```shell
== APP == I0110 11:00:16.001426 2140656 framework.go:83] exec pre hooks: plugin-custom of version v1
== APP == I0110 11:00:16.001456 2140656 framework.go:83] exec pre hooks: plugin-example of version v1
== APP == 2022/01/10 11:00:16 binding - Data: Received
== APP == I0110 11:00:16.019345 2140656 framework.go:94] exec post hooks: plugin-custom of version v1
== APP == I0110 11:00:16.019380 2140656 framework.go:94] exec post hooks: plugin-example of version v1
== APP == I0110 11:00:16.019404 2140656 plugin-example.go:79] the sum is: 4
```

</details>

And the logs of output target app is ...

<details>
<summary>View detailed logs.</summary>

```shell
== APP == 2022/01/10 11:00:16 Receive a message:
== APP == 2022/01/10 11:00:16 Hello
```

</details>
