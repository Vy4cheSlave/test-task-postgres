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
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/bash/create-command": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/bash/"
                ],
                "parameters": [
                    {
                        "description": "input bash string",
                        "name": "new_command",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Vy4cheSlave_test-task-postgres_bash.ReqCreateNewCommandBody"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/bash/get-commands": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/bash/"
                ],
                "responses": {}
            }
        },
        "/bash/get-commands/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/bash/"
                ],
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "uint without 0",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "github_com_Vy4cheSlave_test-task-postgres_bash.ReqCreateNewCommandBody": {
            "type": "object",
            "properties": {
                "bash_strings": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "bash API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
