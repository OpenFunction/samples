# Autoscaling service based on queue depth

## Prerequisites

Follow [this guide](../../../Prerequisites.md#openfunction) to install OpenFunction.

Follow [this guide](../../../Prerequisites.md#kafka) to install a Kafka server named `kafka-server` and a Topic named `sample-topic`.

Follow [this guide](../../../Prerequisites.md#registry-credential) to create a registry credential.

## Deployment

To configure the autoscaling demo we will deploy two functions: `subscriber` which will be used to process messages of the `sample-topic` queue in Kafka, and the `producer`, which will be publishing messages.

### Build and deploy Producer

Build and push the producer image.

```shell
cd producer
docker build -t <your registry name>/v1beta1-autoscaling-producer:latest -f Dockerfile.producer .
docker push <your registry name>/v1beta1-autoscaling-producer:latest
```

Modify the container image in `deploy.yaml`:

> You can set the `NUMBER_OF_PUBLISHERS` env and turn up its value appropriately so that the producer can trigger more subscribers in less time.

```yaml
    spec:
      containers:
        - name: producer
          image: <your registry name>/v1beta1-autoscaling-producer:latest
          imagePullPolicy: Always
          env:
            - name: PUBSUB_NAME
              value: "autoscaling-producer"
            - name: TOPIC_NAME
              value: "sample-topic"
          ports:
            - containerPort: 60034
              name: function-port
              protocol: TCP
```

Deploy the producer:

```shell
kubectl apply -f deploy.yaml
```

### Deploy Subscriber

Modify the ``spec.image`` field in `subscriber/function-subscribe.yaml` to your own container registry address:

```yaml
apiVersion: core.openfunction.io/v1beta1
kind: Function
metadata:
  name: autoscaling-subscriber
spec:
  image: "<your registry name>/autoscaling-subscriber:v1"
```

Use the following commands to create these Functions:

```shell
kubectl apply -f subscriber/function-subscriber.yaml
```

Back in the initial terminal now, some 20-30 seconds after the `producer` starts, you should see the number of `subscriber` pods being adjusted by Keda based on the number of the `sample-topic` topic.

