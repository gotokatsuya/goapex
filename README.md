# goapex

## Setup

```bash
$ curl https://raw.githubusercontent.com/apex/apex/master/install.sh | sh
$ apex upgrade

$ go get github.com/apex/go-apex
```

Attach `AWSLambdaFullAccess` policy to a user.

```bash
$ brew install direnv
$ cd PATH/goapex
$ vim .envrc

export AWS_ACCESS_KEY="hoge-access_key"
export AWS_SECRET_KEY="hoge-secret_key"
export AWS_REGION="hoge-region"

$ direnv allow
```

Attach `AWSLambdaBasicExecutionRole` policy to a role.

```bash
$ cd PATH/goapex
$ vim project.json
{
  "name": "goapex",
  "description": "Sample Go project using apex",
  "runtime": "golang",
  "memory": 128,
  "timeout": 5,
  "role": "arn:aws:iam::<aws_account_id>:role/lambda_basic_execution",
  "defaultEnvironment": "dev",
  "environment": {},
  "nameTemplate": "{{.Project.Name}}_{{.Function.Name}}",
  "retainedVersions": 10,
  "vpc": {
  	"securityGroups": [],
  	"subnets": []
  }
}
```

## Deploy

```bash
$ apex deploy uppercase --dry-run

  + function goapex_uppercase
    runtime: nodejs
    memory: 128
    timeout: 5
    handler: index.handle

  + alias goapex_uppercase
    alias: current
    version: 1


$ apex deploy uppercase

   • creating function         function=uppercase
   • created alias current     function=uppercase version=1
   • function created          function=uppercase name=goapex_uppercase version=1
```

## List

```bash
$ apex list

  uppercase
    description: Uppercase
    runtime: nodejs
    memory: 128mb
    timeout: 5s
    role: arn:aws:iam::<aws_account_id>:role/lambda_basic_execution
    handler: index.handle
    current version: 2


$ apex list --tfvars

apex_function_uppercase="arn:aws:lambda:<aws_region>:<aws_account_id>:function:goapex_uppercase"
```

## Invole

```bash
$ vim event.json
{
  {"value":"eureka"}
}

$ apex invoke uppercase < event.json

{"value":"EUREKA"}
```

## Log

```bash
$ apex logs uppercase -s 10m

/aws/lambda/goapex_uppercase START RequestId: xxx 
Version: 2

/aws/lambda/goapex_uppercase END RequestId: xxx

/aws/lambda/goapex_uppercase REPORT RequestId: xxx	
Duration: 0.96 ms	Billed Duration: 100 ms 	Memory Size: 128 MB	Max Memory Used: 31 MB	
```

## Metrics

```bash
$ apex metrics

  uppercase
    invocations: 3
    duration: 43ms
    throttles: 0
    error: 0
```

## Delete

```bash
$ apex delete -f uppercase

   • deleting                  function=uppercase
   • function deleted          function=uppercase
```


## API Gateway

```
$ brew cask install terraform
$ apex init (create infrastructure dir)
$ apex infra plan
$ apex infra apply
$ apex deploy uppercase
```


## VPC & Database

Launch rds for aurora.
Setting vpc.

```json
{
  "vpc": {
    "securityGroups": ["hoge"],
    "subnets": ["hoge","hoge"]
  }
}
```

Connection and create database.
```bash
$ mysql -h xxx.xxx.xxx.rds.amazonaws.com -P 3306 -u username -p
```


