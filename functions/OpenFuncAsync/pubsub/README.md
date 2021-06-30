# Autoscaling service based on queue depth

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

Create the `metric` topic which we will use in this sample:

> The number of `partitions` is connected to the maximum number of replicas can be scaled.

```shell
kubectl -n kafka exec -it kafka-client -- kafka-topics \
		--zookeeper kafka-cp-zookeeper-headless:2181 \
		--topic metric \
		--create \
		--partitions 10 \
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

To configure the autoscaling demo we will deploy two functions: `subscriber` which will be used to process messages of the `metric` queue in Kafka, and the `producer`, which will be publishing messages.

Modify the ``spec.image`` field in ``producer/function-producer.yaml`` and ``subscriber/function-subscribe.yaml`` to your own container registry address:

    ```yaml
    apiVersion: core.openfunction.io/v1alpha1
    kind: Function
    metadata:
      name: autoscaling-producer
    spec:
      image: "<your registry name>/autoscaling-producer:latest"
    ```

    ```yaml
    apiVersion: core.openfunction.io/v1alpha1
    kind: Function
    metadata:
      name: autoscaling-subscriber
    spec:
      image: "<your registry name>/autoscaling-subscriber:latest"
      ```

Use the following commands to create these Functions:

```shell
kubectl apply -f producer/function-producer-sample.yaml
kubectl apply -f subscriber/function-subscriber-sample.yaml
```

Back in the initial terminal now, some 20-30 seconds after the `producer` starts, you should see the number of `subscriber` pods being adjusted by Keda based on the number of the `metric` topic:


