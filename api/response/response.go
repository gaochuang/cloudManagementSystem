package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PageResult represents paginated response data
type PageResult struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"pageSize"`
}

// Error codes
const (
	SUCCESS = iota + 1000
	ERROR
	ParamError
	AuthError
	UserRegisterFail
	UserNameEmpty
	UserPassEmpty
	UserDisable
	InternalServerError   = http.StatusInternalServerError
	CreateK8SClusterError = iota + 2000
)

// Error messages
const (
	OkMsg                    = "operation success"
	NoOkMsg                  = "operation failed"
	ParamErrorMsg            = "parameters format error"
	LoginCheckErrorMsg       = "user name or password error"
	UserRegisterErrorMsg     = "user register failed"
	UserNameIsEmptyMsg       = "user name is empty"
	UserPasswordIsEmptyMsg   = "user password is empty"
	UserDisableMsg           = "user has been disabled"
	InternalServerErrorMsg   = "server internal error"
	CreateK8SClusterErrorMsg = "create kubernetes cluster failed"
)

// Response represents the structure of an API response
type Response struct {
	ErrCode int         `json:"errCode"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg,omitempty"`
	ErrMsg  string      `json:"errMsg,omitempty"`
}

// customError maps error codes to error messages
var customError = map[int]string{
	SUCCESS:               OkMsg,
	ERROR:                 NoOkMsg,
	ParamError:            ParamErrorMsg,
	AuthError:             LoginCheckErrorMsg,
	UserRegisterFail:      UserRegisterErrorMsg,
	UserNameEmpty:         UserNameIsEmptyMsg,
	UserPassEmpty:         UserPasswordIsEmptyMsg,
	InternalServerError:   InternalServerErrorMsg,
	CreateK8SClusterError: CreateK8SClusterErrorMsg,
}

// Result sends a JSON response with the given status, data, and message
func Result(code int, data interface{}, msg string, ctx *gin.Context) {
	if msg == "" {
		msg = customError[code]
	}
	ctx.JSON(http.StatusOK, Response{
		ErrCode: code,
		Data:    data,
		ErrMsg:  msg,
	})
}

// ResultOk sends a successful JSON response
func ResultOk(code int, data interface{}, msg string, ctx *gin.Context) {
	Result(code, data, msg, ctx)
}

// ResultFail sends a failed JSON response
func ResultFail(code int, data interface{}, msg string, ctx *gin.Context) {
	Result(code, data, msg, ctx)
}

// Convenience functions for sending standard responses
func Ok(c *gin.Context) {
	ResultOk(SUCCESS, map[string]interface{}{}, OkMsg, c)
}

func OkWithMessage(message string, c *gin.Context) {
	ResultOk(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	ResultOk(SUCCESS, data, OkMsg, c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	ResultOk(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	ResultFail(ERROR, map[string]interface{}{}, NoOkMsg, c)
}

func FailWithMessage(code int, message string, c *gin.Context) {
	ResultFail(code, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, code int, message string, c *gin.Context) {
	ResultFail(code, data, message, c)
}
