# Sample Function Go

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

2. Create function

   For sample function below, modify the ``spec.image`` field in ``function-sample.yaml`` to your own container registry address:

    ```yaml
    apiVersion: core.openfunction.io/v1beta1
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

3. Access function

   You can observe the process of a function with the following command:

    ```shell
   kubectl get functions.core.openfunction.io
   
   NAME              BUILDSTATE   SERVINGSTATE   BUILDER         SERVING         URL                                              AGE
   function-sample   Succeeded    Running        builder-jgnzp   serving-q6wdp   http://openfunction.io/default/function-sample   22m
    ```

   The `URL` is the address provided by the OpenFunction Domain that can be accessed. To access the function via this URL address, you need to make sure that DNS can resolve this address.

   > You can use the following command to create a pod in the cluster and access the function from the pod:
   >
   > ```shell
   > kubectl run curl --image=radial/busyboxplus:curl -i --tty
   > ```
   
   Access the function via `URL`:
   
   ```shell
   [ root@curl:/ ]$ curl http://openfunction.io.svc.cluster.local/default/function-sample/World
   Hello, World!
   
   [ root@curl:/ ]$ curl http://openfunction.io.svc.cluster.local/default/function-sample/OpenFunction
   Hello, OpenFunction!
   ```
   
   There is also an alternative way to trigger the function via the access address provided by the Knative Services:

    ```shell
    kubectl get ksvc
     
    NAME                       URL                                                            LATESTCREATED                   LATESTREADY                     READY   REASON
    serving-q6wdp-ksvc-wk6mv   http://serving-q6wdp-ksvc-wk6mv.default.<external-ip>.sslip.io   serving-q6wdp-ksvc-wk6mv-v100   serving-q6wdp-ksvc-wk6mv-v100   True
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
    kubectl get ksvc serving-q6wdp-ksvc-wk6mv -o jsonpath={.status.url}
     
    http://serving-q6wdp-ksvc-wk6mv.default.<external-ip>.sslip.io
    ```
   
   Access the above service address via commands such as ``curl``:
   
    ```shell
    curl http://serving-q6wdp-ksvc-wk6mv.default.<external-ip>.sslip.io/World
     
    Hello, World!
    
    curl http://serving-q6wdp-ksvc-wk6mv.default.<external-ip>.sslip.io/OpenFunction
     
    Hello, OpenFunction!
    ```