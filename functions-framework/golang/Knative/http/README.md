# HTTP function

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
go run main.go
```

Access the function:

```shell
curl http://localhost:8080
```

Check the output:

```shell
I0109 14:41:13.244197    9377 knative.go:44] Knative Function serving http: listening on port 8080
I0109 14:41:22.334777    9377 framework.go:83] exec pre hooks: plugin-custom of version v1
I0109 14:41:22.334818    9377 framework.go:83] exec pre hooks: plugin-example of version v1
I0109 14:41:22.334838    9377 framework.go:94] exec post hooks: plugin-custom of version v1
I0109 14:41:22.334847    9377 framework.go:94] exec post hooks: plugin-example of version v1
I0109 14:41:22.334882    9377 plugin-example.go:79] the sum is: 4
```