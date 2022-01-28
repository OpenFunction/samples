# Autoscaling service based on queue depth

## Prerequisites

Follow [this guide](../../../Prerequisites.md#openfunction) to install OpenFunction.

Follow [this guide](../../../Prerequisites.md#kafka) to install a Kafka server named `kafka-server` and a Topic named `sample-topic`.

Follow [this guide](../../../Prerequisites.md#registry-credential) to create a registry credential.

## Deployment

To configure the autoscaling demo we will deploy two functions: `subscriber` which will be used to process messages of the `sample-topic` queue in Kafka, and the `producer`, which will be publishing messages.

Modify the ``spec.image`` field in ``producer/function-producer.yaml`` and ``subscriber/function-subscribe.yaml`` to your own container registry address:

    ```yaml
    apiVersion: core.openfunction.io/v1beta1
    kind: Function
    metadata:
      name: autoscaling-producer
    spec:
      image: "<your registry name>/autoscaling-producer:latest"
    ```
    
    ```yaml
    apiVersion: core.openfunction.io/v1beta1
    kind: Function
    metadata:
      name: autoscaling-subscriber
    spec:
      image: "<your registry name>/autoscaling-subscriber:latest"
    ```

Use the following commands to create these Functions:

```shell
kubectl apply -f producer/function-producer.yaml
kubectl apply -f subscriber/function-subscriber.yaml
```

Back in the initial terminal now, some 20-30 seconds after the `producer` starts, you should see the number of `subscriber` pods being adjusted by Keda based on the number of the `sample-topic` topic.

