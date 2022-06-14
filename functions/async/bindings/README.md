# Function Input/Output

## Prerequisites

Follow [this guide](../../../Prerequisites.md#openfunction) to install OpenFunction.

Follow [this guide](../../../Prerequisites.md#kafka) to install a Kafka server named `kafka-server` and a Topic named `sample-topic`.

Follow [this guide](../../../Prerequisites.md#registry-credential) to create a registry credential.

## Deployment

### Input only sample

We will use the sample in the `cron-input` directory, which will be triggered by Dapr's `bindings.cron` component at a frequency of once every 2 seconds.

Modify the `spec.image` field in `cron-input/cron-input.yaml` to your own container registry address:

```yaml
apiVersion: core.openfunction.io/v1beta1
kind: Function
metadata:
  name: cron-input
spec:
  image: "<your registry name>/cron-input:v1"
```

Use the following commands to create this Function:

```shell
kubectl apply -f cron-input/cron-input.yaml
```

Afterwards, use the following command to observe the log of the function:

```shell
kubectl logs -f \
  $(kubectl get po -l \
  openfunction.io/serving=$(kubectl get functions cron-input -o jsonpath='{.status.serving.resourceRef}') \
  -o jsonpath='{.items[0].metadata.name}') \
  function
```

You will be able to see messages similar to the following:

```shell
dapr client initializing for: 127.0.0.1:50001
I0125 07:51:47.615204       1 framework.go:107] Plugins for pre-hook stage:
I0125 07:51:47.615266       1 framework.go:115] Plugins for post-hook stage:
I0125 07:51:47.615373       1 async.go:39] Async Function serving grpc: listening on port 8080
2022/01/25 07:51:50 binding - Data: Received
2022/01/25 07:51:52 binding - Data: Received
```

### Input and Output

We will use the sample in the `cron-input-kafka-output` directory, which will be triggered by Dapr's `bindings.cron` component at a frequency of once every 2 seconds. After being triggered, it will send a greeting to another service via Dapr's `bindings.kafka` component. 

Modify the `spec.image` field in `cron-input-kafka-output/cron-input-kafka-output.yaml` to your own container registry address:

```yaml
apiVersion: core.openfunction.io/v1beta1
kind: Function
metadata:
  name: cron-input-kafka-output
spec:
  image: "<your registry name>/cron-input-kafka-output:v1"
```

Use the following commands to create this Function:

```shell
kubectl apply -f cron-input-kafka-output/cron-input-kafka-output.yaml
```

Afterwards, use the following command to observe the log of the function:

```shell
kubectl logs -f \
  $(kubectl get po -l \
  openfunction.io/serving=$(kubectl get functions cron-input-kafka-output -o jsonpath='{.status.serving.resourceRef}') \
  -o jsonpath='{.items[0].metadata.name}') \
  function
```

You will be able to see messages similar to the following:

```shell
I0125 07:42:36.485929       1 framework.go:107] Plugins for pre-hook stage:
I0125 07:42:36.486023       1 framework.go:110] - plugin-custom
I0125 07:42:36.486036       1 framework.go:110] - plugin-example
I0125 07:42:36.486043       1 framework.go:115] Plugins for post-hook stage:
I0125 07:42:36.486057       1 framework.go:118] - plugin-custom
I0125 07:42:36.486065       1 framework.go:118] - plugin-example
dapr client initializing for: 127.0.0.1:50001
I0125 07:42:36.486285       1 async.go:39] Async Function serving grpc: listening on port 8080
2022/01/25 07:42:38 binding - Data: Received
I0125 07:42:38.008694       1 plugin-example.go:83] the sum is: 2
2022/01/25 07:42:40 binding - Data: Received
I0125 07:42:40.004346       1 plugin-example.go:83] the sum is: 2
```

Now we need to start the kafka-input function.

```shell
kubectl apply -f kafka-input/kafka-input.yaml
```

Use the following command to observe the log of the function:

```shell
kubectl logs -f \
  $(kubectl get po -l \
  openfunction.io/serving=$(kubectl get functions kafka-input -o jsonpath='{.status.serving.resourceRef}') \
  -o jsonpath='{.items[0].metadata.name}') \
  function
```

You will be able to see messages similar to the following:

```shell
I0125 07:41:43.086181       1 framework.go:107] Plugins for pre-hook stage:
I0125 07:41:43.086233       1 framework.go:115] Plugins for post-hook stage:
I0125 07:41:43.086322       1 async.go:39] Async Function serving grpc: listening on port 8080
message from Kafka '{Hello}'
message from Kafka '{Hello}'
```
