
So for generating the Swagger, you need to install the Swagger CLI, this is a one time operation and can be done with the following command:

`go get -u github.com/go-swagger/go-swagger/cmd/swagger ` 

Then you can generate the Swagger documentation:

swagger generate spec -o ./swagger.yaml --scan-models

Once you have that run the server , windows will prompt you to add a firewall exception, click ok:

`go run main.go`

To get the documentation you should be able to do:

`curl localhost:9090/swagger.yaml`

Or open the following link in your browser:

`http://localhost:9090/docs`