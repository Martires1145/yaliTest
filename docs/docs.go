// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/user/check": {
            "get": {
                "tags": [
                    "User"
                ],
                "summary": "检验用户名是否重复",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "userName",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/user/login": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "userName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "passWord",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/dchan": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "数据传输模块"
                ],
                "summary": "指定形式传输数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "会话id",
                        "name": "client_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "是否为流式传输",
                        "name": "is_stream",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/en/all": {
            "get": {
                "tags": [
                    "Engineering"
                ],
                "summary": "查看全部施工的简述信息",
                "responses": {}
            }
        },
        "/api/v1/en/delete": {
            "post": {
                "tags": [
                    "Engineering"
                ],
                "summary": "删除施工",
                "parameters": [
                    {
                        "type": "string",
                        "description": "要删除的施工ID",
                        "name": "engineeringID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/en/device/add": {
            "post": {
                "tags": [
                    "Engineering"
                ],
                "summary": "增加施工设备",
                "parameters": [
                    {
                        "description": "施工设备",
                        "name": "dd",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Device"
                            }
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/en/device/delete": {
            "post": {
                "tags": [
                    "Engineering"
                ],
                "summary": "删除施工设备",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "要删除的设备ID",
                        "name": "deviceIDs",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/en/new": {
            "post": {
                "tags": [
                    "Engineering"
                ],
                "summary": "新增施工信息",
                "parameters": [
                    {
                        "description": "施工信息",
                        "name": "ed",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EngineeringJson"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/en/re": {
            "post": {
                "tags": [
                    "Engineering"
                ],
                "summary": "修改施工信息",
                "parameters": [
                    {
                        "description": "施工信息",
                        "name": "ed",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EngineeringJson"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/en/{id}": {
            "get": {
                "tags": [
                    "Engineering"
                ],
                "summary": "查看施工详细信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "施工ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/md/new": {
            "post": {
                "tags": [
                    "Model"
                ],
                "summary": "新增模型参数",
                "parameters": [
                    {
                        "description": "模型信息",
                        "name": "md",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.JsonModel"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/path": {
            "get": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件选项模块"
                ],
                "summary": "查看对应类别的路径下有哪些文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文件类型",
                        "name": "file_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/rs": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "脚本模块"
                ],
                "summary": "运行脚本",
                "parameters": [
                    {
                        "description": "脚本运行的参数",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Params"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/upload/csv": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "上传模块"
                ],
                "summary": "上传对应类型的csv文件到对应路径",
                "parameters": [
                    {
                        "type": "file",
                        "description": "csv文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文件类型",
                        "name": "file_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/all": {
            "get": {
                "tags": [
                    "User"
                ],
                "summary": "获取全部用户信息",
                "responses": {}
            }
        },
        "/api/v1/user/d": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "移除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uid",
                        "name": "uid",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/user/info": {
            "get": {
                "tags": [
                    "User"
                ],
                "summary": "获取当前登录用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/user/new": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "ud",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserJson"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/user/rp": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "重置用户密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uid",
                        "name": "uid",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/user/ru": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "用户修改信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名称",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "旧密码",
                        "name": "oldPassWord",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "newPassWord",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/user/v": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "发送验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "userName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/well/all": {
            "get": {
                "tags": [
                    "Well"
                ],
                "summary": "获取所有井信息",
                "responses": {}
            }
        },
        "/api/v1/well/d": {
            "post": {
                "tags": [
                    "Well"
                ],
                "summary": "删除井信息",
                "parameters": [
                    {
                        "description": "井ID",
                        "name": "wellID",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/well/new": {
            "post": {
                "tags": [
                    "Well"
                ],
                "summary": "新增井信息",
                "parameters": [
                    {
                        "description": "井信息",
                        "name": "wd",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.WellJson"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/well/rw": {
            "post": {
                "tags": [
                    "Well"
                ],
                "summary": "编辑井信息",
                "parameters": [
                    {
                        "description": "井信息",
                        "name": "wd",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.WellJson"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/well/{id}": {
            "get": {
                "tags": [
                    "Well"
                ],
                "summary": "获取井的详细信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "井ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.Device": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "engineeringID": {
                    "type": "integer"
                },
                "nameOfDevice": {
                    "type": "string"
                },
                "numberOfDevice": {
                    "type": "integer"
                },
                "typeOfDevice": {
                    "type": "string"
                }
            }
        },
        "model.EngineeringJson": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "address": {
                    "type": "string"
                },
                "beginTime": {
                    "type": "integer"
                },
                "constructionUnit": {
                    "type": "string"
                },
                "devices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Device"
                    }
                },
                "estimatedCompletionTime": {
                    "type": "integer"
                },
                "head": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "numberOfConstructors": {
                    "type": "integer"
                },
                "progress": {
                    "type": "number"
                },
                "state": {
                    "type": "integer"
                },
                "wellName": {
                    "type": "string"
                }
            }
        },
        "model.JsonModel": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "createTime": {
                    "type": "integer"
                },
                "fileCnt": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "params": {
                    "$ref": "#/definitions/model.ParamsJson"
                },
                "useCnt": {
                    "type": "integer"
                }
            }
        },
        "model.Params": {
            "type": "object",
            "properties": {
                "activation": {
                    "type": "string"
                },
                "anomaly_ratio": {
                    "type": "string"
                },
                "batch_size": {
                    "type": "string"
                },
                "c_out": {
                    "type": "string"
                },
                "checkpoints": {
                    "type": "string"
                },
                "d_ff": {
                    "type": "string"
                },
                "d_layers": {
                    "type": "string"
                },
                "d_model": {
                    "type": "string"
                },
                "data": {
                    "type": "string"
                },
                "data_path": {
                    "type": "string"
                },
                "data_test_path": {
                    "type": "string"
                },
                "data_train_path": {
                    "type": "string"
                },
                "data_vali_path": {
                    "type": "string"
                },
                "dec_in": {
                    "type": "string"
                },
                "des": {
                    "type": "string"
                },
                "devices": {
                    "type": "string"
                },
                "distil": {
                    "type": "string"
                },
                "dropout": {
                    "type": "string"
                },
                "e_layers": {
                    "type": "string"
                },
                "embed": {
                    "type": "string"
                },
                "enc_in": {
                    "type": "string"
                },
                "factor": {
                    "type": "string"
                },
                "features": {
                    "type": "string"
                },
                "freq": {
                    "type": "string"
                },
                "gpu": {
                    "type": "string"
                },
                "is_training": {
                    "type": "string"
                },
                "itr": {
                    "type": "string"
                },
                "label_len": {
                    "type": "string"
                },
                "learning_rate": {
                    "type": "string"
                },
                "loss": {
                    "type": "string"
                },
                "lradj": {
                    "type": "string"
                },
                "mask_rate": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "model_id": {
                    "type": "string"
                },
                "moving_avg": {
                    "type": "string"
                },
                "n_heads": {
                    "type": "string"
                },
                "num_kernels": {
                    "type": "string"
                },
                "num_workers": {
                    "type": "string"
                },
                "optim": {
                    "type": "string"
                },
                "output_attention": {
                    "type": "string"
                },
                "p_hidden_dims": {
                    "type": "string"
                },
                "p_hidden_layers": {
                    "type": "string"
                },
                "patience": {
                    "type": "string"
                },
                "pred_len": {
                    "type": "string"
                },
                "root_path": {
                    "type": "string"
                },
                "scale": {
                    "type": "string"
                },
                "seasonal_patterns": {
                    "type": "string"
                },
                "seq_len": {
                    "type": "string"
                },
                "target": {
                    "type": "string"
                },
                "task_name": {
                    "type": "string"
                },
                "top_k": {
                    "type": "string"
                },
                "train_epochs": {
                    "type": "string"
                },
                "use_amp": {
                    "type": "string"
                },
                "use_gpu": {
                    "type": "string"
                },
                "use_kafka": {
                    "type": "string"
                },
                "use_multi_gpu": {
                    "type": "string"
                },
                "w_lin": {
                    "type": "string"
                }
            }
        },
        "model.ParamsExtra": {
            "type": "object",
            "properties": {
                "activation": {
                    "type": "string"
                },
                "anomaly_ratio": {
                    "type": "string"
                },
                "batch_size": {
                    "type": "string"
                },
                "checkpoints": {
                    "type": "string"
                },
                "d_ff": {
                    "type": "string"
                },
                "d_model": {
                    "type": "string"
                },
                "devices": {
                    "type": "string"
                },
                "distil": {
                    "type": "string"
                },
                "dropout": {
                    "type": "string"
                },
                "embed": {
                    "type": "string"
                },
                "freq": {
                    "type": "string"
                },
                "gpu": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "learning_rate": {
                    "type": "string"
                },
                "loss": {
                    "type": "string"
                },
                "lradj": {
                    "type": "string"
                },
                "mask_rate": {
                    "type": "string"
                },
                "moving_avg": {
                    "type": "string"
                },
                "n_heads": {
                    "type": "string"
                },
                "num_kernels": {
                    "type": "string"
                },
                "num_workers": {
                    "type": "string"
                },
                "output_attention": {
                    "type": "string"
                },
                "p_hidden_dims": {
                    "type": "string"
                },
                "p_hidden_layers": {
                    "type": "string"
                },
                "patience": {
                    "type": "string"
                },
                "seasonal_patterns": {
                    "type": "string"
                },
                "top_k": {
                    "type": "string"
                },
                "train_epochs": {
                    "type": "string"
                },
                "use_amp": {
                    "type": "string"
                },
                "use_gpu": {
                    "type": "string"
                },
                "use_multi_gpu": {
                    "type": "string"
                },
                "w_lin": {
                    "type": "string"
                }
            }
        },
        "model.ParamsJson": {
            "type": "object",
            "properties": {
                "pe": {
                    "$ref": "#/definitions/model.ParamsExtra"
                },
                "pu": {
                    "$ref": "#/definitions/model.ParamsUsual"
                },
                "useExtra": {
                    "type": "boolean"
                }
            }
        },
        "model.ParamsUsual": {
            "type": "object",
            "properties": {
                "c_out": {
                    "type": "string"
                },
                "d_layers": {
                    "type": "string"
                },
                "data": {
                    "type": "string"
                },
                "data_path": {
                    "type": "string"
                },
                "data_test_path": {
                    "type": "string"
                },
                "data_train_path": {
                    "type": "string"
                },
                "data_vali_path": {
                    "type": "string"
                },
                "dec_in": {
                    "type": "string"
                },
                "des": {
                    "type": "string"
                },
                "e_layers": {
                    "type": "string"
                },
                "enc_in": {
                    "type": "string"
                },
                "factor": {
                    "type": "string"
                },
                "features": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_training": {
                    "type": "string"
                },
                "itr": {
                    "type": "string"
                },
                "label_len": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "model_id": {
                    "type": "string"
                },
                "optim": {
                    "type": "string"
                },
                "pred_len": {
                    "type": "string"
                },
                "root_path": {
                    "type": "string"
                },
                "scale": {
                    "type": "string"
                },
                "seq_len": {
                    "type": "string"
                },
                "target": {
                    "type": "string"
                },
                "task_name": {
                    "type": "string"
                },
                "use_kafka": {
                    "type": "string"
                }
            }
        },
        "model.UserJson": {
            "type": "object",
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "passWord": {
                    "type": "string"
                },
                "role": {
                    "type": "integer"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "model.WellJson": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "address": {
                    "type": "string"
                },
                "affiliation": {
                    "type": "string"
                },
                "averageDailyProduction": {
                    "type": "integer"
                },
                "boreholeSize": {
                    "type": "integer"
                },
                "construction": {
                    "type": "integer"
                },
                "depth": {
                    "type": "integer"
                },
                "finishTime": {
                    "type": "integer"
                },
                "life": {
                    "type": "integer"
                },
                "note": {
                    "type": "string"
                },
                "wellName": {
                    "type": "string"
                },
                "wellType": {
                    "type": "string"
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
                "message": {
                    "type": "string"
                }
            }
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
