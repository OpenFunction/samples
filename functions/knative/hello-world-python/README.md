# Sample Function Python

## Run on OpenFunction

1. [Install OpenFunction](https://github.com/OpenFunction/OpenFunction#install-openfunction)
2. [Refer to the go function sample](../hello-world-go/README.md)

## Run it locally

Build the function locally

```sh
pack build python-sample --builder openfunction/gcp-builder:v1 --env GOOGLE_FUNCTION_TARGET=hello_world
```

Run the function

```sh
docker run --rm --env="FUNC_CONTEXT={\"name\":\"python-sample\",\"version\":\"v1.0.0\",\"port\":\"8080\",\"runtime\":\"Knative\"}" --env="CONTEXT_MODE=self-host" --name python-sample -p 8080:8080 python-sample
```

Send a request

```sh
curl http://localhost:8080
# hello, world
```

Definition of a `Function` for `python` is shown below:

```yaml
apiVersion: core.openfunction.io/v1beta2
kind: Function
metadata:
  name: python-sample
spec:
  version: "v1.0.0"
  image: "<your registry name>/sample-python-func:v1"
  imageCredentials:
    name: push-secret
  build:
    builder: "openfunction/gcp-builder:v1"
    env:
      GOOGLE_FUNCTION_TARGET: "hello_world"
      GOOGLE_FUNCTION_SIGNATURE_TYPE: "http"
      GOOGLE_FUNCTION_SOURCE: "main.py"
    srcRepo:
      url: "https://github.com/OpenFunction/samples.git"
      sourceSubPath: "functions/knative/hello-world-python"
      revision: "main"
  serving:
    runtime: knative # default to knative
    template:
      containers:
        - name: function # DO NOT change this
          imagePullPolicy: IfNotPresent 
    triggers:
      http:
        port: 8080
```
