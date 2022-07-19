# Path Parameters Function Go

## Prerequisites

### OpenFunction

You can refer to the [Installation Guide](https://github.com/OpenFunction/OpenFunction#install-openfunction) to setup OpenFunction.

## Run it locally

Build the function locally

```sh
pack build sample-go-path-params-func --builder openfunction/builder-go:v2.3.0-1.16 --env FUNC_NAME="pathParametersFunction"  --env FUNC_CLEAR_SOURCE=true
```

Run the function

```sh
docker compose up
```

Send a request

```sh

# http
curl -X POST "http://localhost:8080/hello/openfunction"
# {"hello":"openfunction"}% 

# cloudevent
curl -X POST "http://localhost:8080/foo/openfunction" \
   -H "Content-Type: application/cloudevents+json" \
   -d '{"specversion":"1.0","type":"dev.knative.samples.helloworld","source":"dev.knative.samples/helloworldsource","id":"536808d3-88be-4077-9d7a-a3f162705f79","data":{"data":"hello"}}'
# in docker compose terminal:
# cloudevent - Data: {"{\"data\":\"hello\"}":"openfunction"}


# http
curl -X POST "http://localhost:8080/bar/openfunction" \
  -d 'hello'
# {"hello":"openfunction"}%  

# Structured CloudEvent
curl -X POST "http://localhost:8080/bar/openfunction" \
   -H "Content-Type: application/cloudevents+json" \
   -d '{"specversion":"1.0","type":"dev.knative.samples.helloworld","source":"dev.knative.samples/helloworldsource","id":"536808d3-88be-4077-9d7a-a3f162705f79","data":{"data":"hello"}}'
# {"{\"data\":\"hello\"}":"openfunction"}%

# Binary CloudEvent
curl "http://localhost:8080/bar/openfunction" \
  -X POST \
  -H "Ce-Specversion: 1.0" \
  -H "Ce-Type: dev.knative.samples.helloworld" \
  -H "Ce-Source: dev.knative.samples/helloworldsource" \
  -H "Ce-Subject: 123" \
  -H "Ce-Id: 536808d3-88be-4077-9d7a-a3f162705f79" \
  -H "Content-Type: application/json" \
  -d '{"data":"hello"}'
# {"{\"data\":\"hello\"}":"openfunction"}%
```


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
      image: "<your registry name>/sample-go-path-params-func:latest"
    ```

   Use the following command to create this Function:

    ```shell
    kubectl apply -f function-sample.yaml
    ```

3. Access function

   You can observe the process of a function with the following command:

    ```shell
   kubectl get functions.core.openfunction.io
   
   NAME              BUILDSTATE   SERVINGSTATE   BUILDER         SERVING         URL                                                        AGE
   function-sample   Succeeded    Running        builder-jgnzp   serving-gsx8g   http://openfunction.io/modelmesh-serving/function-sample   56s
    ```
   
   Trigger the function via the access address provided by the Knative Services:

    ```shell
    kubectl get ksvc
     
    NAME                       URL                                                            LATESTCREATED                   LATESTREADY                     READY   REASON
    serving-gsx8g-ksvc-6fv9l   http://serving-gsx8g-ksvc-6fv9l.default.<external-ip>.sslip.io   serving-gsx8g-ksvc-6fv9l-v100   serving-gsx8g-ksvc-6fv9l-v100   True
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
    kubectl get ksvc serving-gsx8g-ksvc-6fv9l -o jsonpath={.status.url}
     
    http://serving-gsx8g-ksvc-6fv9l.default.<external-ip>.sslip.io
    ```
   
   Access the above service address via commands such as ``curl``:
   
    ```shell
    curl http://serving-gsx8g-ksvc-6fv9l.default.<external-ip>.sslip.io/hello/openfunction
     
    {"hello":"openfunction"}% 
    ```