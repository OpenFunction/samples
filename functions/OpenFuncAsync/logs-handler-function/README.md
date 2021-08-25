# Function Logs Handler

## Prerequisites

### OpenFunction

You can refer to the [Installation Guide](https://github.com/OpenFunction/OpenFunction#readme) to set up OpenFunction.

### Kafka

First, we enable the **logging** components for the KubeSphere platform (see [Enabling Pluggable Components](https://kubesphere.io/docs/pluggable-components/) for more information ). We then build a minimalist Kafka server using [strimzi-kafka-operator](https://github.com/strimzi/strimzi-kafka-operator).

1. Install [strimzi-kafka-operator](https://github.com/strimzi/strimzi-kafka-operator) in the default namespace.

   ```shell
   helm repo add strimzi https://strimzi.io/charts/
   helm install kafka-operator -n default strimzi/strimzi-kafka-operator
   ```

2. Run the following command to create a Kafka cluster and Kafka Topic in the default namespace. The Kafka and Zookeeper clusters created by this command have a storage type of **ephemeral** and are demonstrated using emptyDir.

   > We have created a topic called "logs", which will be used later

   ```shell
   cat <<EOF | kubectl apply -f -
   apiVersion: kafka.strimzi.io/v1beta2
   kind: Kafka
   metadata:
     name: kafka-logs-receiver
     namespace: default
   spec:
     kafka:
       version: 2.8.0
       replicas: 1
       listeners:
         - name: plain
           port: 9092
           type: internal
           tls: false
         - name: tls
           port: 9093
           type: internal
           tls: true
       config:
         offsets.topic.replication.factor: 1
         transaction.state.log.replication.factor: 1
         transaction.state.log.min.isr: 1
         log.message.format.version: '2.8'
         inter.broker.protocol.version: "2.8"
       storage:
         type: ephemeral
     zookeeper:
       replicas: 1
       storage:
         type: ephemeral
     entityOperator:
       topicOperator: {}
       userOperator: {}
   ---
   apiVersion: kafka.strimzi.io/v1beta1
   kind: KafkaTopic
   metadata:
     name: logs
     namespace: default
     labels:
       strimzi.io/cluster: kafka-logs-receiver
   spec:
     partitions: 10
     replicas: 3
     config:
       retention.ms: 7200000
       segment.bytes: 1073741824
   EOF
   ```

3. Run the following command to check Pod status and wait for Kafka and Zookeeper to run and start.

   ```shell
   $ kubectl get po
   NAME                                                   READY   STATUS        RESTARTS   AGE
   kafka-logs-receiver-entity-operator-568957ff84-nmtlw   3/3     Running       0          8m42s
   kafka-logs-receiver-kafka-0                            1/1     Running       0          9m13s
   kafka-logs-receiver-zookeeper-0                        1/1     Running       0          9m46s
   strimzi-cluster-operator-687fdd6f77-cwmgm              1/1     Running       0          11m
   ```

   Run the following command to view the metadata for the Kafka cluster.

   ```shell
   $ kafkacat -L -b kafka-logs-receiver-kafka-brokers:9092
   ```

We add this Kafka server as a log receiver.

1. Log in to KubeSphere's Web Console as **admin**. Click on **Platform** in the top left corner and select **Cluster Management**.

2. On the **Cluster Management** page, select **Log Collection** under **Cluster Settings**.

3. Click on **Add Log Receiver** and select **Kafka**. Enter the Kafka proxy address and port information, then click **OK** to continue.

4. Run the following command to verify that the Kafka cluster is receiving logs from Fluent Bit.

   ```shell
   # Start a util container
   kubectl run utils --image=arunvelsriram/utils -i --tty --rm 
   # Run the following command to consume log messages from kafka topic: my-topic
   kafkacat -C -b kafka-logs-receiver-kafka-0.kafka-logs-receiver-kafka-brokers.default.svc:9092 -t logs
   ```

## Deployment

Generate a secret to access your container registry, such as one on [Docker Hub](https://hub.docker.com/) or [Quay.io](https://quay.io/). You can create this secret by editing the `REGISTRY_SERVER`, `REGISTRY_USER` and `REGISTRY_PASSWORD` fields in following command, and then run it.

```shell
REGISTRY_SERVER=https://index.docker.io/v1/ REGISTRY_USER=<your_registry_user> REGISTRY_PASSWORD=<your_registry_password>
kubectl create secret docker-registry push-secret \
    --docker-server=$REGISTRY_SERVER \
    --docker-username=$REGISTRY_USER \
    --docker-password=$REGISTRY_PASSWORD
```

We create this logs handler function:

```shell
kubectl apply -f logs-handler-function.yaml
```

The logs handler function is then driven by messages from the logs topic in Kafka.
