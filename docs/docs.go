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
        "/episodes": {
            "get": {
                "description": "get all episodes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "episodes"
                ],
                "summary": "List episodes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/deuna-rickandmorty-api_internal_episode.Episode"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/episodes/multiple": {
            "get": {
                "description": "get multiple episodes by ids",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "episodes"
                ],
                "summary": "Get multiple episodes by ids",
                "parameters": [
                    {
                        "type": "string",
                        "description": "episode ids delimited with comma ,",
                        "name": "ids",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/deuna-rickandmorty-api_internal_episode.Episode"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/episodes/{id}": {
            "get": {
                "description": "get episode by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "episodes"
                ],
                "summary": "Get episode by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "episode id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/deuna-rickandmorty-api_internal_episode.Episode"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthCheck"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "deuna-rickandmorty-api_internal_episode.Episode": {
            "type": "object",
            "properties": {
                "air_date": {
                    "type": "string"
                },
                "characters": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "created": {
                    "type": "string"
                },
                "episode": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/deuna-rickandmorty-api/v1",
	Schemes:          []string{"http"},
	Title:            "Rick and Morty Api",
	Description:      "This is a kata api server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
