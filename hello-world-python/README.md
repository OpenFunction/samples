# Sample Function Python

## Run on OpenFunction

1. [Install OpenFunction](https://github.com/OpenFunction/OpenFunction#quickstart)
2. [Run a function](https://github.com/OpenFunction/OpenFunction#sample-run-a-function)

Definition of a ```Function``` for ```python``` is shown below:

```yaml
apiVersion: core.openfunction.io/v1alpha1
kind: Function
metadata:
  name: python-sample
spec:
  funcName: "hello_world"
  funcType: "http"
  funcVersion: "v1.0.0"
  builder: "openfunction/gcp-builder:v1"
  source:
    url: "https://github.com/OpenFunction/function-samples.git"
    sourceSubPath: "hello-world-python"
  image: "<your registry name>/sample-python-func:latest"
  registry:
    url: "https://index.docker.io/v1/"
    account:
      name: "basic-user-pass"
      key: "username"
  runtime: "Knative"
  port: 8080
```