## Prerequisites

Knative runtime based functions can also interact with middleware via dapr components just like the Async functions.

In this case, we will create two functions: `function-front` and `output-target`

The relationship between these functions is shown in the following figure:

![](../../../images/knative-dapr.png)

Follow [this guide](../../../Prerequisites.md#openfunction) to install OpenFunction.

Follow [this guide](../../../Prerequisites.md#kafka) to install a Kafka server named `kafka-server` and a Topic named `sample-topic`.

Follow [this guide](../../../Prerequisites.md#registry-credential) to create a registry credential.

## Deployment

### output-target

`output-target` defines an input source (`serving.inputs`). This input source points to a dapr component of the Kafka server. This means that the `output-target` will be driven by events in the "sample-topic" topic of the Kafka server.

```yaml
serving:
  runtime: Async
  inputs:
    - name: sample
      component: target-topic
      type: bindings
  bindings:
    target-topic:
      type: bindings.kafka
      version: v1
      metadata:
        - name: brokers
          value: "kafka-server-kafka-brokers:9092"
        - name: topics
          value: "sample-topic"
        - name: consumerGroup
          value: "output-target"
        - name: publishTopic
          value: "sample-topic"
        - name: authRequired
          value: "false"
  template:
    containers:
      - name: function
        imagePullPolicy: Always
```

使用以下命令创建函数：

```shell
kubectl apply -f output/output-target.yaml
```

### function-front

> The function plugin mechanism is also demonstrated in `function-front`, and you can observe the following configuration in the definition of `function-front`.
>
> ```yaml
> apiVersion: core.openfunction.io/v1beta1
> kind: Function
> metadata:
>   name: function-front
>   annotations:
>     plugins: |
>       pre:
>       - plugin-custom
>       - plugin-example
>       post:
>       - plugin-custom
>       - plugin-example
> ```
>
> `pre` defines the order of plugins that need to be called before the user function is executed
>
> `post` defines the order of plugins that need to be called after the user function is executed
>
> You can learn about the logic of these two plugins and the effect of the plugins after they are executed here: [plugins mechanism](../../../functions-framework/README.md#plugin-mechanism)
>

In `function-front`, we define an output (`serving.outputs`) that will point to a dapr component definition of the Kafka server.

```yaml
  serving:
    runtime: Knative
    outputs:
      - name: target
        component: kafka-server
        type: bindings
        topic: "sample-topic"
        operation: "create"
    pubsub:
      kafka-server:
        type: bindings.kafka
        version: v1
        metadata:
          - name: brokers
            value: "kafka-server-kafka-brokers:9092"
          - name: authRequired
            value: "false"
          - name: publishTopic
            value: "sample-topic"
          - name: topics
            value: "sample-topic"
          - name: consumerGroup
            value: "function-front"
```

This allows us to send custom content to the output named `target` in the `function-front` function.

```go
func Sender(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
  ...
	_, err := ctx.Send("target", greeting)
	...
}

```

Use the following command to create `function-front` function:

```shell
kubectl apply -f function-front.yaml
```

## Demo

Check the current function status:

```shell
~# kubectl get functions.core.openfunction.io

NAME             BUILDSTATE   SERVINGSTATE   BUILDER         SERVING         URL                                             AGE
function-front   Succeeded    Running        builder-bhbtk   serving-vc6jw   http://openfunction.io/default/function-front   2m41s
output-target    Succeeded    Running        builder-dprfd   serving-75vrt                                                   2m21s
```

The `URL` is the address provided by the OpenFunction Domain that can be accessed. To access the function via this URL address, you need to make sure that DNS can resolve this address.

> You can use the following command to create a pod in the cluster and access the function from the pod:
>
> ```shell
> kubectl run curl --image=radial/busyboxplus:curl -i --tty --rm
> ```

Access the function via `URL`:

```shell
[ root@curl:/ ]$ curl -d '{"message":"Awesome OpenFunction!"}' -H "Content-Type: application/json" -X POST http://openfunction.io.svc.cluster.local/default/function-front
```

Query `function-front`'s log:

```shell
kubectl logs -f \
  $(kubectl get po -l \
  openfunction.io/serving=$(kubectl get functions function-front -o jsonpath='{.status.serving.resourceRef}') \
  -o jsonpath='{.items[0].metadata.name}') \
  function
```

The logs are as follows:

```shell
dapr client initializing for: 127.0.0.1:50001
I0125 06:51:55.584973       1 framework.go:107] Plugins for pre-hook stage:
I0125 06:51:55.585044       1 framework.go:110] - plugin-custom
I0125 06:51:55.585052       1 framework.go:110] - plugin-example
I0125 06:51:55.585057       1 framework.go:115] Plugins for post-hook stage:
I0125 06:51:55.585062       1 framework.go:118] - plugin-custom
I0125 06:51:55.585067       1 framework.go:118] - plugin-example
I0125 06:51:55.585179       1 knative.go:46] Knative Function serving http: listening on port 8080
2022/01/25 06:52:02 http - Data: {"message":"Awesome OpenFunction!"}
I0125 06:52:02.246450       1 plugin-example.go:83] the sum is: 2
```

Query `output-target`'s log:

```shell
kubectl logs -f \
  $(kubectl get po -l \
  openfunction.io/serving=$(kubectl get functions output-target -o jsonpath='{.status.serving.resourceRef}') \
  -o jsonpath='{.items[0].metadata.name}') \
  function
```

The logs are as follows:

```shell
dapr client initializing for: 127.0.0.1:50001
I0125 06:35:28.332381       1 framework.go:107] Plugins for pre-hook stage:
I0125 06:35:28.332863       1 framework.go:115] Plugins for post-hook stage:
I0125 06:35:28.333749       1 async.go:39] Async Function serving grpc: listening on port 8080
message from Kafka '{Awesome OpenFunction!}'
```

