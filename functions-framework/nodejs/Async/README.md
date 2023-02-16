
# Async Runtime use plugin example

When you write a  plugin (like this )[plugins](plugins) , you can use it this way
```
apiVersion: core.openfunction.io/v1beta1
kind: Function
metadata:
  name: sample-node-async-pubsub
  annotations:
    plugins: | 
      pre:
        - sample
      post:
        - sample
spec:
  version: v2.0.0
  image: your-docker-registry/image:v1
...
```