## running the rest-service

You should be able to run this rest-service by doing the following steps:
1. Make sure you have your $GOPATH environment variable set.
2. Run `go get github.com/carrpet/backend-developer-tests`
3. Navigate to the $GOPATH/src/github.com/carrpet/backend-developer-tests/rest-service folder
4. Execute `dep ensure` to download the dependencies
5. Build the app by executing `go build main.go` which will output an executable called `main`
6. Run the app in the terminal by executing `./main`
7. The service will be available on `localhost:8080`