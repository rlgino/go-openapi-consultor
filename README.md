# Golang API with OpenAPI and AWS Lambda

## Steps

### Create Golang API and AWS layer
We need to implement the AWS Proxy event request and response:
```go
func (h Handler) HandleRequest(_ context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)
```
Add API_TOKEN and API_ORG variables into Lambda

### Add API Gateway
1. Create new REST API
2. Name and regional endpoint
3. Create new resource, just set name and url endpoint
4. Select resource then go to `actions` -> `create method`
5. When create method select integration Lambda, check use lambda proxy integration select our lambda then save

## Deploy API Gateway
1. Go to `actions` --> `deploy API`
2. Select new stage with name test (or whatever you want)

## Test API
```json
{
  "Value": "Cual es el mejor equipo de futbol de Argentina?"
}
```

### Create Frontend

### Resources:
* [Writing and deploying your first AWS Lambda with Go](https://levelup.gitconnected.com/writing-and-deploying-your-first-aws-lambda-7a6d5800b443)
* [Consume Open AI (ChatGPT) With Golang](https://betterprogramming.pub/chatgpt-golang-a7bd524e7563)
