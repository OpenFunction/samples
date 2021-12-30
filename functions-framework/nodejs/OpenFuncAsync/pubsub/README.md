# Pub/Sub Example

1. Create a `package.json` file using `npm init`:

   ```bash
   $ npm init
   ```

2. Create an `index.js` file with the following contents:

   ```js
   exports.helloWorld = async (data) => {
     console.log(data)
     const delay = ms => new Promise(resolve => setTimeout(resolve, ms))
     await delay(5000)
     return data
   }
   ```
   
3. Add your `config.json`

   ```json
   {
     "input": {
       "name": "pubsub",
       "uri": "test",
       "params": {
         "type": "pubsub"
       }
     },
     "outputs": {
       "pubsub": {
         "uri": "test",
         "params": {
           "type": "pubsub"
         }
       }
     }
   }
   ```

4. Now install the Functions Framework:

   ```bash
   $ npm install @openfunction/functions-framework
   ```

5. Add a `start` script to `package.json`, with configuration passed via command-line arguments:

   ```json
     "scripts": {
       "start": "functions-framework --target=helloWorld --port 4000"
     }
   ```

6. Note that if we want to test it locally, we need a Dapr config file:

   ```yaml
   apiVersion: dapr.io/v1alpha1
   kind: Component
   metadata:
     name: pubsub # midlleware component name
   spec:
     type: pubsub.redis
     version: v1
     metadata:
     - name: redisHost
       value: ${YOUR_REDIS_ADDRESS}
     - name: redisPassword
       value: ""
   ```

7. Use `dapr` to run  `npm start` and start the built-in local development server:

   ```bash
   $ dapr run --app-id hello-world --app-port 4000 --dapr-http-port 3500 --components-path ./pubsub.yaml  npm start
   ```

8. Send requests to this function using `curl` from another terminal window:

   ```bash
   $ curl -X POST \
        -d'@../mock/payload/structured.json' \
        -H'Content-Type:application/cloudevents+json' \
        http://localhost:4000/test
   ```

9. The `structured.json` is like below: (note that in dapr, pubsub mode respects the cloud event style)

   ```json
   {
     "specversion":"1.0",
     "type":"com.github.pull.create",
     "source":"https://github.com/cloudevents/spec/pull/123",
     "id":"b25e2717-a470-45a0-8231-985a99aa9416",
     "time":"2019-11-06T11:08:00Z",
     "datacontenttype":"application/json",
     "data":{
       "framework":"openfunction"
     }
   }
   ```

10. And the output will be like:

    ```shell
    == APP == { framework: 'openfunction' }
    == APP == ::ffff:127.0.0.1 - - [07/Aug/2021:09:09:31 +0000] "POST /test HTTP/1.1" 200 - "-" "curl/7.58.0"
    == APP == { data: { framework: 'openfunction' } }
    == APP == ::ffff:127.0.0.1 - - [07/Aug/2021:09:09:36 +0000] "POST /test HTTP/1.1" 200 - "-" "fasthttp"
    == APP == { data: { data: { framework: 'openfunction' } } }
    == APP == ::ffff:127.0.0.1 - - [07/Aug/2021:09:09:41 +0000] "POST /test HTTP/1.1" 200 - "-" "fasthttp"
    == APP == { data: { data: { data: [Object] } } }
    == APP == ::ffff:127.0.0.1 - - [07/Aug/2021:09:09:46 +0000] "POST /test HTTP/1.1" 200 - "-" "fasthttp"
    ```

    

