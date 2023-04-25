# golang-aws-serverless-api
Building a Serverless API in AWS with Go


# hello-world-lambda

Learn how to build your first Lambda Function with Go! We'll create a basic hello-world function from scratch, upload it to AWS, and test it within our browser. This is the first in a series of building a Serverless API in AWS with Go!

## initialize go module

```
$ go mod init github.com/favtuts/golang-aws-serverless-api
```

## create entry point

Create `main.go` and print simple message `Hello world`, then run it by command
```
$ go run ./main.go

2023/04/25 15:16:22 Hello world
```

## install go sdk for aws

```
$ go get -u github.com/aws/aws-lambda-go

go: downloading github.com/aws/aws-lambda-go v1.40.0
go: added github.com/aws/aws-lambda-go v1.40.0
```

## handle api gateway proxy request

Build the lambda
```
$ export GOOS="linux"
$ go build main.go
```

## zip the executable file 

```
$ sudo apt install zip unzip
$ zip --version
$ zip golang-hello-world-lambda.zip main
```

## create AWS Lambda from console

Create new AWS Lambda from scratch and then upload a zip file, then edit the runtime settings change handler from `hello` to `main` because our entry point is main function.

## test the AWS lambda with event

Execute the test event
```
START RequestId: e8b036b9-5b1f-40f6-a7e4-93d1698e453b Version: $LATEST
2023/04/25 16:32:24 Hello world
END RequestId: e8b036b9-5b1f-40f6-a7e4-93d1698e453b
REPORT RequestId: e8b036b9-5b1f-40f6-a7e4-93d1698e453b	Duration: 2.33 ms	Billed Duration: 3 ms	Memory Size: 512 MB	Max Memory Used: 30 MB	Init Duration: 97.32 ms	
```

# person json lanbda

We will process the post json data as a Person via API Gateway and our Lambda need to process the request and return a specify message based on incommed request json data.

You can test the lambda by creating new event with template: `apigateway-aws-proxy`, more detail in the file `person-event.json`, and it return the result:
```json
{
  "statusCode": 200,
  "headers": null,
  "multiValueHeaders": null,
  "body": "{\"message\":\"Hello Shane Van Boening\"}"
}
```

## testing the lambda

Install the Testify library
```
$ go get -u github.com/stretchr/testify

go: downloading github.com/stretchr/testify v1.8.2
go: downloading github.com/stretchr/objx v0.5.0
go: added github.com/stretchr/objx v0.5.0
go: upgraded github.com/stretchr/testify v1.7.2 => v1.8.2
```

Reference to an example of using [AWS Lambda with GO](https://github.com/aws-samples/lambda-go-samples)