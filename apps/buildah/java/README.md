# Sample Function Java

## Prerequisites

### OpenFunction

You can refer to the [Installation Guide](https://openfunction.dev/docs/getting-started/installation/) to setup OpenFunction.

## Deployment

1. Create secret

Follow [this guide](https://openfunction.dev/docs/getting-started/quickstarts/prerequisites/#registry-credential) to create a registry credential.

2. Create function

For sample function below, modify the ``spec.image`` field in ``function-buildah-java.yaml`` to your own container registry address:

  ```yaml
  apiVersion: core.openfunction.io/v1beta1
  kind: Function
  metadata:
    name: function-buildah-java
  spec:
    image: "<your registry name>/sample-buildah-java:latest"
  ```

Use the following command to create this Function:

  ```shell
  kubectl apply -f function-buildah-java.yaml
  ```

3. Access function

You can observe the process of a function with the following command:

  ```shell
  kubectl get functions.core.openfunction.io
   
  NAME                    BUILDSTATE   SERVINGSTATE   BUILDER         SERVING         ADDRESS                                                   AGE
  function-buildah-java   Succeeded    Running        builder-jgnzp   serving-q6wdp   http://function-buildah-java.default.svc.cluster.local/   22m
  ```

The `Function.status.addresses` field provides various methods for accessing functions.
Get `Function` addresses by running following command:

  ```shell
  kubectl get function function-buildah-java -o=jsonpath='{.status.addresses}'
  ```

You will get the following address:
   
  ```json
  [{"type":"External","value":"http://function-buildah-java.default.ofn.io/"},
  {"type":"Internal","value":"http://function-buildah-java.default.svc.cluster.local/"}]
  ```

  > You can use the following command to create a pod in the cluster and access the function from the pod:
  >
  > ```shell
  > kubectl run curl --image=radial/busyboxplus:curl -i --tty
  > ```

Access functions by the internal address:

  ```shell
  curl http://function-buildah-java.default.svc.cluster.local
   ```

Access functions by the external address:
  > To access the function via the Address of type `External` in `Funtion.status`, you should configure local domain first, see [Configure Local Domain](https://openfunction.dev/docs/operations/networking/local-domain/).

  ```shell
  curl http://function-buildah-java.default.ofn.io
  ```
   