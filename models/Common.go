package models

import "go-web-server/enums"

type JsonResult struct{
	Code enums.JsonResultCode `json:"code"`//把struct编码成json字符串时，Code字段的key是code
	Msg string `json:"msg"`
	Obj interface{} `json:"obj"`
}
