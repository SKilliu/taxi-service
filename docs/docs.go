// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/operators/assign": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Assign the car to the driver",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "operators"
                ],
                "summary": "Assign car to the driver",
                "parameters": [
                    {
                        "description": "Body for assign car",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AssignCarToDriverReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/operators/car": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Add new car to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "operators"
                ],
                "summary": "Add new car",
                "parameters": [
                    {
                        "description": "Body for add new car request",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/AddCarReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/operators/user": {
            "post": {
                "description": "Create a new user by operator",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "operators"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "Body for creation of a new user",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/SignUpReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Create a new order as a client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get available orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetAvailableOrdersResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Create a new order as a client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create order",
                "parameters": [
                    {
                        "description": "Body for creating new order",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateOrderReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Accept or close the order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Actions with order",
                "parameters": [
                    {
                        "description": "Body for order actions",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OrderActionsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/sign_in": {
            "post": {
                "description": "Sign in with login and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "Body for sign in",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/SignInReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/SignUpResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/sign_up": {
            "post": {
                "description": "Sign up with login, password and account type (driver, client or operator)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "Body for sign up",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/SignUpReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/SignUpResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Get your profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetProfileResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Edit user's name and email in profile info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Edit profile",
                "parameters": [
                    {
                        "description": "Body for edit profile request",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/EditProfileReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Delete user profile from db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "AddCarReq": {
            "type": "object",
            "properties": {
                "model": {
                    "type": "string",
                    "example": "BMW"
                },
                "number": {
                    "type": "string",
                    "example": "AX1234XA"
                },
                "series": {
                    "type": "string",
                    "example": "M5"
                },
                "status": {
                    "type": "string",
                    "example": "available"
                }
            }
        },
        "CreateOrderReq": {
            "type": "object",
            "properties": {
                "car_arrival_time": {
                    "type": "string",
                    "example": "2020-11-11T23:30:00Z"
                },
                "destination_point": {
                    "type": "object",
                    "$ref": "#/definitions/Location"
                },
                "starting_point": {
                    "type": "object",
                    "$ref": "#/definitions/Location"
                }
            }
        },
        "EditProfileReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "new-email@example.com"
                },
                "name": {
                    "type": "string",
                    "example": "Tester"
                }
            }
        },
        "ErrResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "message": {
                    "type": "string",
                    "example": "INTERNAL_SERVER_ERROR"
                }
            }
        },
        "GetAvailableOrdersResp": {
            "type": "object",
            "properties": {
                "car_arrival_time": {
                    "type": "string"
                },
                "client_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "trip_info": {
                    "type": "object",
                    "$ref": "#/definitions/TripInfo"
                }
            }
        },
        "GetProfileResp": {
            "type": "object",
            "properties": {
                "account_type": {
                    "type": "string",
                    "example": "operator"
                },
                "email": {
                    "type": "string",
                    "example": "test@example.com"
                },
                "id": {
                    "type": "string",
                    "example": "Yh34te-saaiud3322chadsc-asdvcsf"
                },
                "name": {
                    "type": "string",
                    "example": "Tester"
                },
                "profile_image_url": {
                    "type": "string",
                    "example": "http://simple-service-backend/simple-service/photo-924y82hde7ce.jpg"
                }
            }
        },
        "Location": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "Some City, some Street, 123"
                },
                "destination_point_longitude": {
                    "type": "number",
                    "example": 12.12345
                },
                "latitude": {
                    "type": "number",
                    "example": 12.12345
                }
            }
        },
        "SignInReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "qwerty1234"
                }
            }
        },
        "SignUpReq": {
            "type": "object",
            "properties": {
                "account_type": {
                    "type": "string",
                    "example": "client"
                },
                "email": {
                    "type": "string",
                    "example": "test@example.com"
                },
                "name": {
                    "type": "string",
                    "example": "TestName"
                },
                "password": {
                    "type": "string",
                    "example": "qwerty1234"
                }
            }
        },
        "SignUpResp": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "nausdgtGTGAjndfsKijIYbsgfsuadfe34r"
                }
            }
        },
        "TripInfo": {
            "type": "object",
            "properties": {
                "destination_point": {
                    "type": "object",
                    "$ref": "#/definitions/Location"
                },
                "distance": {
                    "type": "number",
                    "example": 15
                },
                "starting_point": {
                    "type": "object",
                    "$ref": "#/definitions/Location"
                }
            }
        },
        "dto.AssignCarToDriverReq": {
            "type": "object",
            "properties": {
                "car_id": {
                    "type": "string"
                },
                "driver_id": {
                    "type": "string"
                }
            }
        },
        "dto.OrderActionsReq": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Taxi-service",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
