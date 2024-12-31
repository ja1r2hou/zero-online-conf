package respx

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type RespOk struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"traceId"`
}

type RespErr struct {
	Status int `json:"status"`
	//Errors errorx.CodeError `json:"errors"`
	//Code int   `json:"code"` //错误码
	Msg     string `json:"msg"` //错误码对应key
	TraceId string `json:"traceId"`
}

// Ok writes HTTP 200 OK into w.
func Ok(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

// OkJson writes v into w with 200 OK.
func OkJson(w http.ResponseWriter, v interface{}, ctx context.Context) {
	r := &RespOk{
		Status:  0,
		Data:    v,
		TraceId: trace.TraceIDFromContext(ctx),
	}
	WriteJson(w, http.StatusOK, r)
}

// ErrJson writes v into w with 500 OK.
func ErrJson(w http.ResponseWriter, err error) {
	r := &RespErr{
		Status: 500,
		Msg:    err.Error(),
	}
	WriteJson(w, http.StatusInternalServerError, r)
}

// Error writes err into w.
func Error(w http.ResponseWriter, err error, ctx context.Context) {
	r := &RespErr{
		Status:  500,
		Msg:     err.Error(),
		TraceId: trace.TraceIDFromContext(ctx),
	}
	WriteJson(w, http.StatusInternalServerError, r)
}

// WriteJson writes v as json string into w with code.
func WriteJson(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set(httpx.ContentType, httpx.JsonContentType)
	w.WriteHeader(code)

	if bs, err := json.Marshal(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			logx.Errorf("write response failed, error: %s", err)
		}
	} else if n < len(bs) {
		logx.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}
}
