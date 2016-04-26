variable "apex_function_uppercase" {}

resource "aws_api_gateway_rest_api" "Uppercase" {
  name = "Uppercase API"
  description = "Value to uppercase"
}

resource "aws_api_gateway_resource" "Uppercase" {
  rest_api_id = "${aws_api_gateway_rest_api.Uppercase.id}"
  parent_id = "${aws_api_gateway_rest_api.Uppercase.root_resource_id}"
  path_part = "uppercase"
}

resource "aws_api_gateway_method" "UppercasePost" {
  rest_api_id = "${aws_api_gateway_rest_api.Uppercase.id}"
  resource_id = "${aws_api_gateway_resource.Uppercase.id}"
  http_method = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_method_response" "200" {
  rest_api_id = "${aws_api_gateway_rest_api.Uppercase.id}"
  resource_id = "${aws_api_gateway_resource.Uppercase.id}"
  http_method = "${aws_api_gateway_method.UppercasePost.http_method}"
  status_code = "200"
  response_models = {
    "application/json" = "Empty"
  }
}

resource "aws_api_gateway_integration_response" "Uppercase" {
  rest_api_id = "${aws_api_gateway_rest_api.Uppercase.id}"
  resource_id = "${aws_api_gateway_resource.Uppercase.id}"
  http_method = "${aws_api_gateway_method.UppercasePost.http_method}"
  status_code = "${aws_api_gateway_method_response.200.status_code}"
}

resource "aws_api_gateway_integration" "UppercasePost" {
  rest_api_id = "${aws_api_gateway_rest_api.Uppercase.id}"
  resource_id = "${aws_api_gateway_resource.Uppercase.id}"
  http_method = "${aws_api_gateway_method.UppercasePost.http_method}"
  type = "AWS"
  integration_http_method = "POST" # Must be POST for invoking Lambda function
  credentials = "${aws_iam_role.gateway_invoke_lambda.arn}"
  # http://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#uri
  uri = "arn:aws:apigateway:${var.aws_region}:lambda:path/2015-03-31/functions/${var.apex_function_uppercase}/invocations"
}

resource "aws_api_gateway_deployment" "UppercaseDeployment" {
  depends_on = ["aws_api_gateway_integration.UppercasePost"]

  rest_api_id = "${aws_api_gateway_rest_api.Uppercase.id}"
  stage_name = "dev"
}