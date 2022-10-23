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
        "/base/advertList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "获取轮播图",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/response.PageResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/system.SysAdvert"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/base/captcha": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "生成验证码",
                "responses": {
                    "200": {
                        "description": "生成验证码,返回包括随机数id,base64,验证码长度",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.SysCaptchaResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/base/deleteAdvert/:id": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "删除轮播图",
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        " msg": {
                                            "type": "string"
                                        },
                                        "data": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/base/insertAdvert": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "添加轮播图",
                "parameters": [
                    {
                        "description": " ",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.InsertAdvert"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/base/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名, 密码, 验证码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回包括用户信息,token,过期时间",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.LoginResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/base/updateAdvert": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "更新轮播图",
                "parameters": [
                    {
                        "description": " ",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateAdvert"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        " msg": {
                                            "type": "string"
                                        },
                                        "data": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/base/wxLogin": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "授权登录",
                "parameters": [
                    {
                        "description": "用户名, 密码, 验证码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.WxLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回包括用户信息,token,过期时间",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.WxLoginResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jwt/jsonInBlacklist": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jwt"
                ],
                "summary": "jwt加入黑名单",
                "responses": {
                    "200": {
                        "description": "jwt加入黑名单",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.InsertAdvert": {
            "type": "object",
            "properties": {
                "action": {
                    "description": "链接值",
                    "type": "string"
                },
                "name": {
                    "description": "广告名称",
                    "type": "string"
                },
                "photo": {
                    "description": "图片地址",
                    "type": "string"
                },
                "position": {
                    "description": "广告位置 1首页轮播;",
                    "type": "integer"
                },
                "sort": {
                    "description": "排序",
                    "type": "integer"
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                },
                "type": {
                    "description": "广告类型 product 产品 news 资讯 index 首页",
                    "type": "string"
                }
            }
        },
        "request.Login": {
            "type": "object",
            "properties": {
                "captcha": {
                    "description": "验证码",
                    "type": "string"
                },
                "captchaId": {
                    "description": "验证码ID",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "userName": {
                    "description": "用户名-手机号码",
                    "type": "string"
                }
            }
        },
        "request.UpdateAdvert": {
            "type": "object",
            "properties": {
                "action": {
                    "description": "链接值",
                    "type": "string"
                },
                "id": {
                    "description": "id",
                    "type": "string"
                },
                "name": {
                    "description": "广告名称",
                    "type": "string"
                },
                "photo": {
                    "description": "图片地址",
                    "type": "string"
                },
                "position": {
                    "description": "广告位置 1首页轮播;",
                    "type": "integer"
                },
                "sort": {
                    "description": "排序",
                    "type": "integer"
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                },
                "type": {
                    "description": "广告类型 product 产品 news 资讯 index 首页",
                    "type": "string"
                }
            }
        },
        "request.WxLogin": {
            "type": "object",
            "properties": {
                "appid": {
                    "description": "微信应用id",
                    "type": "string"
                },
                "code": {
                    "description": "微信用户登录凭证（有效期五分钟）",
                    "type": "string"
                },
                "encryptedData": {
                    "description": "微信用户信息的加密数据",
                    "type": "string"
                },
                "iv": {
                    "description": "加密算法的初始向量",
                    "type": "string"
                }
            }
        },
        "response.LoginResponse": {
            "type": "object",
            "properties": {
                "expiresAt": {
                    "description": "凭证过期时间",
                    "type": "integer"
                },
                "openId": {
                    "description": "openid",
                    "type": "string"
                },
                "token": {
                    "description": "登录凭证",
                    "type": "string"
                },
                "user": {
                    "description": "用户信息",
                    "$ref": "#/definitions/system.SysUser"
                }
            }
        },
        "response.PageResult": {
            "type": "object",
            "properties": {
                "list": {},
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "response.SysCaptchaResponse": {
            "type": "object",
            "properties": {
                "captchaId": {
                    "type": "string"
                },
                "captchaLength": {
                    "type": "integer"
                },
                "picPath": {
                    "type": "string"
                }
            }
        },
        "response.WxLoginResponse": {
            "type": "object",
            "properties": {
                "expiresAt": {
                    "description": "凭证过期时间",
                    "type": "integer"
                },
                "openId": {
                    "description": "openid",
                    "type": "string"
                },
                "token": {
                    "description": "登录凭证",
                    "type": "string"
                },
                "user": {
                    "description": "用户信息",
                    "$ref": "#/definitions/system.SysUser"
                }
            }
        },
        "system.SysAdvert": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "createdTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "deletedTime": {
                    "description": "删除时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "position": {
                    "type": "integer"
                },
                "sort": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedTime": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "system.SysUser": {
            "type": "object",
            "properties": {
                "Mobile": {
                    "description": "用户手机号",
                    "type": "string"
                },
                "avatar": {
                    "description": "用户头像",
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "createdTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "deletedTime": {
                    "description": "删除时间",
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "enable": {
                    "description": "用户是否被冻结 1正常 2冻结",
                    "type": "integer"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "openid": {
                    "type": "string"
                },
                "points": {
                    "type": "number"
                },
                "province": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "updatedTime": {
                    "description": "更新时间",
                    "type": "string"
                },
                "userName": {
                    "description": "用户账号-用户登录名",
                    "type": "string"
                },
                "uuid": {
                    "description": "用户UUID",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample Server pets",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
