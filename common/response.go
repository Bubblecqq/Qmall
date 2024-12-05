package common

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SUCCESS    = 200
	TOKEN_FAIL = 403
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	TraceId string      `json:"trace_id"`
	Data    interface{} `json:"data,omitempty"`
	Total   interface{} `json:"total"`
	Rows    interface{} `json:"rows"`
}

func JsonBaseResponseCtx(ctx context.Context, w http.ResponseWriter, data any, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := &BaseResponse{
		Data:    data,
		Message: message,
		Code:    code,
	}
	ret, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = w.Write(ret)
}

/*
*
200 OKLoginSuccessVO
201 Created
401 Unauthorized
403 Forbidden
404 Not Found
*
*/

func RespOk(ctx context.Context, w http.ResponseWriter, data interface{}, message string, code int) {
	JsonBaseResponseCtx(ctx, w, data, message, code)
}
