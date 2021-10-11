# Function Input/Output

## Prerequisites

Follow [this guide](../../../../Prerequisites.md#openfunction) to install OpenFunction.

Follow [this guide](../../../../Prerequisites.md#kafka) to install a Kafka server named `kafka-bindings-server` and a Topic named `sample`.

Follow [this guide](../../../../Prerequisites.md#registry-credential) to create a registry credential.

## Deployment

### Input only sample

We will use the sample in the `without-output` directory, which will be triggered by Dapr's `bindings.cron` component at a frequency of once every 2 seconds.

Modify the `spec.image` field in `without-output/function-bindings.yaml` to your own container registry address:

```yaml
apiVersion: core.openfunction.io/v1alpha2
kind: Function
metadata:
  name: bindings-without-output
spec:
  image: "<your registry name>/bindings-without-output:latest"
```

Use the following commands to create this Function:

```shell
kubectl apply -f without-output/function-bindings.yaml
```

Afterwards, use the following command to observe the log of the function:

```shell
kubectl logs -f \
  $(kubectl get po -l \
  openfunction.io/serving=$(kubectl get functions bindings-without-output -o jsonpath='{.status.serving.resourceRef}') \
  -o jsonpath='{.items[0].metadata.name}') \
  function
```

You will be able to see messages similar to the following:

```shell
2021/07/02 09:02:02 binding - Data: Received
2021/07/02 09:02:04 binding - Data: Received
2021/07/02 09:02:06 binding - Data: Received
2021/07/02 09:02:08 binding - Data: Received
```

### Input and Output

We will use the sample in the `with-output` directory, which will be triggered by Dapr's `bindings.cron` component at a frequency of once every 2 seconds. After being triggered, it will send a greeting to another service via Dapr's `bindings.kafka` component. 

Modify the `spec.image` field in `without-output/function-bindings.yaml` to your own container registry address:

```yaml
apiVersion: core.openfunction.io/v1alpha2
kind: Function
metadata:
  name: bindings-with-output
spec:
  image: "<your registry name>/bindings-with-output:latest"
```

Use the following commands to create this Function:

```shell
kubectl apply -f with-output/function-bindings.yaml
```

Afterwards, use the following command to observe the log of the function:

```shell
kubectl logs -f \
  $(kubectl get po -l \
  openfunction.io/serving=$(kubectl get functions bindings-with-output -o jsonpath='{.status.serving.resourceRef}') \
  -o jsonpath='{.items[0].metadata.name}') \
  function
```

You will be able to see messages similar to the following:

```shell
2021/07/02 09:02:02 binding - Data: Received
2021/07/02 09:02:04 binding - Data: Received
2021/07/02 09:02:06 binding - Data: Received
2021/07/02 09:02:08 binding - Data: Received
```

Now we need to start the output target service. (Use `kubectl delete -f with-output/output/goapp.yaml` for cleaning.)

```shell
kubectl apply -f with-output/output/goapp.yaml
```

Use command `kubectl logs -f $(kubectl get po -l app=bindingsgoapp -o jsonpath='{.items[0].metadata.name}') go` to observe the logs of output target service:

```shell
data from Event Hubs '{Hello}'
data from Event Hubs '{Hello}'
data from Event Hubs '{Hello}'
data from Event Hubs '{Hello}'
```

