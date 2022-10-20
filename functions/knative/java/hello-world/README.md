# Sample Function Java

## Prerequisites

### OpenFunction

You can refer to the [Installation Guide](https://openfunction.dev/docs/getting-started/installation/) to setup OpenFunction.

## Run it locally

Build the function locally

  ```sh
  pack build func-http-java --path ../ --builder openfunction/builder-java:v2-18 --env FUNC_NAME="dev.openfunction.samples.HttpFunctionImpl"  --env FUNC_CLEAR_SOURCE=true
  ```

Run the function

  ```sh
  docker run --rm --env="FUNC_CONTEXT={\"name\":\"HelloWorld\",\"version\":\"v1.0.0\",\"port\":\"8080\",\"runtime\":\"Knative\"}" --env="CONTEXT_MODE=self-host" --name func-http-java -p 8080:8080 func-http-java
  ```

Send a request

  ```sh
  curl http://localhost:8080/
  # Hello World
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
    name: function-http-java
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
   
  NAME                 BUILDSTATE   SERVINGSTATE   BUILDER         SERVING         ADDRESS                                                         AGE
  function-http-java   Succeeded    Running        builder-jgnzp   serving-q6wdp   http://function-http-java.default.svc.cluster.local/               22m
  ```

The `Function.status.addresses` field provides various methods for accessing functions.
Get `Function` addresses by running following command:

  ```shell
  kubectl get function function-http-java -o=jsonpath='{.status.addresses}'
  ```

You will get the following address:
   
  ```json
  [{"type":"External","value":"http://function-http-java.default.ofn.io/"},
  {"type":"Internal","value":"http://function-http-java.default.svc.cluster.local/"}]
  ```

  > You can use the following command to create a pod in the cluster and access the function from the pod:
  >
  > ```shell
  > kubectl run curl --image=radial/busyboxplus:curl -i --tty
  > ```

Access functions by the internal address:

  ```shell
  curl http://function-http-java.default.svc.cluster.local
  Hello, World!
   ```

Access functions by the external address:
  > To access the function via the Address of type `External` in `Funtion.status`, you should configure local domain first, see [Configure Local Domain](https://openfunction.dev/docs/operations/networking/local-domain/).

  ```shell
  curl http://function-http-java.default.ofn.io
  Hello, World!
  ```
   