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