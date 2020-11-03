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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
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
        "/algo_centrality/{node}/{relationship}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "algorithms"
                ],
                "summary": "Show centrality nodes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Node type",
                        "name": "node",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Relationship projection",
                        "name": "relationship",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/algo_community/{node}/{relationship}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "algorithms"
                ],
                "summary": "Show comunity nodes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Node type",
                        "name": "node",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Relationship projection",
                        "name": "relationship",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/algo_pagerank/{node}/{relationship}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "algorithms"
                ],
                "summary": "Show page rank nodes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Node type",
                        "name": "node",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Relationship projection",
                        "name": "relationship",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/algo_path/{node}/{relationship}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "algorithms"
                ],
                "summary": "Show path nodes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Node type",
                        "name": "node",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Relationship projection",
                        "name": "relationship",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/movie": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "Show all movies",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/movie/": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "Create a movies",
                "parameters": [
                    {
                        "description": "Movie model",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/movie/{title}/{released}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "Get movie by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie name",
                        "name": "title",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Movie released year",
                        "name": "released",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "Edit a movies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie name",
                        "name": "title",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Movie released year",
                        "name": "released",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Movie model",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "Delete a movie",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie name",
                        "name": "title",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Movie released year",
                        "name": "released",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/person": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Show all persons",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/person/": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Create a person",
                "parameters": [
                    {
                        "description": "Person model",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Person"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/person/GETwith_relationship": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Get person by relationship",
                "parameters": [
                    {
                        "description": "name Person",
                        "name": "pKp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PersonKNOWSPerson"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/person/with_relationship": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Create a person with relationship",
                "parameters": [
                    {
                        "description": "Relationship model",
                        "name": "pKp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PersonKNOWSPerson"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/person/{name}/{born}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Get person by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name Person",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "year person born",
                        "name": "born",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Edit a person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name Person",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "year person born",
                        "name": "born",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Person model",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Person"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "Delete a person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name Person",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "year person born",
                        "name": "born",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Movie": {
            "type": "object",
            "properties": {
                "released": {
                    "type": "integer"
                },
                "tagline": {
                    "type": "string"
                },
                "title": {
                    "description": "gorm.Model",
                    "type": "string"
                }
            }
        },
        "model.Person": {
            "type": "object",
            "properties": {
                "born": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.PersonKNOWSPerson": {
            "type": "object",
            "properties": {
                "ano": {
                    "description": "RelationType string ` + "`" + `json:\"relationship_type\"` + "`" + `",
                    "type": "string"
                },
                "mes": {
                    "type": "string"
                },
                "p1": {
                    "$ref": "#/definitions/model.Person"
                },
                "p2": {
                    "$ref": "#/definitions/model.Person"
                }
            }
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
	Host:        "localhost:3000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "fiber aplication with neo4j",
	Description: "This is a sample swagger for Fiber",
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
