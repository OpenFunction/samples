## Prerequisites

Knative runtime based functions can also interact with middleware via dapr components just like the Async functions.

In this case, we will create two functions: `function-front` and `kafka-input`. `function-front` will send the content 
to the kafka after a HTTP request received. The `kafka-input` will watch the kafka and consume it.

You can refer to the [Installation Guide](https://openfunction.dev/docs/getting-started/installation/) to setup OpenFunction.

Follow [this guide](https://openfunction.dev/docs/getting-started/quickstarts/prerequisites/#kafka) to install a Kafka server named `kafka-server` and a Topic named `sample-topic`.

Follow [this guide](https://openfunction.dev/docs/getting-started/quickstarts/prerequisites/#registry-credential) to create a registry credential.

## Deployment

### kafka-input

[kafka-input](kafka-input.yaml) defines an input source (`serving.inputs`). This input source points to a dapr component of the Kafka server. This means that the `kafka-input` will be driven by events in the "sample-topic" topic of the Kafka server.

  ```yaml
  serving:
    runtime: async
    inputs:
      - name: sample
        component: target-topic
        type: bindings
    bindings:
      target-topic:
        type: bindings.kafka
        version: v1
        metadata:
          - name: brokers
            value: "kafka-server-kafka-brokers:9092"
          - name: topics
            value: "sample-topic"
          - name: consumerGroup
            value: "kafka-input"
          - name: publishTopic
            value: "sample-topic"
          - name: authRequired
            value: "false"
    template:
      containers:
        - name: function # DO NOT change this
          imagePullPolicy: IfNotPresent 
  ```

Use the following command to create the function:

  ```shell
  kubectl apply -f ../../../async/java/kafka-input/kafka-input.yaml
  ```

### function-front

The function plugin mechanism is also demonstrated in `function-front`, and you can observe the following configuration in the definition of `function-front`.

  ```yaml
  apiVersion: core.openfunction.io/v1beta1
  kind: Function
  metadata:
    name: function-front
    annotations:
      plugins: |
        pre:
        - dev.openfunction.samples.plugins.ExamplePlugin
        post:
        - dev.openfunction.samples.plugins.ExamplePlugin
  ```
  >
  > `pre` defines the order of plugins that need to be called before the user function is executed
  >
  > `post` defines the order of plugins that need to be called after the user function is executed
  >
  > You can learn about the logic of these two plugins and the effect of the plugins after they are executed here: [plugins mechanism](../../../../functions-framework/README.md#plugin-mechanism)
  >

In `function-front`, we define an output (`serving.outputs`) that will point to a dapr component definition of the Kafka server.

  ```yaml
    serving:
      runtime: knative
      outputs:
        - name: target
          component: kafka-server
          type: bindings
          topic: "sample-topic"
          operation: "create"
      pubsub:
        kafka-server:
          type: bindings.kafka
          version: v1
          metadata:
            - name: brokers
              value: "kafka-server-kafka-brokers:9092"
            - name: authRequired
              value: "false"
            - name: publishTopic
              value: "sample-topic"
            - name: topics
              value: "sample-topic"
            - name: consumerGroup
              value: "function-front"
  ```

The custom content will send to the output named `target` in the `function-front` function.

For sample function below, modify the ``spec.image`` field in ``function-front.yaml`` to your own container registry address:

  ```yaml
  apiVersion: core.openfunction.io/v1beta1
  kind: Function
  metadata:
    name: function-sample
  spec:
    image: "<your registry name>/sample-knative-dapr-java:v1"
  ```

Use the following command to create `function-front` function:

  ```shell
  kubectl apply -f function-front.yaml
  ```

## Demo

Check the current function status:

  ```shell
  kubectl get functions.core.openfunction.io
  
  NAME               BUILDSTATE   SERVINGSTATE   BUILDER         SERVING         ADDRESS                                            AGE
  function-front     Succeeded    Running        builder-sk9s2   serving-dgqq7   http://function-front.default.svc.cluster.local/   11m
  kafka-input-java   Succeeded    Running        builder-29s69   serving-5mrq5                                                            2m21s
  ```

The `Function.status.addresses` field provides various methods for accessing functions.
Get `Function` addresses by running following command:

  ```shell
  kubectl get function function-front -o=jsonpath='{.status.addresses}'
  ```

You will get the following address:

  ```json
  [{"type":"External","value":"http://function-front.default.ofn.io/"},
  {"type":"Internal","value":"http://function-front.default.svc.cluster.local/"}]
  ```

> You can use the following command to create a pod in the cluster and access the function from the pod:
>
> ```shell
  > kubectl run curl --image=radial/busyboxplus:curl -i --tty
  > ```

Access functions by the internal address:

  ```shell
  [ root@curl:/ ]$ curl -d '{"message":"Awesome OpenFunction!"}' -H "Content-Type: application/json" -X POST http://function-front.default.svc.cluster.local/
  ```

Access functions by the external address:
  > To access the function via the Address of type `External` in `Funtion.status`, you should configure local domain first, see [Configure Local Domain](https://openfunction.dev/docs/operations/networking/local-domain/).

  ```shell
  [ root@curl:/ ]$ curl -d '{"message":"Awesome OpenFunction!"}' -H "Content-Type: application/json" -X POST http://function-front.default.ofn.io/
  ```

Query `function-front`'s log:

  ```shell
  kubectl logs -f \
    $(kubectl get po -l \
    openfunction.io/serving=$(kubectl get functions function-front -o jsonpath='{.status.serving.resourceRef}') \
    -o jsonpath='{.items[0].metadata.name}') \
    function
  ```

The logs are as follows:
  
  ```shell
  [main] INFO org.eclipse.jetty.server.Server - jetty-11.0.9; built: 2022-03-30T17:44:47.085Z; git: 243a48a658a183130a8c8de353178d154ca04f04; jvm 18.0.1.1+2
  [main] INFO org.eclipse.jetty.server.handler.ContextHandler - Started o.e.j.s.ServletContextHandler@-234de47e{/,null,AVAILABLE}
  [main] INFO org.eclipse.jetty.server.AbstractConnector - Started ServerConnector@dafe6b9b{HTTP/1.1, (http/1.1)}{0.0.0.0:8080}
  [main] INFO org.eclipse.jetty.server.Server - Started Server@b6fa4d5c{STARTING}[11.0.9,sto=0] @2880ms
  plugin plugin-example:v1.0.0 exec pre hook for http function at 2022-10-18 08:08:55.Z
  receive event: {"message":"Awesome OpenFunction!"}
  send to output target
  plugin plugin-example:v1.0.0 exec post hook for http function at 2022-10-18 08:08:58.Z
  ```

Query `kafka-input`'s log:
  
  ```shell
  kubectl logs -f \
    $(kubectl get po -l \
    openfunction.io/serving=$(kubectl get functions kafka-input-java -o jsonpath='{.status.serving.resourceRef}') \
    -o jsonpath='{.items[0].metadata.name}') \
    function
  ```

The logs are as follows:
  
  ```shell
  plugin plugin-example:v1.0.0 exec pre hook for binding serving-922d6-component-target-topic-jpvlq at 2022-10-18 09:55:08.Z
  receive event: "{\"message\":\"Awesome OpenFunction!\"}"
  plugin plugin-example:v1.0.0 exec post hook for binding serving-922d6-component-target-topic-jpvlq at 2022-10-18 09:55:08.Z
  ```
