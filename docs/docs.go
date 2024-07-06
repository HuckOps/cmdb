// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/idc": {
            "get": {
                "description": "Get IDC List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "idc"
                ],
                "summary": "Get IDC List",
                "parameters": [
                    {
                        "type": "string",
                        "description": "IDC Code",
                        "name": "code",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "IDC Name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/restful.ResponseBody-model_PaginationResult-array_model_IDC"
                        }
                    }
                }
            },
            "post": {
                "description": "Create IDC",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "idc"
                ],
                "summary": "Create IDC",
                "parameters": [
                    {
                        "description": "IDC",
                        "name": "idc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.IDC"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/restful.ResponseBody-model_IDC"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "model.IDC": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "contact": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "locate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "orderSite": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                },
                "telephone": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "vendor": {
                    "description": "供应商信息",
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "model.PaginationResult-array_model_IDC": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IDC"
                    }
                }
            }
        },
        "restful.ResponseBody-model_IDC": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/model.IDC"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "restful.ResponseBody-model_PaginationResult-array_model_IDC": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/model.PaginationResult-array_model_IDC"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}