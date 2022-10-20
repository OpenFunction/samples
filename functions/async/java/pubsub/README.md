# Autoscaling service based on queue depth

## Prerequisites

You can refer to the [Installation Guide](https://openfunction.dev/docs/getting-started/installation/) to setup OpenFunction.

Follow [this guide](https://openfunction.dev/docs/getting-started/quickstarts/prerequisites/#kafka) to install a Kafka server named `kafka-server` and a Topic named `sample-topic`.

Follow [this guide](https://openfunction.dev/docs/getting-started/quickstarts/prerequisites/#registry-credential) to create a registry credential.

## Deployment

To configure the autoscaling demo we will deploy two functions: `subscriber` which will be used to process messages of the `sample-topic` queue in Kafka, and the `producer`, which will be publishing messages.

### Build and deploy Producer

Follow [this guide](../../pubsub/README.md#Build-and-deploy-Producer) to deploy a producer.

### Deploy Subscriber

Modify the ``spec.image`` field in `function-subscribe.yaml` to your own container registry address:

```yaml
apiVersion: core.openfunction.io/v1beta1
kind: Function
metadata:
  name: autoscaling-subscriber-java
spec:
  image: "<your registry name>/autoscaling-subscriber-java:v1"
```

Use the following commands to create these Functions:

```shell
kubectl apply -f function-subscriber.yaml
```

Back in the initial terminal now, some 20-30 seconds after the `producer` starts, you should see the number of `subscriber` pods being adjusted by Keda based on the number of the `sample-topic` topic.

