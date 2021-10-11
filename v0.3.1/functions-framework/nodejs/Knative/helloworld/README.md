# Hello World Example

1. Create a `package.json` file using `npm init`:

   ```bash
   $ npm init
   ```

2. Create an `index.js` file with the following contents:

   ```js
   exports.helloWorld = (req, res) => {
     res.send('Hello, World');
   };
   ```

3. Now install the Functions Framework:

   ```bash
   $ npm install @openfunction/functions-framework
   ```

4. Add a `start` script to `package.json`, with configuration passed via command-line arguments:

   ```json
     "scripts": {
       "start": "functions-framework --target=helloWorld --source http"
     }
   ```

5. Use `npm start` to start the built-in local development server:

   ```bash
   $ npm start
   ...
   The functionModulePath is: /.../.../index.js
   Openfunction functions framework listening at http://localhost:8080
   ```

6. Send requests to this function using `curl` from another terminal window:

   ```bash
   $ curl localhost:8080
   # Output: Hello, World
   ```
