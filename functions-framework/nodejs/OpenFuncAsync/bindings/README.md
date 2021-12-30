# BIndings Example

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
       "name": "sample-topic",
       "uri": "test",
       "params": {
         "type": "bindings"
       }
     },
     "outputs": {
       "sample-topic": {
         "uri": "test",
         "params": {
           "type": "bindings",
           "operation": "create"
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
     name: sample-topic
   spec:
     type: bindings.kafka
     version: v1
     metadata:
     # Kafka broker connection setting
     - name: brokers
       value: localhost:9092
     # consumer configuration: topic and consumer group
     - name: topics
       value: sample
     - name: consumerGroup
       value: group1
     # publisher configuration: topic
     - name: publishTopic
       value: sample
     - name: authRequired
       value: "false"
   ```

7. Use `dapr` to run  `npm start` and start the built-in local development server:

   ```bash
   $ dapr run --app-id hello-world --app-port 4000 --dapr-http-port 3500 --components-path ./bindings.yaml  npm start
   ```

8. Send requests to this function using `curl` from another terminal window:

   ```bash
   $ curl -X POST \
        -d'{"data": "hello"}' \
        -H'Content-Type:application/json' \
        http://localhost:4000/test
   ```

9. And the output will be like:

   ```shell
   == APP == { data: 'hello' }
   == APP == ::ffff:127.0.0.1 - - [07/Aug/2021:09:09:31 +0000] "POST /test HTTP/1.1" 200 - "-" "curl/7.58.0"
   == APP == { data: 'hello' }
   == APP == ::ffff:127.0.0.1 - - [07/Aug/2021:09:09:36 +0000] "POST /test HTTP/1.1" 200 - "-" "fasthttp"
   == APP == { data: 'hello' }
   == APP == ::ffff:127.0.0.1 - - [07/Aug/2021:09:09:41 +0000] "POST /test HTTP/1.1" 200 - "-" "fasthttp"
   == APP == { data: 'hello' }
   == APP == ::ffff:127.0.0.1 - - [07/Aug/2021:09:09:46 +0000] "POST /test HTTP/1.1" 200 - "-" "fasthttp"
   ```

   

