# Autoscaling service based on queue depth

## Prerequisites

Follow [this guide](../../../../Prerequisites.md#openfunction) to install OpenFunction.

Follow [this guide](../../../../Prerequisites.md#kafka) to install a Kafka server named `kafka-pubsub-server` and a Topic named `metric`.

Follow [this guide](../../../../Prerequisites.md#registry-credential) to create a registry credential.Kafka

## Deployment

To configure the autoscaling demo we will deploy two functions: `subscriber` which will be used to process messages of the `metric` queue in Kafka, and the `producer`, which will be publishing messages.

Modify the ``spec.image`` field in ``producer/function-producer.yaml`` and ``subscriber/function-subscribe.yaml`` to your own container registry address:

    ```yaml
    apiVersion: core.openfunction.io/v1alpha1
    kind: Function
    metadata:
      name: autoscaling-producer
    spec:
      image: "<your registry name>/autoscaling-producer:v0.3.1"
    ```
    
    ```yaml
    apiVersion: core.openfunction.io/v1alpha1
    kind: Function
    metadata:
      name: autoscaling-subscriber
    spec:
      image: "<your registry name>/autoscaling-subscriber:v0.3.1"
    ```

Use the following commands to create these Functions:

```shell
kubectl apply -f producer/function-producer.yaml
kubectl apply -f subscriber/function-subscriber.yaml
```

Back in the initial terminal now, some 20-30 seconds after the `producer` starts, you should see the number of `subscriber` pods being adjusted by Keda based on the number of the `metric` topic.

