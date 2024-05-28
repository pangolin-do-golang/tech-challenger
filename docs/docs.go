// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cart/add-product": {
            "post": {
                "description": "Adds a Product to Customer's Cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "parameters": [
                    {
                        "description": "AddProductPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.AddProductPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "{\\\"error\\\": \\\"Internal Server Error\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/cart/edit-product": {
            "post": {
                "description": "Edits a Product from Customer's Cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "parameters": [
                    {
                        "description": "EditProductPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.EditProductPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "{\\\"error\\\": \\\"Internal Server Error\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/cart/remove-product": {
            "post": {
                "description": "Removes a Product from Customer's Cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "parameters": [
                    {
                        "description": "RemoveProductPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RemoveProductPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/customer": {
            "get": {
                "description": "Get all customer's list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Get customer list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/customer.Customer"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Create customer",
                "parameters": [
                    {
                        "description": "CustomerPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CustomerPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "customer.Customer",
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\": \\\"Invalid CPF\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/customer/{cpf}": {
            "get": {
                "description": "Get a customer by cpf",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Get customer by cpf",
                "parameters": [
                    {
                        "type": "string",
                        "description": "customer cpf",
                        "name": "cpf",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    },
                    "404": {
                        "description": "Customer not found"
                    }
                }
            }
        },
        "/customer/{id}": {
            "put": {
                "description": "Update a customer by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Update customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the customer",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "CustomerPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CustomerPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "customer.Customer",
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\": \\\"Invalid CPF\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a customer by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Delete customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "123456789",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Invalid identifier informed"
                    }
                }
            }
        },
        "/order": {
            "get": {
                "description": "Get all order's list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get order list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/order.Order"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create order from Cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Create order from Cart",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.Order"
                        }
                    },
                    "500": {
                        "description": "{\\\"error\\\": \\Internal Server Error\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/order/{id}": {
            "get": {
                "description": "Get a order by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get order by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the order",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.Order"
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\": \\Invalid identifier informed\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.AddProductPayload": {
            "type": "object",
            "required": [
                "client_id",
                "product_id",
                "quantity"
            ],
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "comments": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                },
                "quantity": {
                    "description": "TODO validação de quantidade \u003e= 0 / estoque",
                    "type": "integer"
                }
            }
        },
        "controller.CustomerPayload": {
            "type": "object",
            "required": [
                "cpf",
                "email",
                "name"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 120,
                    "minimum": 18
                },
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 5
                }
            }
        },
        "controller.EditProductPayload": {
            "type": "object",
            "required": [
                "client_id",
                "product_id",
                "quantity"
            ],
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "comments": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "controller.RemoveProductPayload": {
            "type": "object",
            "required": [
                "client_id",
                "product_id"
            ],
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "customer.Customer": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "order.Order": {
            "type": "object",
            "properties": {
                "clientID": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "totalAmount": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Tech Challenge Food API",
	Description:      "Fast Food API for FIAP Tech course",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
