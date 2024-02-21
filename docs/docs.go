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
        "/events": {
            "get": {
                "description": "Retrieves a list of all events with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Get all events",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page for pagination",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of events",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.EventListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/events/test-store-all": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Store all events from the json file in the frontend repo",
                "responses": {}
            }
        },
        "/events/{id}": {
            "get": {
                "description": "Retrieves an event by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Get an event by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.EventResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates an event by ID with new details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Update an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Event Update Information",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.CreateEventRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Event successfully updated",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.EventResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid input",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/line-login/auth": {
            "get": {
                "description": "Redirects the user to LINE's OAuth service for authentication.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OAuth"
                ],
                "summary": "Initiate LINE OAuth login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Redirect path after login",
                        "name": "redirect_path",
                        "in": "query"
                    }
                ],
                "responses": {
                    "301": {
                        "description": "Redirect to the target URL",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/line-login/callback": {
            "get": {
                "description": "Handles the callback from LINE's OAuth service and redirects the user to the frontend with the tokens in the query and cookies.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OAuth"
                ],
                "summary": "Handle LINE OAuth callback",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization code",
                        "name": "code",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "State",
                        "name": "state",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "301": {
                        "description": "Redirect to the frontend",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Retrieves a list of users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Users",
                "responses": {
                    "200": {
                        "description": "List of users successfully retrieved",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.UserListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/users/events": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves a list of events associated with a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User Events",
                "responses": {
                    "200": {
                        "description": "List of events associated with the user",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.EventListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Subscribes a user to an event with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Subscribe to an event",
                "parameters": [
                    {
                        "description": "Event Subscription Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.CreateEventRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully subscribed to the event",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.EventResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid input, such as invalid time format or missing required fields",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict - User already subscribed to the event",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity - User has exceeded the maximum number of subscriptions",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Error storing the event, subscribing the user, or enqueuing the event notification",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/users/events/all": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Subscribes a user to all events",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Subscribe to all events, and if remind is true, it will send all the event notification to user immediately",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "Send the Line notification immediately",
                        "name": "remind",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully subscribed to the event",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.EventResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid input",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/users/events/delete/{event_id}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deletes a specific event by its ID for a given user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "event_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Event successfully deleted",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/users/events/{event_id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deletes a specific event by its ID for a given user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "event_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Event successfully deleted",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/users/login/{user_id}": {
            "get": {
                "description": "Simulates a login process for a user by generating fake access and refresh tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Fake Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful, tokens generated",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.TokenResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/users/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Logs out a user by invalidating their authentication token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User logout",
                "responses": {
                    "200": {
                        "description": "Logout successful",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized: Invalid token format",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/users/profile": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Fetches the profile of a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Profile",
                "responses": {
                    "200": {
                        "description": "Profile successfully retrieved",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized: Invalid or expired token",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/users/refresh_token": {
            "get": {
                "description": "Refreshes the access and refresh tokens for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Refresh User Token",
                "responses": {
                    "200": {
                        "description": "Access and Refresh Tokens successfully generated",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid request format",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Invalid or expired refresh token",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Error generating tokens",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Register a fake user for testing purposes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Fake Register",
                "parameters": [
                    {
                        "description": "Create Fake User Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.CreateFakeUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Fake register successful",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        },
        "/users/{user_id}": {
            "get": {
                "description": "Retrieves a user's information by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully retrieved",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.UserResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bikefest_pkg_model.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "bikefest_pkg_model.CreateEventRequest": {
            "type": "object",
            "properties": {
                "event_detail": {
                    "type": "string",
                    "example": "{\"title\":\"test event\",\"description\":\"test event description\"}"
                },
                "event_time_end": {
                    "type": "string",
                    "example": "2021/01/01 00:00"
                },
                "event_time_start": {
                    "type": "string",
                    "example": "2021/01/01 00:00"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "bikefest_pkg_model.CreateFakeUserRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "bikefest_pkg_model.Event": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "event_detail": {
                    "description": "the ` + "`" + `EventDetail` + "`" + ` field store the event detail in json format, this would be parsed when send to line message API",
                    "type": "string"
                },
                "event_time_end": {
                    "type": "string"
                },
                "event_time_start": {
                    "type": "string"
                },
                "id": {
                    "description": "the event id is defne at the frontend, if frontend don't have event id, the event id would be calculated by the hash of event detail and event time",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "bikefest_pkg_model.EventListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/bikefest_pkg_model.Event"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "bikefest_pkg_model.EventResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/bikefest_pkg_model.Event"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "bikefest_pkg_model.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "bikefest_pkg_model.Token": {
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
        "bikefest_pkg_model.TokenResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/bikefest_pkg_model.Token"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "bikefest_pkg_model.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/bikefest_pkg_model.Event"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "picture_url": {
                    "description": "TODO: add more user info for line login and line message API identity",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "bikefest_pkg_model.UserListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/bikefest_pkg_model.User"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "bikefest_pkg_model.UserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/bikefest_pkg_model.User"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
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
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
