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
    "description": "Post messages to Slack",
    "title": "slack",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "social"
      ],
      "container": "gcr.io/direktiv/functions/slack",
      "issues": "https://github.com/direktiv-apps/slack/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function can post messages to Slack channels. It is required to have a Slack account and an app create at the Slack [API portal](https://api.slack.com/apps).  This function can send simple messages as well as blocks. \nAfter creating a Slack app an \"Incoming Webhook\" has to be created which is the ` + "`" + `webhook-url` + "`" + ` for this function. ",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/slack"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "content": {
                  "type": "object",
                  "additionalProperties": true
                },
                "verbose": {
                  "description": "Enables verbose output",
                  "type": "boolean"
                },
                "webhook-url": {
                  "description": "URL for Slack's incoming webhook",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "type": "object",
              "properties": {
                "slack": {
                  "type": "object",
                  "properties": {
                    "result": {
                      "type": "string"
                    },
                    "success": {
                      "type": "boolean"
                    }
                  }
                }
              }
            },
            "examples": {
              "slack": [
                {
                  "result": "ok",
                  "success": true
                }
              ]
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "echo sending message to slack",
              "print": false,
              "silent": false
            },
            {
              "action": "exec",
              "exec": "curl -X POST -H 'Content-type: application/json' --data '{{ .Content | toJson }}' {{ .WebhookURL }}",
              "print": "{{ .Verbose }}",
              "silent": "{{- if .Verbose }}false{{- else}}true{{- end }}"
            }
          ],
          "output": "{\n  \"slack\": {{ index . 1 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: slack\n  type: action\n  action:\n    function: slack\n    secrets: [\"webhook-url\"]\n    input: \n      webhook-url: jq(.secrets.\"webhook-url\")\n    content:\n      text: Hello World",
            "title": "Simple text"
          },
          {
            "content": "- id: slack\n  type: action\n  action:\n    function: slack\n    secrets: [\"webhook-url\"]\n    input: \n      webhook-url: jq(.secrets.\"webhook-url\")\n      content:\n        - type: section\n          text:\n            type: mrkdwn\n            text: Test Me",
            "title": "Post with blocks"
          }
        ],
        "x-direktiv-function": "functions:\n- id: slack\n  image: gcr.io/direktiv/functions/slack:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "Incoming webhook URL from Slack",
            "name": "webhook-url"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
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
    "description": "Post messages to Slack",
    "title": "slack",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "social"
      ],
      "container": "gcr.io/direktiv/functions/slack",
      "issues": "https://github.com/direktiv-apps/slack/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function can post messages to Slack channels. It is required to have a Slack account and an app create at the Slack [API portal](https://api.slack.com/apps).  This function can send simple messages as well as blocks. \nAfter creating a Slack app an \"Incoming Webhook\" has to be created which is the ` + "`" + `webhook-url` + "`" + ` for this function. ",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/slack"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/postParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "$ref": "#/definitions/postOKBody"
            },
            "examples": {
              "slack": [
                {
                  "result": "ok",
                  "success": true
                }
              ]
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "echo sending message to slack",
              "print": false,
              "silent": false
            },
            {
              "action": "exec",
              "exec": "curl -X POST -H 'Content-type: application/json' --data '{{ .Content | toJson }}' {{ .WebhookURL }}",
              "print": "{{ .Verbose }}",
              "silent": "{{- if .Verbose }}false{{- else}}true{{- end }}"
            }
          ],
          "output": "{\n  \"slack\": {{ index . 1 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: slack\n  type: action\n  action:\n    function: slack\n    secrets: [\"webhook-url\"]\n    input: \n      webhook-url: jq(.secrets.\"webhook-url\")\n    content:\n      text: Hello World",
            "title": "Simple text"
          },
          {
            "content": "- id: slack\n  type: action\n  action:\n    function: slack\n    secrets: [\"webhook-url\"]\n    input: \n      webhook-url: jq(.secrets.\"webhook-url\")\n      content:\n        - type: section\n          text:\n            type: mrkdwn\n            text: Test Me",
            "title": "Post with blocks"
          }
        ],
        "x-direktiv-function": "functions:\n- id: slack\n  image: gcr.io/direktiv/functions/slack:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "Incoming webhook URL from Slack",
            "name": "webhook-url"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    },
    "postOKBody": {
      "type": "object",
      "properties": {
        "slack": {
          "$ref": "#/definitions/postOKBodySlack"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postOKBodySlack": {
      "type": "object",
      "properties": {
        "result": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBody": {
      "type": "object",
      "properties": {
        "content": {
          "type": "object",
          "additionalProperties": true
        },
        "verbose": {
          "description": "Enables verbose output",
          "type": "boolean"
        },
        "webhook-url": {
          "description": "URL for Slack's incoming webhook",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    }
  }
}`))
}
