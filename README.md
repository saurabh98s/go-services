
So for generating the Swagger, you need to install the Swagger CLI, this is a one time operation and can be done with the following command:

`go get -u github.com/go-swagger/go-swagger/cmd/swagger ` 

Then you can generate the Swagger documentation:

swagger generate spec -o ./swagger.yaml --scan-models

Use the below command to the client folder(create one at the same level at data)
`swagger generate client -f ../swagger.yaml -A micro-services`

Once you have that run the server , windows will prompt you to add a firewall exception, click ok:

`go run main.go`

To get the documentation you should be able to do:

`curl localhost:9090/swagger.yaml`

Or open the following link in your browser:

`http://localhost:9090/docs`

## FRONTEND APP
<br> The front end app is a basic react app , which you can install by

`yarn` or
`npm install`

`yarn start` or
`npm start`

whichever you prefer<br>
you might come across a webpack error, in that case simply use the .env file in your project, with the code<br>

`SKIP_PREFLIGHT_CHECK=true`
 also delete all the node_modules if they are still in the repo.

 ### webpack version issue fix

 To fix the dependency tree, try following the steps below in the exact order:

  1. Delete package-lock.json (not package.json!) and/or yarn.lock in your project folder.
  2. Delete node_modules in your project folder.
  3. Remove "webpack" from dependencies and/or devDependencies in the package.json file in your project folder.
  4. Run npm install or yarn, depending on the package manager you use.
