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
kubectl apply -n kafka -f deployment/kafka-client.yaml
cat <<EOF | kubectl apply -f -
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
		--replication-factor 3 \
		--if-not-exists
```

In order to access your container registry, you need to create a secret. You can create this secret by editing the ``username`` and ``password`` fields in following command, and then run it.

```shell
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: basic-user-pass
type: kubernetes.io/basic-auth
stringData:
  username: <USERNAME>
  password: <PASSWORD>
EOF
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
kubectl logs -f $(kubectl get po -l serving=bindings-without-output-serving -o jsonpath='{.items[0].metadata.name}') function
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

First we need to start the output target service. (Use `kubectl delete -f with-output/output/goapp.yaml` for cleaning.)

```shell
kubectl apply -f with-output/output/goapp.yaml
```

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
kubectl logs -f $(kubectl get po -l serving=bindings-with-output-serving -o jsonpath='{.items[0].metadata.name}') function
```

You will be able to see messages similar to the following:

```shell
2021/07/02 09:02:02 binding - Data: Received
2021/07/02 09:02:04 binding - Data: Received
2021/07/02 09:02:06 binding - Data: Received
2021/07/02 09:02:08 binding - Data: Received
```

Use command `kubectl logs -f $(kubectl get po -l app=bindingsgoapp -o jsonpath='{.items[0].metadata.name}') go` to observe the logs of output target service:

```shell
data from Event Hubs '{Hello}'
data from Event Hubs '{Hello}'
data from Event Hubs '{Hello}'
data from Event Hubs '{Hello}'
```

