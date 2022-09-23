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

## Development

### AWS Lambda execution

These directions assume you want to develop on your development environment.

To work on the sample code, you'll need to clone your project's repository to your
local computer. If you haven't, do that first.

1. Install Go.  See https://golang.org/dl/ for details.

1. Install your dependencies:

    ```bash
    go mod init
    ```

    or if already installed

    ```bash
    go mod tidy
    ```

1. Install the SAM CLI. For details see
https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html

1. Run the following command in your repository to build the main.go file.

    ```bash
    GOARCH=amd64 GOOS=linux go build -o main cmd/http/lambda/*.go
    ```

1. Start the development server:

    ```bash
    sam local start-api -p 8080
    ```

1. Open http://127.0.0.1:8080/ in a web browser to view your webapp or execute

  ```bash
  scripts/test-integration.sh
  ```

### Server mode

Run server mode

```bash
go run cmd/http/server/*.go
```

Then open http://127.0.0.1:8080/ in a web browser to view your webapp or execute

  ```bash
  scripts/test-integration.sh
  ```

### Dynamo DB 

https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.html

## Contributing

### Vulnerability checking

Requires Go version 1.18 - see https://go.dev/blog/vuln

### Format

```bash
gofmt -w -s .
```

## Deployment

Validate your Cloudformation template using below command: 

```bash
aws cloudformation validate-template --template-body file://./template.yml
```

## Security
 
```bash
curl -X POST --user ':'  'https://democars.auth.us-east-2.amazoncognito.com/oauth2/token?grant_type=client_credentials&scope=profile' -H 'Content-Type: application/x-www-form-urlencoded'
```


## References

1. 
1. Directory structure :
- https://www.gobeyond.dev/packages-as-layers/ , https://www.gobeyond.dev/standard-package-layout/  and  https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091
- https://leonardqmarcq.com/posts/go-project-structure-for-api-gateway-lambda-with-aws-sam 
- https://github.com/golang-standards/project-layout 

https://changelog.com/posts/on-go-application-structure? 
2. Dynamo DB , https://www.trek10.com/blog/dynamodb-single-table-relational-modeling

3. Go Cloud https://github.com/google/go-cloud
4. AWS Lamdba Golang https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
5. AWS Lambda EnvVars https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html
6. https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-dynamodb-with-go-sdk.html