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
  "port": "50003",
  "inputs": {
    "sub": {
      "uri": "sample-topic",
      "componentName": "msg",
      "componentType": "pubsub.kafka"
    }
  },
  "outputs": {},
  "runtime": "Async"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"subscriber","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"50003","inputs":{"sub":{"uri":"sample-topic","componentName":"msg","componentType":"pubsub.kafka"}},"outputs":{},"runtime":"Async"}'
export CONTEXT_MODE='self-host'
```

### Run

Start the service and watch the logs.

```shell
cd sub/
go mod tidy
dapr run --app-id subscriber \
    --app-protocol grpc \
    --app-port 50003 \
    --dapr-grpc-port 50001 \
    --components-path ../../components \
    go run ./main.go
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
  "inputs": {
    "cron": {
      "componentName": "cron_input",
      "componentType": "bindings.cron"
    }
  },
  "outputs": {
    "pub": {
      "uri": "sample-topic",
      "componentName": "msg",
      "componentType": "pubsub.kafka"
    }
  },
  "runtime": "Async"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"producer","version":"v1","requestID":"a0f2ad8d-5062-4812-91e9-95416489fb01","port":"50004","inputs":{"cron":{"componentName":"cron_input","componentType":"bindings.cron"}},"outputs":{"pub":{"uri":"sample-topic","componentName":"msg","componentType":"pubsub.kafka"}},"runtime":"Async"}'
export CONTEXT_MODE='self-host'
export DAPR_GRPC_PORT="50002"
```

### Run

Start the service with another terminal to publish message.

```shell
cd pub/
go mod tidy
dapr run --app-id producer \
    --app-protocol grpc \
    --app-port 50004 \
    --dapr-grpc-port 50002 \
    --components-path ../../components \
    go run ./main.go
```

## Results

<details>
<summary>View detailed producer logs.</summary>

```shell
== APP == dapr client initializing for: 127.0.0.1:50002
== APP == I0308 12:00:57.853501 2463688 framework.go:110] Plugins for pre-hook stage:
== APP == I0308 12:00:57.853575 2463688 framework.go:118] Plugins for post-hook stage:
== APP == I0308 12:00:57.855762 2463688 async.go:111] registered bindings handler: cron_input
== APP == I0308 12:00:57.855780 2463688 async.go:53] Async Function serving grpc: listening on port 50004
INFO[0001] application discovered on port 50004          app_id=producer instance=crab scope=dapr.runtime type=log ver=1.5.1
INFO[0001] actor runtime started. actor idle timeout: 1h0m0s. actor scan interval: 30s  app_id=producer instance=crab scope=dapr.runtime.actor type=log ver=1.5.1
INFO[0001] dapr initialized. Status: Running. Init Elapsed 1666.519766ms  app_id=producer instance=crab scope=dapr.runtime type=log ver=1.5.1
INFO[0001] placement tables updated, version: 0          app_id=producer instance=crab scope=dapr.runtime.actor.internal.placement type=log ver=1.5.1
== APP == I0308 12:00:59.008107 2463688 pub.go:22] send msg and receive result:
== APP == I0308 12:01:01.003672 2463688 pub.go:22] send msg and receive result:
== APP == I0308 12:01:03.010105 2463688 pub.go:22] send msg and receive result:
== APP == I0308 12:01:05.028111 2463688 pub.go:22] send msg and receive result:
```
</details>

<details>
<summary>View detailed subscriber logs.</summary>

```shell
== APP == I0308 11:59:54.548833 2459950 framework.go:110] Plugins for pre-hook stage:
== APP == I0308 11:59:54.549080 2459950 framework.go:118] Plugins for post-hook stage:
== APP == dapr client initializing for: 127.0.0.1:50001
== APP == I0308 11:59:54.555141 2459950 async.go:143] registered pubsub handler: msg, topic: sample-topic
== APP == I0308 11:59:54.555166 2459950 async.go:53] Async Function serving grpc: listening on port 50003
INFO[0006] application discovered on port 50003          app_id=subscriber instance=crab scope=dapr.runtime type=log ver=1.5.1
INFO[0006] actor runtime started. actor idle timeout: 1h0m0s. actor scan interval: 30s  app_id=subscriber instance=crab scope=dapr.runtime.actor type=log ver=1.5.1
INFO[0006] app is subscribed to the following topics: [sample-topic] through pubsub=msg  app_id=subscriber instance=crab scope=dapr.runtime type=log ver=1.5.1
INFO[0006] placement tables updated, version: 0          app_id=subscriber instance=crab scope=dapr.runtime.actor.internal.placement type=log ver=1.5.1
INFO[0009] dapr initialized. Status: Running. Init Elapsed 9673.663393ms  app_id=subscriber instance=crab scope=dapr.runtime type=log ver=1.5.1
== APP == 2022/03/08 12:00:59 event - Data: {"hello":"world"}
== APP == 2022/03/08 12:00:59 event - Data: map[hello:world]
== APP == 2022/03/08 12:01:01 event - Data: {"hello":"world"}
== APP == 2022/03/08 12:01:01 event - Data: map[hello:world]
== APP == 2022/03/08 12:01:03 event - Data: {"hello":"world"}
== APP == 2022/03/08 12:01:03 event - Data: map[hello:world]
== APP == 2022/03/08 12:01:05 event - Data: {"hello":"world"}
== APP == 2022/03/08 12:01:05 event - Data: map[hello:world]
```
</details>

