# Sample Function Node

## Run on OpenFunction

1. [Install OpenFunction](https://github.com/OpenFunction/OpenFunction#install-openfunction)
2. [Refer to the go function sample for the full test steps](../../hello-world-go/README.md)

You can define a common JavaScript Function as below:

```yaml
apiVersion: core.openfunction.io/v1beta1
kind: Function
metadata:
  name: node-sample
spec:
  version: "v1.0.0"
  image: "<your registry name>/sample-node-func:v1"
  imageCredentials:
    name: push-secret
  port: 8080 # default to 8080
  build:
    builder: "openfunction/builder-node:latest"
    env:
      FUNC_NAME: "helloWorld"
      FUNC_TYPE: "http"
    srcRepo:
      url: "https://github.com/OpenFunction/samples.git"
      sourceSubPath: "functions/knative/hello-world-node/commonjs"
      revision: "main"
  serving:
    runtime: "knative" # default to knative
    template:
      containers:
        - name: function # DO NOT change this
          imagePullPolicy: IfNotPresent 
```