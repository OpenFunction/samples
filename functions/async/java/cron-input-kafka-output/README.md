# Autoscaling service based on queue depth

## Prerequisites

You can refer to the [Installation Guide](https://openfunction.dev/docs/getting-started/installation/) to setup OpenFunction.

Follow [this guide](https://openfunction.dev/docs/getting-started/quickstarts/prerequisites/#kafka) to install a Kafka server named `kafka-server` and a Topic named `sample-topic`.

Follow [this guide](https://openfunction.dev/docs/getting-started/quickstarts/prerequisites/#registry-credential) to create a registry credential.

## Deployment

The `cron-input-kafka-output-java` function will be triggered by Dapr's `bindings.cron` component at a frequency of once every 2 seconds. After being triggered, it will send a greeting to another service via Dapr's `bindings.kafka` component.

Modify the `spec.image` field in `cron-input-kafka-output.yaml` to your own container registry address:

  ```yaml
  apiVersion: core.openfunction.io/v1beta1
  kind: Function
  metadata:
    name: cron-input-kafka-output-java
  spec:
    image: "<your registry name>/cron-input-kafka-output-java:v1"
  ```

Use the following commands to create this Function:

  ```shell
  kubectl apply -f cron-input-kafka-output.yaml
  ```

Afterwards, use the following command to observe the log of the function:

  ```shell
  kubectl logs -f \
    $(kubectl get po -l \
    openfunction.io/serving=$(kubectl get functions cron-input-kafka-output-java -o jsonpath='{.status.serving.resourceRef}') \
    -o jsonpath='{.items[0].metadata.name}') \
    function
  ```

You will be able to see messages similar to the following:

  ```shell
  plugin plugin-example:v1.0.0 exec pre hook for binding serving-2w6ft-component-cron-jsvxg at 2022-10-19 08:00:25.Z
  receive event: 
  send to output sample
  plugin plugin-example:v1.0.0 exec post hook for binding serving-2w6ft-component-cron-jsvxg at 2022-10-19 08:00:25.Z
  ```

Now we need to start the kafka-input function.

  ```shell
  kubectl apply -f ../kafka-input/kafka-input.yaml
  ```

Use the following command to observe the log of the function:
  
  ```shell
  kubectl logs -f \
    $(kubectl get po -l \
    openfunction.io/serving=$(kubectl get functions kafka-input-java -o jsonpath='{.status.serving.resourceRef}') \
    -o jsonpath='{.items[0].metadata.name}') \
    function
  ```

You will be able to see messages similar to the following:

  ```shell
 plugin plugin-example:v1.0.0 exec pre hook for binding serving-p7nll-component-target-topic-kjh7q at 2022-10-19 08:12:58.Z
 receive event: ""
 plugin plugin-example:v1.0.0 exec post hook for binding serving-p7nll-component-target-topic-kjh7q at 2022-10-19 08:12:58.Z
  ```
