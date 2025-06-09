package serializer

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Response Base serializer
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// TrackedErrorResponse Error response with tracking information
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// Three-digit error codes reuse HTTP original meanings
// Five-digit error codes are application custom errors
// Five-digit error codes starting with 5 are server-side errors, such as database operation failures
// Five-digit error codes starting with 4 are client-side errors, sometimes due to client code errors, sometimes due to user operation errors
const (
	// CodeCheckLogin Not logged in
	CodeCheckLogin = 401
	// CodeNoRightErr Unauthorized access
	CodeNoRightErr = 403
	// CodeDBError Database operation failed
	CodeDBError = 50001
	// CodeEncryptError Encryption failed
	CodeEncryptError = 50002
	// CodeParamErr Various parameter errors
	CodeParamErr = 40001
)

// CheckLogin Check login status
func CheckLogin() Response {
	return Response{
		Code: CodeCheckLogin,
		Msg:  "Not logged in",
	}
}

// Err Common error handling
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	// Hide underlying errors in production environment
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = fmt.Sprintf("%+v", err)
	}
	return res
}

// DBErr Database operation failed
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "Database operation failed"
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr Various parameter errors
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "Parameter error"
	}
	return Err(CodeParamErr, msg, err)
}
