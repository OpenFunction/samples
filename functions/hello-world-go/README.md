# Sample Function Go

## Prerequisites

### OpenFunction

You can refer to the [Installation Guide](https://github.com/OpenFunction/OpenFunction#readme) to setup OpenFunction.

## Deployment

1. Creating a secret

In order to access your container registry, you need to create a secret. You can create this secret by editing the ``username`` and ``password`` fields in following command, and then run it.

```shell
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: basic-user-pass
type: kubernetes.io/basic-auth
stringData:
  username: <USERNAME>
  password: <PASSWORD>
EOF
```

2. Creating functions

   For sample function below, modify the ``spec.image`` field in ``function-sample.yaml`` to your own container registry address:

    ```yaml
    apiVersion: core.openfunction.io/v1alpha1
    kind: Function
    metadata:
      name: function-sample
    spec:
      image: "<your registry name>/sample-go-func:latest"
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