// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplatemanager = `{
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
        "/health-check": {
            "get": {
                "description": "get service health check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthCheck"
                ],
                "summary": "Show health check",
                "responses": {}
            }
        },
        "/projects": {
            "get": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "description": "List projects",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "List projects",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Last token fetched",
                        "name": "last_token_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Per page count",
                        "name": "per_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/projectparam.ListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "description": "Create project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Create project",
                "parameters": [
                    {
                        "description": "Create project request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/projectparam.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/projectparam.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}": {
            "post": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "description": "Update project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Update project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project identifier",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update project request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/projectparam.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/projectparam.UpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "description": "Delete project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Delete project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project identifier",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/projectparam.DeleteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/sources": {
            "get": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "description": "List sources",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Source"
                ],
                "summary": "List sources",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Last token fetched",
                        "name": "last_token_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Per page count",
                        "name": "per_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sourceparam.ListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "description": "Create source",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Source"
                ],
                "summary": "Create source",
                "parameters": [
                    {
                        "description": "Create source request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sourceparam.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/sourceparam.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/param.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/param.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Register request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/param.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/param.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Project": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted": {
                    "type": "boolean"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "token_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.Source": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deleted": {
                    "type": "boolean"
                },
                "deletedAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "metadata": {
                    "$ref": "#/definitions/entity.SourceMetadata"
                },
                "name": {
                    "type": "string"
                },
                "ownerID": {
                    "type": "string"
                },
                "projectID": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/entity.Status"
                },
                "updatedAt": {
                    "type": "string"
                },
                "writeKey": {
                    "type": "string"
                }
            }
        },
        "entity.SourceMetadata": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                }
            }
        },
        "entity.Status": {
            "type": "string",
            "enum": [
                "active",
                "not active"
            ],
            "x-enum-varnames": [
                "SourceStatusActive",
                "SourceStatusNotActive"
            ]
        },
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "param.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "name@test.com"
                },
                "password": {
                    "type": "string",
                    "example": "123Qwe!@#"
                }
            }
        },
        "param.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "$ref": "#/definitions/param.Token"
                },
                "user": {
                    "$ref": "#/definitions/param.UserInfo"
                }
            }
        },
        "param.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "name@test.com"
                },
                "name": {
                    "type": "string",
                    "example": "name"
                },
                "password": {
                    "type": "string",
                    "example": "123Qwe!@#"
                }
            }
        },
        "param.RegisterResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "name@test.com"
                },
                "id": {
                    "type": "string",
                    "example": "f90631e0-aad3-4eb1-8cef-1478711e16e9"
                }
            }
        },
        "param.Token": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "param.UserInfo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "projectparam.CreateRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "description"
                },
                "name": {
                    "type": "string",
                    "example": "name"
                }
            }
        },
        "projectparam.CreateResponse": {
            "type": "object",
            "properties": {
                "project": {
                    "$ref": "#/definitions/entity.Project"
                }
            }
        },
        "projectparam.DeleteResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "projectparam.ListResponse": {
            "type": "object",
            "properties": {
                "has_more": {
                    "type": "boolean"
                },
                "last_token": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Project"
                    }
                }
            }
        },
        "projectparam.UpdateRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "description"
                },
                "name": {
                    "type": "string",
                    "example": "name"
                }
            }
        },
        "projectparam.UpdateResponse": {
            "type": "object",
            "properties": {
                "project": {
                    "$ref": "#/definitions/entity.Project"
                }
            }
        },
        "sourceparam.CreateRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "description"
                },
                "name": {
                    "type": "string",
                    "example": "test"
                },
                "project_id": {
                    "type": "string"
                }
            }
        },
        "sourceparam.CreateResponse": {
            "type": "object",
            "properties": {
                "source": {
                    "$ref": "#/definitions/entity.Source"
                }
            }
        },
        "sourceparam.ListResponse": {
            "type": "object",
            "properties": {
                "has_more": {
                    "type": "boolean"
                },
                "last_token": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "sources": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Source"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "JWTToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfomanager holds exported Swagger Info so clients can modify it
var SwaggerInfomanager = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "manager",
	SwaggerTemplate:  docTemplatemanager,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfomanager.InstanceName(), SwaggerInfomanager)
}
