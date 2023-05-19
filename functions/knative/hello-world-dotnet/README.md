# Sample Function Dotnet

## Run on OpenFunction

1. [Install OpenFunction](https://github.com/OpenFunction/OpenFunction#install-openfunction)
2. [Refer to the go function sample](../hello-world-go/README.md)

Definition of a ```Function``` for ```dotnet``` is shown below:

```yaml
apiVersion: core.openfunction.io/v1beta2
kind: Function
metadata:
  name: dotnet-sample
spec:
  version: "v1.0.0"
  image: "<your registry name>/sample-dotnet-func:v1"
  imageCredentials:
    name: push-secret
  build:
    builder: "openfunction/gcp-builder:v1"
    env:
      GOOGLE_FUNCTION_TARGET: "helloworld"
      GOOGLE_FUNCTION_SIGNATURE_TYPE: "http"
    srcRepo:
      url: "https://github.com/OpenFunction/samples.git"
      sourceSubPath: "functions/knative/hello-world-dotnet"
      revision: "release-0.6"
  serving:
    template:
      containers:
        - name: function # DO NOT change this
          imagePullPolicy: IfNotPresent 
    triggers:
      http:
        port: 8080
```
