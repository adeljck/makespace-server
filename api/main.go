package api

import (
	"encoding/json"
	"makespace-remaster/serializer"
)


func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
		}
	}

	return serializer.Response{
		Status: 40002,
		Msg:    "参数错误",
	}
}

