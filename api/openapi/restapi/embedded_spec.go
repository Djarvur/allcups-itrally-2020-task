// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "# IT RALLY 2020 HighLoad Cup\n## Usage\n- The goal is to dig up and cash out treasures to increase your balance\n  as much as possible in fixed period of time (10 minutes).\n- [Optional] Use /explore to find out where treasures were buried.\n- Use /license to ensure you have valid license ID before calling /dig.\n  Each /dig call will increment license's DigUsed up to DigAllowed,\n  then this license will become inactive.\n- Use /dig with Depth increasing serially from 1 to 10, and sometimes you'll\n  find a treasure.\n- Use /cash on each treasure found by /dig to increase your balance.\n- [Optional] You can pay for a license if you like to - this will\n  decrease your balance, but paid licenses have more DigAllowed, so\n  you'll have to issue less licenses.\n  Each coin in the wallet returned by /cash is unique, and may be spent to get\n  paid license just once.\n- [Optional] You can get your current balance and coins using /balance, and your\n  current active licenses using /licenses.\n## List of all custom errors\nFirst number is HTTP Status code, second is value of \"code\" field in returned JSON object, text description may or may not match \"message\" field in returned JSON object.\n- 422.1000: wrong coordinates\n- 422.1001: wrong depth\n- 409.1002: no more active licenses allowed\n- 409.1003: treasure is not digged\n",
    "title": "HighLoad Cup 2020",
    "version": "0.3.0"
  },
  "basePath": "/",
  "paths": {
    "/balance": {
      "get": {
        "description": "Returns a current balance.",
        "operationId": "getBalance",
        "responses": {
          "200": {
            "$ref": "#/responses/balance"
          },
          "default": {
            "$ref": "#/responses/error"
          }
        }
      }
    },
    "/cash": {
      "post": {
        "description": "Exchange provided treasure for money.",
        "operationId": "cash",
        "parameters": [
          {
            "description": "Treasure for exchange.",
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/treasure"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/cash"
          },
          "default": {
            "description": "- 409.1003: treasure is not digged\n",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/dig": {
      "post": {
        "description": "Dig at given point and depth, returns found treasures.",
        "operationId": "dig",
        "parameters": [
          {
            "description": "License, place and depth to dig.",
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dig"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/dig"
          },
          "default": {
            "description": "- 422.1000: wrong coordinates\n- 422.1001: wrong depth\n",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/explore": {
      "post": {
        "description": "Returns amount of treasures in the provided area at full depth.",
        "operationId": "exploreArea",
        "parameters": [
          {
            "description": "Area to be explored.",
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/area"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/explore"
          },
          "default": {
            "description": "- 422.1000: wrong coordinates\n",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/health-check": {
      "get": {
        "security": [],
        "description": "Returns 200 if service works okay.",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "Extra details about service status, if any.",
            "schema": {
              "type": "object",
              "additionalProperties": true
            }
          },
          "default": {
            "$ref": "#/responses/error"
          }
        }
      }
    },
    "/licenses": {
      "get": {
        "description": "Returns a list of issued licenses.",
        "operationId": "listLicenses",
        "responses": {
          "200": {
            "$ref": "#/responses/licenseList"
          },
          "default": {
            "$ref": "#/responses/error"
          }
        }
      },
      "post": {
        "description": "Issue a new license.",
        "operationId": "issueLicense",
        "parameters": [
          {
            "description": "Amount of money to spend for a license.",
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/wallet"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/license"
          },
          "default": {
            "description": "- 409.1002: no more active licenses allowed\n",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "amount": {
      "description": "Non-negative amount of treasures/etc.",
      "type": "integer"
    },
    "area": {
      "type": "object",
      "required": [
        "posX",
        "posY"
      ],
      "properties": {
        "posX": {
          "type": "integer",
          "x-order": 0
        },
        "posY": {
          "type": "integer",
          "x-order": 1
        },
        "sizeX": {
          "type": "integer",
          "minimum": 1,
          "x-order": 2
        },
        "sizeY": {
          "type": "integer",
          "minimum": 1,
          "x-order": 3
        }
      }
    },
    "balance": {
      "description": "Current balance and wallet with up to 1000 coins.",
      "type": "object",
      "required": [
        "balance",
        "wallet"
      ],
      "properties": {
        "balance": {
          "type": "integer",
          "format": "uint32",
          "x-order": 0
        },
        "wallet": {
          "x-order": 1,
          "$ref": "#/definitions/wallet"
        }
      }
    },
    "dig": {
      "type": "object",
      "required": [
        "licenseID",
        "posX",
        "posY",
        "depth"
      ],
      "properties": {
        "depth": {
          "type": "integer",
          "maximum": 100,
          "minimum": 1,
          "x-order": 3
        },
        "licenseID": {
          "description": "ID of the license this request is attached to.",
          "type": "integer",
          "x-order": 0
        },
        "posX": {
          "type": "integer",
          "x-order": 1
        },
        "posY": {
          "type": "integer",
          "x-order": 2
        }
      }
    },
    "error": {
      "description": "This model should match output of errors returned by go-swagger\n(like failed validation), to ensure our handlers use same format.\n",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600 with HTTP Status Code 422.",
          "type": "integer",
          "format": "int32",
          "x-order": 0
        },
        "message": {
          "type": "string",
          "x-order": 1
        }
      }
    },
    "license": {
      "description": "License for digging.",
      "type": "object",
      "required": [
        "id",
        "digAllowed",
        "digUsed"
      ],
      "properties": {
        "digAllowed": {
          "x-order": 1,
          "$ref": "#/definitions/amount"
        },
        "digUsed": {
          "x-order": 2,
          "$ref": "#/definitions/amount"
        },
        "id": {
          "type": "integer",
          "x-order": 0
        }
      }
    },
    "licenseList": {
      "description": "List of issued licenses.",
      "type": "array",
      "items": {
        "$ref": "#/definitions/license"
      }
    },
    "report": {
      "type": "object",
      "required": [
        "area",
        "amount"
      ],
      "properties": {
        "amount": {
          "x-order": 1,
          "$ref": "#/definitions/amount"
        },
        "amountPerDepth": {
          "description": "Histogram, key is depth (\"1\", \"2\", …).",
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/amount"
          },
          "x-order": 2
        },
        "area": {
          "x-order": 0,
          "$ref": "#/definitions/area"
        }
      }
    },
    "treasure": {
      "description": "Treasure ID.",
      "type": "string"
    },
    "treasureList": {
      "description": "List of treasures.",
      "type": "array",
      "items": {
        "$ref": "#/definitions/treasure"
      }
    },
    "wallet": {
      "description": "Wallet with some coins.",
      "type": "array",
      "maxItems": 1000,
      "uniqueItems": true,
      "items": {
        "type": "integer",
        "format": "uint32"
      }
    }
  },
  "responses": {
    "balance": {
      "description": "Current balance.",
      "schema": {
        "$ref": "#/definitions/balance"
      }
    },
    "cash": {
      "description": "Payment for treasure.",
      "schema": {
        "$ref": "#/definitions/wallet"
      }
    },
    "dig": {
      "description": "List of treasures found.",
      "schema": {
        "$ref": "#/definitions/treasureList"
      }
    },
    "error": {
      "description": "General errors using same model as used by go-swagger for validation errors.",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "explore": {
      "description": "Report about found treasures.",
      "schema": {
        "$ref": "#/definitions/report"
      }
    },
    "license": {
      "description": "Issued license.",
      "schema": {
        "$ref": "#/definitions/license"
      }
    },
    "licenseList": {
      "description": "List of issued licenses.",
      "schema": {
        "$ref": "#/definitions/licenseList"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "# IT RALLY 2020 HighLoad Cup\n## Usage\n- The goal is to dig up and cash out treasures to increase your balance\n  as much as possible in fixed period of time (10 minutes).\n- [Optional] Use /explore to find out where treasures were buried.\n- Use /license to ensure you have valid license ID before calling /dig.\n  Each /dig call will increment license's DigUsed up to DigAllowed,\n  then this license will become inactive.\n- Use /dig with Depth increasing serially from 1 to 10, and sometimes you'll\n  find a treasure.\n- Use /cash on each treasure found by /dig to increase your balance.\n- [Optional] You can pay for a license if you like to - this will\n  decrease your balance, but paid licenses have more DigAllowed, so\n  you'll have to issue less licenses.\n  Each coin in the wallet returned by /cash is unique, and may be spent to get\n  paid license just once.\n- [Optional] You can get your current balance and coins using /balance, and your\n  current active licenses using /licenses.\n## List of all custom errors\nFirst number is HTTP Status code, second is value of \"code\" field in returned JSON object, text description may or may not match \"message\" field in returned JSON object.\n- 422.1000: wrong coordinates\n- 422.1001: wrong depth\n- 409.1002: no more active licenses allowed\n- 409.1003: treasure is not digged\n",
    "title": "HighLoad Cup 2020",
    "version": "0.3.0"
  },
  "basePath": "/",
  "paths": {
    "/balance": {
      "get": {
        "description": "Returns a current balance.",
        "operationId": "getBalance",
        "responses": {
          "200": {
            "description": "Current balance.",
            "schema": {
              "$ref": "#/definitions/balance"
            }
          },
          "default": {
            "description": "General errors using same model as used by go-swagger for validation errors.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/cash": {
      "post": {
        "description": "Exchange provided treasure for money.",
        "operationId": "cash",
        "parameters": [
          {
            "description": "Treasure for exchange.",
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/treasure"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Payment for treasure.",
            "schema": {
              "$ref": "#/definitions/wallet"
            }
          },
          "default": {
            "description": "- 409.1003: treasure is not digged\n",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/dig": {
      "post": {
        "description": "Dig at given point and depth, returns found treasures.",
        "operationId": "dig",
        "parameters": [
          {
            "description": "License, place and depth to dig.",
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dig"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of treasures found.",
            "schema": {
              "$ref": "#/definitions/treasureList"
            }
          },
          "default": {
            "description": "- 422.1000: wrong coordinates\n- 422.1001: wrong depth\n",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/explore": {
      "post": {
        "description": "Returns amount of treasures in the provided area at full depth.",
        "operationId": "exploreArea",
        "parameters": [
          {
            "description": "Area to be explored.",
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/area"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Report about found treasures.",
            "schema": {
              "$ref": "#/definitions/report"
            }
          },
          "default": {
            "description": "- 422.1000: wrong coordinates\n",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/health-check": {
      "get": {
        "security": [],
        "description": "Returns 200 if service works okay.",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "Extra details about service status, if any.",
            "schema": {
              "type": "object",
              "additionalProperties": true
            }
          },
          "default": {
            "description": "General errors using same model as used by go-swagger for validation errors.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/licenses": {
      "get": {
        "description": "Returns a list of issued licenses.",
        "operationId": "listLicenses",
        "responses": {
          "200": {
            "description": "List of issued licenses.",
            "schema": {
              "$ref": "#/definitions/licenseList"
            }
          },
          "default": {
            "description": "General errors using same model as used by go-swagger for validation errors.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "Issue a new license.",
        "operationId": "issueLicense",
        "parameters": [
          {
            "description": "Amount of money to spend for a license.",
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/wallet"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Issued license.",
            "schema": {
              "$ref": "#/definitions/license"
            }
          },
          "default": {
            "description": "- 409.1002: no more active licenses allowed\n",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "amount": {
      "description": "Non-negative amount of treasures/etc.",
      "type": "integer",
      "minimum": 0
    },
    "area": {
      "type": "object",
      "required": [
        "posX",
        "posY"
      ],
      "properties": {
        "posX": {
          "type": "integer",
          "minimum": 0,
          "x-order": 0
        },
        "posY": {
          "type": "integer",
          "minimum": 0,
          "x-order": 1
        },
        "sizeX": {
          "type": "integer",
          "minimum": 1,
          "x-order": 2
        },
        "sizeY": {
          "type": "integer",
          "minimum": 1,
          "x-order": 3
        }
      }
    },
    "balance": {
      "description": "Current balance and wallet with up to 1000 coins.",
      "type": "object",
      "required": [
        "balance",
        "wallet"
      ],
      "properties": {
        "balance": {
          "type": "integer",
          "format": "uint32",
          "x-order": 0
        },
        "wallet": {
          "x-order": 1,
          "$ref": "#/definitions/wallet"
        }
      }
    },
    "dig": {
      "type": "object",
      "required": [
        "licenseID",
        "posX",
        "posY",
        "depth"
      ],
      "properties": {
        "depth": {
          "type": "integer",
          "maximum": 100,
          "minimum": 1,
          "x-order": 3
        },
        "licenseID": {
          "description": "ID of the license this request is attached to.",
          "type": "integer",
          "x-order": 0
        },
        "posX": {
          "type": "integer",
          "minimum": 0,
          "x-order": 1
        },
        "posY": {
          "type": "integer",
          "minimum": 0,
          "x-order": 2
        }
      }
    },
    "error": {
      "description": "This model should match output of errors returned by go-swagger\n(like failed validation), to ensure our handlers use same format.\n",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600 with HTTP Status Code 422.",
          "type": "integer",
          "format": "int32",
          "x-order": 0
        },
        "message": {
          "type": "string",
          "x-order": 1
        }
      }
    },
    "license": {
      "description": "License for digging.",
      "type": "object",
      "required": [
        "id",
        "digAllowed",
        "digUsed"
      ],
      "properties": {
        "digAllowed": {
          "x-order": 1,
          "$ref": "#/definitions/amount"
        },
        "digUsed": {
          "x-order": 2,
          "$ref": "#/definitions/amount"
        },
        "id": {
          "type": "integer",
          "x-order": 0
        }
      }
    },
    "licenseList": {
      "description": "List of issued licenses.",
      "type": "array",
      "items": {
        "$ref": "#/definitions/license"
      }
    },
    "report": {
      "type": "object",
      "required": [
        "area",
        "amount"
      ],
      "properties": {
        "amount": {
          "x-order": 1,
          "$ref": "#/definitions/amount"
        },
        "amountPerDepth": {
          "description": "Histogram, key is depth (\"1\", \"2\", …).",
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/amount"
          },
          "x-order": 2
        },
        "area": {
          "x-order": 0,
          "$ref": "#/definitions/area"
        }
      }
    },
    "treasure": {
      "description": "Treasure ID.",
      "type": "string"
    },
    "treasureList": {
      "description": "List of treasures.",
      "type": "array",
      "items": {
        "$ref": "#/definitions/treasure"
      }
    },
    "wallet": {
      "description": "Wallet with some coins.",
      "type": "array",
      "maxItems": 1000,
      "uniqueItems": true,
      "items": {
        "type": "integer",
        "format": "uint32"
      }
    }
  },
  "responses": {
    "balance": {
      "description": "Current balance.",
      "schema": {
        "$ref": "#/definitions/balance"
      }
    },
    "cash": {
      "description": "Payment for treasure.",
      "schema": {
        "$ref": "#/definitions/wallet"
      }
    },
    "dig": {
      "description": "List of treasures found.",
      "schema": {
        "$ref": "#/definitions/treasureList"
      }
    },
    "error": {
      "description": "General errors using same model as used by go-swagger for validation errors.",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "explore": {
      "description": "Report about found treasures.",
      "schema": {
        "$ref": "#/definitions/report"
      }
    },
    "license": {
      "description": "Issued license.",
      "schema": {
        "$ref": "#/definitions/license"
      }
    },
    "licenseList": {
      "description": "List of issued licenses.",
      "schema": {
        "$ref": "#/definitions/licenseList"
      }
    }
  }
}`))
}
