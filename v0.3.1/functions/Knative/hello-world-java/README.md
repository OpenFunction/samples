# Sample Function Java

## Run on OpenFunction

1. [Install OpenFunction](https://github.com/OpenFunction/OpenFunction#quickstart)
2. [Run a function](https://github.com/OpenFunction/OpenFunction#sample-run-a-function)

Definition of a ```Function``` for ```java``` is shown below:

```yaml
apiVersion: core.openfunction.io/v1alpha1
kind: Function
metadata:
  name: java-sample
spec:
  version: "v1.0.0"
  image: "<your registry name>/sample-java-func:v0.3.1"
  # port: 8080 # default to 8080
  build:
    builder: "openfunction/gcp-builder:v1"
    params:
      GOOGLE_FUNCTION_TARGET: "com.openfunction.HelloWorld"
      GOOGLE_FUNCTION_SIGNATURE_TYPE: "http"
    srcRepo:
      url: "https://github.com/OpenFunction/samples.git"
      sourceSubPath: "v0.3.1/functions/Knative/hello-world-java"
    registry:
      url: "https://index.docker.io/v1/"
      account:
        name: "basic-user-pass"
        key: "username"
  # serving:
    # runtime: "Knative" # default to Knative
```
