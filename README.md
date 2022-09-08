# Demo Cars Function

This sample code helps get you started with a simple Go web application deployed by AWS CloudFormation to AWS Lambda and Amazon API Gateway.

## What's Here?

This sample includes:

* README.md - this file
* buildspec.yml - this file is used by AWS CodeBuild to package your
  application for deployment to AWS Lambda
* main.go - this file contains the sample Go code for the web application
* main_test.go - this file contains unit tests for the sample Go code
* template.yml - this file contains the AWS Serverless Application Model (AWS SAM) used
  by AWS CloudFormation to deploy your application to AWS Lambda and Amazon API
  Gateway.
* template-configuration.json - this file contains the project ARN with placeholders used for tagging resources with the project ID  

## Getting Started

These directions assume you want to develop on your development environment or a Cloud9 environment.

To work on the sample code, you'll need to clone your project's repository to your
local computer. If you haven't, do that first. You can find instructions in the
AWS CodeStar user guide at https://docs.aws.amazon.com/codestar/latest/userguide/getting-started.html#clone-repo

1. Install Go.  See https://golang.org/dl/ for details.

1. Install your dependencies:

    ```bash
    go mod init
    ```

1. Install the SAM CLI. For details see
https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html

1. Run the following command in your repository to build the main.go file.

    ```bash
    GOARCH=amd64 GOOS=linux go build cmd/lambda/main.go
    ```

1. Start the development server:

    ```bash
    sam local start-api -p 8080
    ```

1. Open http://127.0.0.1:8080/ in a web browser to view your webapp.

## Development

```bash
github.com/craguilar/demo-cars-fn/internal
```

Run server mode

```bash
go run cmd/server/*.go
```

### Format

```bash
gofmt -w -s .
```
