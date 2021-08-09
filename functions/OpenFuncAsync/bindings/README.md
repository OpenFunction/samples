# Function Input/Output

## Prerequisites

### OpenFunction

You can refer to the [Installation Guide](https://github.com/OpenFunction/OpenFunction#readme) to set up OpenFunction.

### Kafka

If you don't have access to Kafka you can use these instructions to install Kafka into the cluster:

```shell
helm repo add confluentinc https://confluentinc.github.io/cp-helm-charts/
helm repo update
kubectl create ns kafka
helm install kafka confluentinc/cp-helm-charts -n kafka \
    --set cp-schema-registry.enabled=false \
    --set cp-kafka-rest.enabled=false \
    --set cp-kafka-connect.enabled=false \
    --set dataLogDirStorageClass=default \
    --set dataDirStorageClass=default \
    --set storageClass=default
kubectl rollout status deployment.apps/kafka-cp-control-center -n kafka
kubectl rollout status deployment.apps/kafka-cp-ksql-server -n kafka
kubectl rollout status statefulset.apps/kafka-cp-kafka -n kafka
kubectl rollout status statefulset.apps/kafka-cp-zookeeper -n kafka
```

When done, also deploy Kafka client and wait until it's ready:

```shell
cat <<EOF | kubectl apply -n kafka -f -
apiVersion: v1
kind: Pod
metadata:
 name: kafka-client
spec:
 containers:
  - name: kafka-client
    image: confluentinc/cp-enterprise-kafka:5.5.0
    command:
     - sh
     - -c
     - "exec tail -f /dev/null"
EOF

kubectl wait -n kafka --for=condition=ready pod kafka-client --timeout=120s
```

## Deployment

Create the `sample` topic which we will use in this sample:

```shell
kubectl -n kafka exec -it kafka-client -- kafka-topics \
    --zookeeper kafka-cp-zookeeper-headless:2181 \
    --topic sample \
    --create \
    --partitions 10 \
    --replication-factor 3 \
    --if-not-exists
```

Generate a secret to access your container registry, such as one on [Docker Hub](https://hub.docker.com/) or [Quay.io](https://quay.io/).
You can create this secret by editing the ``REGISTRY_SERVER``, ``REGISTRY_USER`` and ``REGISTRY_PASSWORD`` fields in following command, and then run it.

  ```bash
  REGISTRY_SERVER=https://index.docker.io/v1/ REGISTRY_USER=<your_registry_user> REGISTRY_PASSWORD=<your_registry_password>
  kubectl create secret docker-registry push-secret \
      --docker-server=$REGISTRY_SERVER \
      --docker-username=$REGISTRY_USER \
      --docker-password=$REGISTRY_PASSWORD
  ```

### Input only sample

We will use the sample in the `without-output` directory, which will be triggered by Dapr's `bindings.cron` component at a frequency of once every 2 seconds.

Modify the `spec.image` field in `without-output/function-bindings.yaml` to your own container registry address:

```yaml
apiVersion: core.openfunction.io/v1alpha1
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
apiVersion: core.openfunction.io/v1alpha1
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

