# Cloudevent

The functions framework can unmarshall incoming [CloudEvents](http://cloudevents.io/) payloads to a `cloudevent` object. Note that your function must use the `cloudevent-style` function signature.

```js
exports.helloCloudEvents = (cloudevent) => {
  return cloudevent
}
```

Your `package.json`

```json
  "scripts": {
    "start": "functions-framework --target=helloCloudEvents --source=cloudevent"
  }
```

**A binary one**

See [payload content](https://github.com/OpenFunction/functions-framework-nodejs/blob/main/mock/payload/binary.json)

```bash
$ curl -X POST \
     -d'@../mock/payload/binary.json' \
     -H'Content-Type:application/json' \
     -H'ce-specversion:1.0' \
     -H'ce-type:com.github.pull.create' \
     -H'ce-source:https://github.com/cloudevents/spec/pull/123' \
     -H'ce-id:45c83279-c8a1-4db6-a703-b3768db93887' \
     -H'ce-time:2019-11-06T11:17:00Z' \
     -H'ce-myextension:extension value' \
     http://localhost:8080/
# The response is
{
    "datacontenttype": "application/json",
    "data": {
        "runtime": "cloudevent"
    },
    "specversion": "1.0",
    "type": "com.github.pull.create",
    "source": "https://github.com/cloudevents/spec/pull/123",
    "id": "45c83279-c8a1-4db6-a703-b3768db93887",
    "time": "2019-11-06T11:17:00Z",
    "myextension": "extension value"
}
```

**A structed one**

See [payload content](https://github.com/OpenFunction/functions-framework-nodejs/blob/main/mock/payload/structured.json)

```bash
$ curl -X POST \
     -d'@../mock/payload/structured.json' \
     -H'Content-Type:application/cloudevents+json' \
     http://localhost:8080/
# The response is
{
    "specversion": "1.0",
    "type": "com.github.pull.create",
    "source": "https://github.com/cloudevents/spec/pull/123",
    "id": "b25e2717-a470-45a0-8231-985a99aa9416",
    "time": "2019-11-06T11:08:00Z",
    "datacontenttype": "application/json",
    "data": {
        "framework": "openfunction"
    }
}
```

