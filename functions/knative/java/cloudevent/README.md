# Sample Function Java

## Prerequisites

### OpenFunction

You can refer to the [Installation Guide](https://openfunction.dev/docs/getting-started/installation/) to setup OpenFunction.

## Run it locally

Build the function locally

  ```sh
  pack build func-cloudevent-java --path ../ --builder openfunction/builder-java:v2-18 --env FUNC_NAME="dev.openfunction.samples.CloudEventFunctionImpl"  --env FUNC_CLEAR_SOURCE=true
  ```

Run the function

  ```sh
  docker run --rm --env="FUNC_CONTEXT={\"name\":\"CloudEvent\",\"version\":\"v1.0.0\",\"port\":\"8080\",\"runtime\":\"Knative\"}" --env="CONTEXT_MODE=self-host" --name func-cloudevent-java -p 8080:8080 func-cloudevent-java
  ```

Send a request

  ```sh
  curl -X POST "http://localhost:8080/" \
     -H "Content-Type: application/cloudevents+json" \
     -d '{"specversion":"1.0","type":"dev.knative.samples.helloworld","source":"dev.knative.samples/helloworldsource","id":"536808d3-88be-4077-9d7a-a3f162705f79","data":"{\"data\":\"hello world\"}"}'
  # in docker compose terminal:
  # receive event: {"data":"hello world"}

  # Binary CloudEvent
  curl "http://localhost:8080/" \
    -X POST \
    -H "Ce-Specversion: 1.0" \
    -H "Ce-Type: dev.knative.samples.helloworld" \
    -H "Ce-Source: dev.knative.samples/helloworldsource" \
    -H "Ce-Subject: 123" \
    -H "Ce-Id: 536808d3-88be-4077-9d7a-a3f162705f79" \
    -H "Content-Type: application/json" \
    -d '{"data":"hello world"}'
  # in docker compose terminal:
  # receive event: {"data":"hello world"}
  ```

## Deployment

1. Create secret

Follow [this guide](https://openfunction.dev/docs/getting-started/quickstarts/prerequisites/#registry-credential) to create a registry credential.

2. Create function
   
For sample function below, modify the ``spec.image`` field in ``function-sample.yaml`` to your own container registry address:

  ```yaml
  apiVersion: core.openfunction.io/v1beta1
  kind: Function
  metadata:
    name: function-cloudevent-java
  spec:
    image: "<your registry name>/sample-java-func:v1"
  ```

Use the following command to create this Function:

  ```shell
  kubectl apply -f function-sample.yaml
  ```

3. Access function

You can observe the process of a function with the following command:

  ```shell
  kubectl get functions.core.openfunction.io

  NAME                       BUILDSTATE   SERVINGSTATE   BUILDER         SERVING         ADDRESS                                                      AGE
  function-cloudevent-java   Succeeded    Running        builder-75mq9   serving-xzx49   http://function-cloudevent-java.default.svc.cluster.local/   9m1s
  ```

The `Function.status.addresses` field provides various methods for accessing functions.
Get `Function` addresses by running following command:

  ```shell
  kubectl get function function-cloudevent-java -o=jsonpath='{.status.addresses}'
  ```

You will get the following address:

  ```json
  [{"type":"External","value":"http://function-cloudevent-java.default.ofn.io/"},
  {"type":"Internal","value":"http://function-cloudevent-java.default.svc.cluster.local/"}]
  ```

  > You can use the following command to create a pod in the cluster and access the function from the pod:
  >
  > ```shell
  > kubectl run curl --image=radial/busyboxplus:curl -i --tty
  > ```

Access functions by the internal address:

  ```shell
  curl http://function-cloudevent-java.default.svc.cluster.local \
    -X POST \
    -H "Ce-Specversion: 1.0" \
    -H "Ce-Type: dev.knative.samples.helloworld" \
    -H "Ce-Source: dev.knative.samples/helloworldsource" \
    -H "Ce-Subject: 123" \
    -H "Ce-Id: 536808d3-88be-4077-9d7a-a3f162705f79" \
    -H "Content-Type: application/json" \
    -d '{"data":"hello world"}'
  ```

Access functions by the external address:
  > To access the function via the Address of type `External` in `Funtion.status`, you should configure local domain first, see [Configure Local Domain](https://openfunction.dev/docs/operations/networking/local-domain/).

  ```shell
  curl http://function-cloudevent-java.default.ofn.io \
    -X POST \
    -H "Ce-Specversion: 1.0" \
    -H "Ce-Type: dev.knative.samples.helloworld" \
    -H "Ce-Source: dev.knative.samples/helloworldsource" \
    -H "Ce-Subject: 123" \
    -H "Ce-Id: 536808d3-88be-4077-9d7a-a3f162705f79" \
    -H "Content-Type: application/json" \
    -d '{"data":"hello world"}'
  ```
   
Query `function-cloudevent-java`'s log:

  ```shell
  kubectl logs -f \
    $(kubectl get po -l \
    openfunction.io/serving=$(kubectl get functions function-cloudevent-java -o jsonpath='{.status.serving.resourceRef}') \
    -o jsonpath='{.items[0].metadata.name}') \
    function
  ```

The logs are as follows:

  ```shell
  [main] INFO org.eclipse.jetty.server.Server - jetty-11.0.9; built: 2022-03-30T17:44:47.085Z; git: 243a48a658a183130a8c8de353178d154ca04f04; jvm 18.0.1.1+2
  [main] INFO org.eclipse.jetty.server.handler.ContextHandler - Started o.e.j.s.ServletContextHandler@64ed162{/,null,AVAILABLE}
  [main] INFO org.eclipse.jetty.server.AbstractConnector - Started ServerConnector@7b929271{HTTP/1.1, (http/1.1)}{0.0.0.0:8080}
  [main] INFO org.eclipse.jetty.server.Server - Started Server@41b1c863{STARTING}[11.0.9,sto=0] @836ms
  receive event: {"data":"hello world"}
  ```
