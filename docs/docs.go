// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Hossein Yazdani",
            "url": "https://GodlyNice.ir"
        },
        "license": {
            "name": "MIT license",
            "url": "https://opensource.org/license/mit/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "post": {
                "description": "creates a new user and returns the newly created data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user management"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structure.UserCreationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "single object containing user's data",
                        "schema": {
                            "$ref": "#/definitions/doc.createNewUserHandlerResponse"
                        }
                    },
                    "400": {
                        "description": "bad input",
                        "schema": {
                            "$ref": "#/definitions/doc.httpResponseError"
                        }
                    },
                    "500": {
                        "description": "unexpected error",
                        "schema": {
                            "$ref": "#/definitions/doc.httpResponseError"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "get a single user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user management"
                ],
                "summary": "Get user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "single object containing user's data",
                        "schema": {
                            "$ref": "#/definitions/doc.getUserByIDHandlerResponse"
                        }
                    },
                    "400": {
                        "description": "invalid ID",
                        "schema": {
                            "$ref": "#/definitions/doc.httpResponseError"
                        }
                    },
                    "404": {
                        "description": "user not found",
                        "schema": {
                            "$ref": "#/definitions/doc.httpResponseError"
                        }
                    },
                    "500": {
                        "description": "unexpected error",
                        "schema": {
                            "$ref": "#/definitions/doc.httpResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "doc.createNewUserHandlerResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/structure.User"
                }
            }
        },
        "doc.getUserByIDHandlerResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/structure.User"
                }
            }
        },
        "doc.httpResponseError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "error message"
                }
            }
        },
        "structure.Order": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/structure.Product"
                    }
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "structure.Product": {
            "type": "object",
            "properties": {
                "available": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "rating_count": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "structure.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@email.com"
                },
                "firstname": {
                    "type": "string",
                    "example": "John"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "lastname": {
                    "type": "string",
                    "example": "Doe"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/structure.Order"
                    }
                },
                "phone": {
                    "type": "string",
                    "example": "0123456789"
                },
                "username": {
                    "type": "string",
                    "example": "user"
                }
            }
        },
        "structure.UserCreationRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@email.com"
                },
                "firstname": {
                    "type": "string",
                    "example": "John"
                },
                "lastname": {
                    "type": "string",
                    "example": "Doe"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "type": "string",
                    "example": "0123456789"
                },
                "username": {
                    "type": "string",
                    "example": "user"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Kala",
	Description:      "Online shop backend.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
