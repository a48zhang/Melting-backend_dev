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
        "contact": {
            "name": "@a48zhang \u0026 @Cg1028",
            "email": "3557695455@qq.com 2194028175@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/join": {
            "get": {
                "description": "Join a proposal with infoId",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Join certain proposal (login required)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "infoId",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    },
                    "400": {
                        "description": "user exists",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "login and return id\u0026token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "register and login"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "the User who is logging in",
                        "name": "loginAuth",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    },
                    {
                        "type": "string",
                        "description": "type of login(use 'qq' to login with qq)",
                        "name": "loginType",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.loginResponse"
                        }
                    },
                    "401": {
                        "description": "username or password incorrect",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "403": {
                        "description": "param not satisfied",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "token generation failed",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/project": {
            "get": {
                "description": "Get a project with its id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Get a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the id of the project",
                        "name": "info_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.ProposalInfo"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update user's project(login required)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Update one's project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the id of the project",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "operating project",
                        "name": "newproject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.ProposalInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user's project(login required)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Delete one's project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the id of the project",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/project/games": {
            "get": {
                "description": "Get a game's info by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Get a game's info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "game_id",
                        "name": "game_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/project/games/details": {
            "get": {
                "description": "Get a game's detail by id",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Get a game's detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "game_id",
                        "name": "game_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/project/games/find": {
            "post": {
                "description": "Get some games' info with certain circumstances",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Get some games' info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "game circumstances",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.Game"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/project/newproject": {
            "post": {
                "description": "Create user's project(login required)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Create a new project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "operating project",
                        "name": "newproject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.ProposalInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/project/template": {
            "get": {
                "description": "Get a template with its id or name",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "template"
                ],
                "summary": "Get a template",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the id of the template",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "the name of the template",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Template"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Create user's project(login required)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "template"
                ],
                "summary": "Update certain template",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "operating project",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.Template"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new template(login required)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "template"
                ],
                "summary": "Create a new template",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "operating project",
                        "name": "newproject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.Template"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Create user's project(login required)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "template"
                ],
                "summary": "Delete certain template",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "operating project",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "create a new account in certain way",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "register and login"
                ],
                "summary": "register account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the type of the register",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "description": "register data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    },
                    "400": {
                        "description": "param not satisfied",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Get User's info with its userID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get User's info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uid",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get User's info",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get User's info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    }
                }
            },
            "put": {
                "description": "upload user's profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "upload profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "new user profile",
                        "name": "Profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "update failed",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/users/myproject": {
            "get": {
                "description": "Get all the projects from user(login required)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Get one's projects",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.ProposalInfo"
                        }
                    },
                    "404": {
                        "description": "Resource requested not found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/users/photo": {
            "put": {
                "description": "upload user's avatar",
                "consumes": [
                    "image/jpeg"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "upload photo",
                "parameters": [
                    {
                        "type": "object",
                        "description": "the photo of the user",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "file not received",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "failed to upload file",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "db.Game": {
            "type": "object",
            "properties": {
                "crowd": {
                    "type": "string"
                },
                "gameid": {
                    "type": "integer"
                },
                "gamename": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "venue": {
                    "type": "string"
                }
            }
        },
        "db.ProposalInfo": {
            "type": "object",
            "properties": {
                "aim": {
                    "description": "活动目的",
                    "type": "string"
                },
                "budget": {
                    "description": "活动预算",
                    "type": "string"
                },
                "corporates": {
                    "type": "string"
                },
                "department": {
                    "description": "部门",
                    "type": "string"
                },
                "game": {
                    "description": "游戏项目",
                    "type": "string"
                },
                "info_id": {
                    "description": "活动序号",
                    "type": "integer"
                },
                "name": {
                    "description": "活动",
                    "type": "string"
                },
                "nodes": {
                    "description": "项目环节",
                    "type": "string"
                },
                "optional_tables": {
                    "description": "可选标签",
                    "type": "string"
                },
                "place": {
                    "description": "活动位置",
                    "type": "string"
                },
                "time": {
                    "description": "活动时间",
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                }
            }
        },
        "db.Template": {
            "type": "object",
            "properties": {
                "context": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "temid": {
                    "type": "integer"
                }
            }
        },
        "db.User": {
            "type": "object",
            "properties": {
                "auth": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "muxipass": {
                    "type": "integer"
                },
                "nick_name": {
                    "description": "最多七个汉字",
                    "type": "string"
                },
                "photo": {
                    "description": "头像",
                    "type": "string"
                },
                "position": {
                    "description": "职位",
                    "type": "string"
                },
                "qq": {
                    "type": "string"
                },
                "studentid": {
                    "type": "integer"
                },
                "tag": {
                    "description": "标签",
                    "type": "string"
                },
                "uid": {
                    "description": "序号",
                    "type": "integer"
                }
            }
        },
        "handler.loginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.7",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{"https"},
	Title:            "Melting API",
	Description:      "# TODO List\r\n* ~~使用MongoDB存储resource/~~\r\n* 优化图片上传\r\n* 日志记录\r\n* 调整环境变量名称\r\n* websocket\r\n\r\n\r\n\r\n",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
