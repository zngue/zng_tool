package httpx

import (
	"encoding/json"
	stderrors "errors"
	"net/http"

	"github.com/zngue/zng_app/pkg/errors_ez"
)

type MessageCode int

const (
	CodeSuccess    MessageCode = 200
	CodeParamError MessageCode = 10001
	CodeError      MessageCode = 400
)

func (m MessageCode) String() string {
	switch m {
	case CodeSuccess:
		return "success"
	case CodeParamError:
		return "参数错误"
	case CodeError:
		return "响应错误"
	default:
		return "未知错误"
	}
}

type Response struct {
	Code      MessageCode         `json:"statusCode"`
	Message   string              `json:"message"`
	Reason    string              `json:"reason,omitempty"`
	RequestId string              `json:"requestId,omitempty"`
	Data      any                 `json:"data,omitempty"`
	Err       []*errors_ez.EzInfo `json:"err,omitempty"`
}

func DataWithErr(w http.ResponseWriter, r *http.Request, err error, data any) {
	if err != nil {
		Error(w, r, err)
		return
	}
	Success(w, r, data)
}

func Success(w http.ResponseWriter, r *http.Request, data any) {
	write(w, &Response{
		Code:      CodeSuccess,
		Message:   CodeSuccess.String(),
		RequestId: requestID(r),
		Data:      data,
	})
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	writeError(w, r, CodeError, err)
}

func ParamError(w http.ResponseWriter, r *http.Request, err error) {
	writeError(w, r, CodeParamError, err)
}

func writeError(w http.ResponseWriter, r *http.Request, code MessageCode, err error) {
	res := &Response{
		Code:      code,
		RequestId: requestID(r),
	}
	var ezErr *errors_ez.EzError
	if stderrors.As(err, &ezErr) {
		res.Err = ezErr.ReasonMessage()
		res.Message = ezErr.LastCustomReason()
	} else if err != nil {
		res.Reason = err.Error()
	}
	if res.Message == "" {
		res.Message = code.String()
	}
	write(w, res)
}

func write(w http.ResponseWriter, res *Response) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}
