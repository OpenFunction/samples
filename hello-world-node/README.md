# Sample Function Node

## Run on OpenFunction

1. [Install OpenFunction](https://github.com/OpenFunction/OpenFunction#quickstart)
2. [Run a function](https://github.com/OpenFunction/OpenFunction#sample-run-a-function)

Definition of a ```Function``` for ```node``` is shown below:

```yaml
apiVersion: core.openfunction.io/v1alpha1
kind: Function
metadata:
  name: node-sample
spec:
  funcName: "helloWorld"
  funcType: "http"
  funcVersion: "v1.0.0"
  builder: "openfunction/gcp-builder:v1"
  source:
    url: "https://github.com/OpenFunction/function-samples.git"
    sourceSubPath: "hello-world-node"
  image: "<your registry name>/sample-node-func:latest"
  registry:
    url: "https://index.docker.io/v1/"
    account:
      name: "basic-user-pass"
      key: "username"
  runtime: "Knative"
  port: 8080
```