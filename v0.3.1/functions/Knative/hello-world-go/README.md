# Sample Function Go

## Prerequisites

### OpenFunction

You can refer to the [Installation Guide](https://github.com/OpenFunction/OpenFunction#install-openfunction) to setup OpenFunction.

## Deployment

1. Creating a secret

Generate a secret to access your container registry, such as one on [Docker Hub](https://hub.docker.com/) or [Quay.io](https://quay.io/).
You can create this secret by editing the ``REGISTRY_SERVER``, ``REGISTRY_USER`` and ``REGISTRY_PASSWORD`` fields in following command, and then run it.

  ```bash
  REGISTRY_SERVER=https://index.docker.io/v1/ REGISTRY_USER=<your_registry_user> REGISTRY_PASSWORD=<your_registry_password>
  kubectl create secret docker-registry push-secret \
      --docker-server=$REGISTRY_SERVER \
      --docker-username=$REGISTRY_USER \
      --docker-password=$REGISTRY_PASSWORD
  ```

2. Creating functions

   For sample function below, modify the ``spec.image`` field in ``function-sample.yaml`` to your own container registry address:

    ```yaml
    apiVersion: core.openfunction.io/v1alpha1
    kind: Function
    metadata:
      name: function-sample
    spec:
      image: "<your registry name>/sample-go-func:v0.3.1"
    ```

   Use the following command to create this Function:

    ```shell
    kubectl apply -f function-sample.yaml
    ```

3. Result observation

   You can observe the process of a function with the following command:

    ```shell
    kubectl get functions.core.openfunction.io
   
    NAME              AGE
    function-sample   5s
    ```

   You can also observe the process of a builder in the [Tekton Dashboard](https://tekton.dev/docs/dashboard/).

   Finally, you can observe the final state of the function workload in the Serving:

    ```shell
    kubectl get servings.core.openfunction.io
     
    NAME                      AGE
    function-sample-serving   15s
    ```

   You can now find out the service entry of the function with the following command:

    ```shell
    kubectl get ksvc
     
    NAME                           URL                                                                  LATESTCREATED                        LATESTREADY                          READY   REASON
    function-sample-serving-ksvc   http://function-sample-serving-ksvc.default.<external-ip>.sslip.io   function-sample-serving-ksvc-00001   function-sample-serving-ksvc-00001   True
    ```

   Or get the service address directly with the following command:

   > where` <external-ip> `indicates the external address of your gateway service.
   >
   > You can do a simple configuration to use the node ip as the `<external-ip>` as follows  (Assuming you are using Kourier as network layer of Knative). Where `1.2.3.4` can be replaced by your node ip.
   >
   > ```shell
    > kubectl patch svc -n kourier-system kourier \
    >   -p '{"spec": {"type": "LoadBalancer", "externalIPs": ["1.2.3.4"]}}'
    > 
    > kubectl patch configmap/config-domain -n knative-serving \
    >   --type merge --patch '{"data":{"1.2.3.4.sslip.io":""}}'
    > ```

    ```shell
    kubectl get ksvc function-sample-serving-ksvc -o jsonpath={.status.url}
     
    http://function-sample-serving-ksvc.default.<external-ip>.sslip.io
    ```

   Access the above service address via commands such as ``curl``:

    ```shell
    curl http://function-sample-serving-ksvc.default.<external-ip>.sslip.io
     
    Hello, World!
    ```