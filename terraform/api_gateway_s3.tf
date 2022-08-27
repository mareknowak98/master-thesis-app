resource "aws_api_gateway_rest_api" "mylearn_s3" {
  name        = "mylearn-s3-api"
  description = "REST API fort S3 related stuff"

  body = jsonencode(
    {
      "openapi" : "3.0.0",
      "info" : {
        "version" : "2016-10-21T17:26:28Z",
        "title" : "ApiName"
      },
      "paths" : {
        "/s3" : {
          "get" : {
            "parameters" : [
              {
                "name" : "key",
                "in" : "query",
                "required" : false,
                "schema" : {
                  "type" : "string"
                }
              }
            ],
            "responses" : {
              "200" : {
                "description" : "200 response",
                "content" : {
                  "application/json" : {
                    "schema" : {
                      "$ref" : "#/components/schemas/Empty"
                    }
                  }
                }
              },
              "500" : {
                "description" : "500 response"
              }
            },
            "x-amazon-apigateway-integration" : {
              "credentials" : aws_iam_role.s3_content_gateway.arn,
              "responses" : {
                "default" : {
                  "statusCode" : "500"
                },
                "2\\d{2}" : {
                  "statusCode" : "200"
                }
              },
              "requestParameters" : {
                "integration.request.path.key" : "method.request.querystring.key"
              },
              "uri" : "arn:aws:apigateway:eu-central-1:s3:path/{key}",
              "passthroughBehavior" : "when_no_match",
              "httpMethod" : "GET",
              "type" : "aws"
            }
          },
          "put" : {
            "parameters" : [
              {
                "name" : "key",
                "in" : "query",
                "required" : false,
                "schema" : {
                  "type" : "string"
                }
              }
            ],
            "responses" : {
              "200" : {
                "description" : "200 response",
                "content" : {
                  "application/json" : {
                    "schema" : {
                      "$ref" : "#/components/schemas/Empty"
                    }
                  },
                  "application/octet-stream" : {
                    "schema" : {
                      "$ref" : "#/components/schemas/Empty"
                    }
                  }
                }
              },
              "500" : {
                "description" : "500 response"
              }
            },
            "x-amazon-apigateway-integration" : {
              "credentials" : aws_iam_role.s3_content_gateway.arn,
              "responses" : {
                "default" : {
                  "statusCode" : "500"
                },
                "2\\d{2}" : {
                  "statusCode" : "200"
                }
              },
              "requestParameters" : {
                "integration.request.path.key" : "method.request.querystring.key"
              },
              "uri" : "arn:aws:apigateway:eu-central-1:s3:path/{key}",
              "passthroughBehavior" : "when_no_match",
              "httpMethod" : "PUT",
              "type" : "aws",
              "contentHandling" : "CONVERT_TO_BINARY"
            }
          }
        }
      },
      "x-amazon-apigateway-binary-media-types" : [
        "*/*"
      ],
      "components" : {
        "schemas" : {
          "Empty" : {
            "type" : "object",
            "title" : "Empty Schema"
          }
        },
#        "securitySchemes": {
#          "UserPool": {
#            "type": "apiKey",
#            "name": "Authorization",
#            "x-amazon-apigateway-authtype": "cognito_user_pools",
#            "x-amazon-apigateway-authorizer": {
#              "type": "cognito_user_pools",
#              "providerARNs": [
#                  aws_cognito_user_pool.mylearn.arn
#                ],
#              "identitySource": "$request.header.Authorization"
#            }
#          }
#        }
        "securitySchemes": {
          "jwt-authorizer-oauth": {
            "type": "oauth2",
            "x-amazon-apigateway-authorizer": {
              "type": "jwt",
              "jwtConfiguration": {
                "issuer": "https://cognito-idp.region.amazonaws.com/userPoolId",
                "audience": [
                  "audience1",
                  "audience2"
                ]
              },
              "identitySource": "$request.header.Authorization"
            }
          }
        }
      }
    }
  )

}


## API deployment
resource "aws_api_gateway_stage" "mylearn_s3" {
  deployment_id = aws_api_gateway_deployment.mylearn_s3.id
  rest_api_id   = aws_api_gateway_rest_api.mylearn_s3.id
  stage_name    = var.deployment
}

##
resource "aws_api_gateway_deployment" "mylearn_s3" {
  rest_api_id = aws_api_gateway_rest_api.mylearn_s3.id

  lifecycle {
    create_before_destroy = true
  }

  depends_on = [
    data.aws_api_gateway_resource.s3,
    aws_api_gateway_rest_api.mylearn_s3
  ]
}

data "aws_api_gateway_resource" "s3" {
  rest_api_id = aws_api_gateway_rest_api.mylearn_s3.id
  path        = "/s3"
}

module "cors_s3_1" {
  source  = "squidfunk/api-gateway-enable-cors/aws"
  version = "0.3.3"

  api_id          = aws_api_gateway_rest_api.mylearn_s3.id
  api_resource_id = data.aws_api_gateway_resource.s3.id
}