// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/": {
            "post": {
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "summary": "Create new short url",
                "parameters": [
                    {
                        "description": "Full url for shortening",
                        "name": "url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "http://shortener.org/123",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid url: http//example",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "http://shortener.org/123",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/shorten": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new short url",
                "parameters": [
                    {
                        "description": "Full url for shortening",
                        "name": "fullURL",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateShortRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "URL saved, short url returned in result field",
                        "schema": {
                            "$ref": "#/definitions/responses.CreateShortResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Full URL already added earlier, old short url is returned in result field",
                        "schema": {
                            "$ref": "#/definitions/responses.CreateShortResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/shorten/batch": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create many new short urls",
                "parameters": [
                    {
                        "description": "List of full urls for shortening",
                        "name": "fullURList",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/requests.CreateShortBatchItem"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.CreateShortBatchResponseItem"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/user/urls": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get all urls user shortened",
                "responses": {
                    "200": {
                        "description": "List of urls, user added",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.ListShortItem"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content. User not added any urls yet"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete urls user shortened earlier",
                "parameters": [
                    {
                        "description": "List of urls to delete",
                        "name": "deleteURLs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Delete request accepted and put on queue, urls will be deleted eventually"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "summary": "returns \"OK\" if service is working and storage is available",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "DB unavailable",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "summary": "Redirects to the full url, if found in storage by {id}",
                "parameters": [
                    {
                        "type": "string",
                        "description": "url id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "307": {
                        "description": "redirects to full url",
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "http://example.com/"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "410": {
                        "description": "Record has been deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "requests.CreateShortBatchItem": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "type": "string"
                },
                "original_url": {
                    "type": "string"
                }
            }
        },
        "requests.CreateShortRequest": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string",
                    "example": "http://example.com/asd"
                }
            }
        },
        "responses.CreateShortBatchResponseItem": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "type": "string"
                },
                "short_url": {
                    "type": "string"
                }
            }
        },
        "responses.CreateShortResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string",
                    "example": "http://shortener.org/123"
                }
            }
        },
        "responses.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Not found"
                }
            }
        },
        "responses.ListShortItem": {
            "type": "object",
            "properties": {
                "original_url": {
                    "type": "string",
                    "example": "http://example.com/"
                },
                "short_url": {
                    "type": "string",
                    "example": "http://shortener.org/123"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Shortener API",
	Description:      "API server for shorting log urls to short ones",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}