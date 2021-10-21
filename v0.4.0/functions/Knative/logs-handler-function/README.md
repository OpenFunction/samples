# Function Logs Handler

> You can refer to the original article by visiting this: [以 Serverless 的方式实现 Kubernetes 日志告警](https://mp.weixin.qq.com/s/EZWYqtXJ7Cj-Yd7Fro6uyA)

## Prerequisites

Follow [this guide](../../../../Prerequisites.md#openfunction) to install OpenFunction.

Follow [this guide](../../../../Prerequisites.md#kafka) to install a Kafka server named `kafka-logs-receiver` and a Topic named `logs`.

Follow [this guide](../../../../Prerequisites.md#registry-credential) to create a registry credential.

## Deployment

Create logs handler function:

```shell
kubectl apply -f logs-handler-function.yaml
```

The logs handler function is then driven by messages from the logs topic in Kafka.

