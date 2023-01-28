// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/demo": {
            "get": {
                "description": "向你说Hello",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "测试"
                ],
                "summary": "测试SayHello",
                "parameters": [
                    {
                        "type": "string",
                        "description": "人名",
                        "name": "who",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"hello sos\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"who are you\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/file": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "随便",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "文件",
                        "name": "other",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"hello sos\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"who are you\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "127.0.0.1:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}