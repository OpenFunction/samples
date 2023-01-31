# Build with local source code

## Prerequisites

### OpenFunction

You can refer to the [Installation Guide](https://github.com/OpenFunction/OpenFunction#install-openfunction) to setup OpenFunction.

## Deployment

1. Create secret

Generate a secret to access your container registry, such as one on [Docker Hub](https://hub.docker.com/) or [Quay.io](https://quay.io/).
You can create this secret by editing the ``REGISTRY_SERVER``, ``REGISTRY_USER`` and ``REGISTRY_PASSWORD`` fields in following command, and then run it.

  ```bash
  REGISTRY_SERVER=https://index.docker.io/v1/ REGISTRY_USER=<your_registry_user> REGISTRY_PASSWORD=<your_registry_password>
  kubectl create secret docker-registry push-secret \
      --docker-server=$REGISTRY_SERVER \
      --docker-username=$REGISTRY_USER \
      --docker-password=$REGISTRY_PASSWORD
  ```

2. Create bundle image

   A bundle image include the source code which used to build function.

   We can use this command to build a bundle image.

   ```
   docker build -t <your registry name>/local-source:latest -f Dockerfile ../../../
   docker push <your registry name>/local-source:latest
   ```

   > We suggest using a empty image such as `scratch` as the base image of the bundle image, a non-empty base image may cause the source code copy to fail.

3. Create function

   For sample function below, modify the ``spec.image`` field in ``function-sample.yaml`` to your own container registry address:

    ```yaml
    apiVersion: core.openfunction.io/v1beta1
    kind: Function
    metadata:
      name: function-sample-local
    spec:
      image: "<your registry name>/sample-local-func:v1"
    ```

   Use the following command to create this Function:

    ```shell
    kubectl apply -f function-sample.yaml
    ```

    > The `sourceSubPath` is the absolute path of the source code in the bundle image.

4. Access function

   You can refer to the [Installation Guide](https://openfunction.dev/docs/concepts/networking/function-entrypoints/) to access the function.
