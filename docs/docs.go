// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "AlvISs_Reimu",
            "url": "https://github.com/AlvISsReimu",
            "email": "alvissreimu@gmail.com"
        },
        "license": {
            "name": "MIT License",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v2/items": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Get all Items",
                "deprecated": true,
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/shims.Item"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "existence": {
                                                "$ref": "#/definitions/models.Existence"
                                            },
                                            "name_i18n": {
                                                "$ref": "#/definitions/models.I18nString"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        },
        "/v2/items/{itemId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Get an Item with ID",
                "deprecated": true,
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "itemId",
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
                                    "$ref": "#/definitions/shims.Item"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "existence": {
                                            "$ref": "#/definitions/models.Existence"
                                        },
                                        "name_i18n": {
                                            "$ref": "#/definitions/models.I18nString"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid or missing itemId. Notice that this shall be the **string ID** of the item, instead of the internally used numerical ID of the item.",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        },
        "/v2/stages": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stage"
                ],
                "summary": "Get all Stages",
                "deprecated": true,
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/shims.Stage"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "code_i18n": {
                                                "$ref": "#/definitions/models.I18nString"
                                            },
                                            "existence": {
                                                "$ref": "#/definitions/models.Existence"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        },
        "/v2/stages/{stageId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stage"
                ],
                "summary": "Get an Stage with ID",
                "deprecated": true,
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Stage ID",
                        "name": "stageId",
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
                                    "$ref": "#/definitions/shims.Stage"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code_i18n": {
                                            "$ref": "#/definitions/models.I18nString"
                                        },
                                        "existence": {
                                            "$ref": "#/definitions/models.Existence"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid or missing stageId. Notice that this shall be the **string ID** of the stage, instead of the internally used numerical ID of the stage.",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        },
        "/v2/zones": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Zone"
                ],
                "summary": "Get all Zones",
                "deprecated": true,
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/shims.Zone"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "existence": {
                                                "$ref": "#/definitions/models.Existence"
                                            },
                                            "zoneName_i18n": {
                                                "$ref": "#/definitions/models.I18nString"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        },
        "/v2/zones/{zoneId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Zone"
                ],
                "summary": "Get a Zone with ID",
                "deprecated": true,
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Zone ID",
                        "name": "zoneId",
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
                                    "$ref": "#/definitions/shims.Zone"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "existence": {
                                            "$ref": "#/definitions/models.Existence"
                                        },
                                        "zoneName_i18n": {
                                            "$ref": "#/definitions/models.I18nString"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid or missing zoneId. Notice that this shall be the **string ID** of the zone, instead of the v3 API internally used numerical ID of the zone.",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        },
        "/v3/items": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Get all Items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/models.Item"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "existence": {
                                                "$ref": "#/definitions/models.Existence"
                                            },
                                            "keywords": {
                                                "$ref": "#/definitions/models.Keywords"
                                            },
                                            "name": {
                                                "$ref": "#/definitions/models.I18nString"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        },
        "/v3/items/{itemId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Get an Item with ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "itemId",
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
                                    "$ref": "#/definitions/models.Item"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "existence": {
                                            "$ref": "#/definitions/models.Existence"
                                        },
                                        "keywords": {
                                            "$ref": "#/definitions/models.Keywords"
                                        },
                                        "name": {
                                            "$ref": "#/definitions/models.I18nString"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid or missing itemId. Notice that this shall be the **string ID** of the item, instead of the internally used numerical ID of the item.",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        },
        "/v3/stages": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stage"
                ],
                "summary": "Get all Stages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/models.Stage"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "code": {
                                                "$ref": "#/definitions/models.I18nString"
                                            },
                                            "existence": {
                                                "$ref": "#/definitions/models.Existence"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        },
        "/v3/stages/{stageId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stage"
                ],
                "summary": "Get an Stage with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Stage ID",
                        "name": "stageId",
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
                                    "$ref": "#/definitions/models.Stage"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "$ref": "#/definitions/models.I18nString"
                                        },
                                        "existence": {
                                            "$ref": "#/definitions/models.Existence"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid or missing stageId. Notice that this shall be the **string ID** of the stage, instead of the internally used numerical ID of the stage.",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        },
        "/v3/zones": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Zone"
                ],
                "summary": "Get all Zones",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/models.Zone"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "existence": {
                                                "$ref": "#/definitions/models.Existence"
                                            },
                                            "name": {
                                                "$ref": "#/definitions/models.I18nString"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        },
        "/v3/zones/{zoneId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Zone"
                ],
                "summary": "Get a Zone with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Zone ID",
                        "name": "zoneId",
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
                                    "$ref": "#/definitions/models.Zone"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "existence": {
                                            "$ref": "#/definitions/models.Existence"
                                        },
                                        "name": {
                                            "$ref": "#/definitions/models.I18nString"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid or missing zoneId. Notice that this shall be the **string ID** of the zone, instead of the internally used numerical ID of the zone.",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    },
                    "500": {
                        "description": "An unexpected error occurred",
                        "schema": {
                            "$ref": "#/definitions/errors.PenguinError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.PenguinError": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "string",
                    "example": "INVALID_REQUEST"
                },
                "message": {
                    "type": "string",
                    "example": "invalid request: request parameters are invalid"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "models.Existence": {
            "type": "object",
            "required": [
                "CN",
                "JP",
                "KR",
                "US"
            ],
            "properties": {
                "CN": {
                    "description": "CN: 国服 Mainland China Server (maintained by Hypergryph Network Technology Co., Ltd.)",
                    "$ref": "#/definitions/models.ServerExistence"
                },
                "JP": {
                    "description": "JP: 日服 Japan Server (maintained by Yostar Inc,.)",
                    "$ref": "#/definitions/models.ServerExistence"
                },
                "KR": {
                    "description": "KR: 韩服 Korea Server (maintained by Yostar Limited)",
                    "$ref": "#/definitions/models.ServerExistence"
                },
                "US": {
                    "description": "US: 美服/国际服 Global Server (maintained by Yostar Limited)",
                    "$ref": "#/definitions/models.ServerExistence"
                }
            }
        },
        "models.I18nOptionalString": {
            "type": "object",
            "properties": {
                "en": {
                    "description": "EN: English (en)",
                    "type": "string"
                },
                "ja": {
                    "description": "JP: 日本語 (ja)",
                    "type": "string"
                },
                "ko": {
                    "description": "KR: 한국어 (ko)",
                    "type": "string"
                },
                "zh": {
                    "description": "ZH: 中文 (zh-CN)",
                    "type": "string"
                }
            }
        },
        "models.I18nString": {
            "type": "object",
            "required": [
                "en",
                "ja",
                "ko",
                "zh"
            ],
            "properties": {
                "en": {
                    "description": "EN: English (en)",
                    "type": "string"
                },
                "ja": {
                    "description": "JP: 日本語 (ja)",
                    "type": "string"
                },
                "ko": {
                    "description": "KR: 한국어 (ko)",
                    "type": "string"
                },
                "zh": {
                    "description": "ZH: 中文 (zh-CN)",
                    "type": "string"
                }
            }
        },
        "models.Item": {
            "type": "object",
            "properties": {
                "existence": {
                    "description": "Existence is a map with server code as key and the existence of the item in that server as value.",
                    "type": "object"
                },
                "group": {
                    "description": "Group is an identifier of what the item actually is. For example, both orirock and orirock cube would have the same group, ` + "`" + `orirock` + "`" + `.",
                    "type": "string"
                },
                "itemId": {
                    "description": "ArkItemID (itemId) is the previously used, string form ID of the item; in JSON-representation ` + "`" + `itemId` + "`" + ` is used as key.",
                    "type": "string"
                },
                "keywords": {
                    "description": "Keywords is an arbitrary JSON object containing the keywords of the item, for optimizing the results of the frontend built-in search engine.",
                    "type": "object"
                },
                "name": {
                    "description": "Name is a map with language code as key and the name of the item in that language as value.",
                    "type": "object"
                },
                "penguinItemId": {
                    "description": "ItemID (penguinItemId) is the numerical ID of the item.",
                    "type": "integer"
                },
                "rarity": {
                    "type": "integer"
                },
                "sortId": {
                    "description": "SortID is the sort position of the item.",
                    "type": "integer"
                },
                "sprite": {
                    "description": "Sprite describes the location of the item's sprite on the sprite image, in a form of Y:X.",
                    "type": "string"
                }
            }
        },
        "models.Keywords": {
            "type": "object",
            "properties": {
                "alias": {
                    "description": "Alias of the item,",
                    "$ref": "#/definitions/models.I18nOptionalString"
                },
                "pron": {
                    "description": "Pronounciation hints of the item",
                    "$ref": "#/definitions/models.I18nOptionalString"
                }
            }
        },
        "models.ServerExistence": {
            "type": "object",
            "required": [
                "exist"
            ],
            "properties": {
                "endTime": {
                    "type": "integer"
                },
                "exist": {
                    "type": "boolean"
                },
                "startTime": {
                    "type": "integer"
                }
            }
        },
        "models.Stage": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Code is a map with language code as key and the code of the stage in that language as value.",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "existence": {
                    "description": "Existence is a map with server code as key and the existence of the item in that server as value.",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "minClearTime": {
                    "description": "MinClearTime is the minimum time (in milliseconds as a duration) it takes to clear the stage, referencing from prts.wiki",
                    "type": "number"
                },
                "penguinStageId": {
                    "description": "StageID (penguinStageId) is the numerical ID of the stage.",
                    "type": "integer"
                },
                "sanity": {
                    "description": "Sanity is the sanity requirement for a full clear of the stage.",
                    "type": "number"
                },
                "stageId": {
                    "description": "ArkStageID (stageId) is the previously used, string form ID of the stage; in JSON-representation ` + "`" + `stageId` + "`" + ` is used as key.",
                    "type": "string"
                },
                "zoneId": {
                    "description": "ZoneID is the numerical ID of the zone the stage is in.",
                    "type": "integer"
                }
            }
        },
        "models.Zone": {
            "type": "object",
            "properties": {
                "background": {
                    "description": "Background is the path of the background image of the zone, relative to the CDN endpoint.",
                    "type": "string"
                },
                "category": {
                    "description": "Category of the zone.",
                    "type": "string",
                    "example": "MAINLINE"
                },
                "existence": {
                    "description": "Existence is a map with server code as key and the existence of the item in that server as value.",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "index": {
                    "type": "integer"
                },
                "name": {
                    "description": "Name is a map with language code as key and the name of the item in that language as value.",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "penguinZoneId": {
                    "description": "ZoneID is the numerical ID of the zone.",
                    "type": "integer"
                },
                "type": {
                    "description": "Type of the zone, e.g. \"AWAKENING_HOUR\" or \"VISION_SHATTER\". Optional and only occurres when ` + "`" + `category` + "`" + ` is \"MAINLINE\".",
                    "type": "string",
                    "example": "AWAKENING_HOUR"
                },
                "zoneId": {
                    "type": "string"
                }
            }
        },
        "shims.Item": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "existence": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "groupID": {
                    "type": "string"
                },
                "itemId": {
                    "type": "string"
                },
                "itemType": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "name_i18n": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "pron": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "rarity": {
                    "type": "integer"
                },
                "sortId": {
                    "type": "integer"
                },
                "spriteCoord": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "shims.Stage": {
            "type": "object",
            "properties": {
                "apCost": {
                    "type": "number"
                },
                "code": {
                    "type": "string"
                },
                "code_i18n": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "existence": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "minClearTime": {
                    "type": "number"
                },
                "stageId": {
                    "type": "string"
                },
                "zoneId": {
                    "type": "string"
                }
            }
        },
        "shims.Zone": {
            "type": "object",
            "properties": {
                "background": {
                    "type": "string"
                },
                "existence": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "stages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "subType": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "zoneId": {
                    "type": "string"
                },
                "zoneIndex": {
                    "type": "integer"
                },
                "zoneName": {
                    "type": "string"
                },
                "zoneName_i18n": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
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
	Version:     "3.0.0-alpha.1",
	Host:        "localhost:9010",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Penguin Statistics API",
	Description: "This is the Penguin Statistics v3 API, re-designed to aim for lightweight on wire.",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
