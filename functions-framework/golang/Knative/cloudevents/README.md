# CloudEvents function

## FUNC_CONTEXT

```json
{
  "name": "function-demo",
  "version": "v1.0.0",
  "port": "8080",
  "runtime": "Knative",
  "prePlugins": ["plugin-custom", "plugin-example"],
  "postPlugins": ["plugin-custom", "plugin-example"]
}
```

## Run

Export `FUNC_CONTEXT` environment variable:

```shell
export FUNC_CONTEXT='{"name":"function-demo","version":"v1.0.0","port":"8080","runtime":"Knative","prePlugins":["plugin-custom","plugin-example"],"postPlugins":["plugin-custom","plugin-example"]}'
```

Start the function:

```shell
go run main.go plugin.go
```

Access the function:

```shell
curl -v "http://localhost:8080" \
  -X POST \
  -H "Ce-Specversion: 1.0" \
  -H "Ce-Type: dev.knative.samples.helloworld" \
  -H "Ce-Source: dev.knative.samples/helloworldsource" \
  -H "Ce-Id: 536808d3-88be-4077-9d7a-a3f162705f79" \
  -H "Content-Type: application/json" \
  -d '{"msg":"Hello Knative!"}'
```

Check the output:

```shell
I0111 10:08:00.971063  373434 framework.go:83] exec pre hooks: plugin-custom of version v1
I0111 10:08:00.971100  373434 framework.go:83] exec pre hooks: plugin-example of version v1
{"msg":"Hello Knative!"}
I0111 10:08:00.971111  373434 framework.go:94] exec post hooks: plugin-custom of version v1
I0111 10:08:00.971253  373434 framework.go:94] exec post hooks: plugin-example of version v1
I0111 10:08:00.971282  373434 plugin-example.go:79] the sum is: 4
```

