package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = iota + 1000
	ERROR
	ParamError
	AuthError
	UserRegisterFail
	UserNameEmpty
	UserPassEmpty

	InternalServerError = http.StatusInternalServerError
)

const (
	OkMsg                  = "operation success"
	NoOkMsg                = "operation failed"
	ParamErrorMsg          = "parameters format error"
	LoginCheckErrorMsg     = "user name or password error"
	UserRegisterErrorMsg   = "user register failed"
	UserNameIsEmptyMsg     = "user name is empty"
	UserPasswordIsEmptyMsg = "user password is empty"
	InternalServerErrorMsg = "server internal error"
)

type response struct {
	ErrCode int         `json:"errCode"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	ErrMsg  string      `json:"errMsg"`
}

var customError = map[int]string{
	SUCCESS:             OkMsg,
	ERROR:               NoOkMsg,
	ParamError:          ParamErrorMsg,
	AuthError:           LoginCheckErrorMsg,
	UserRegisterFail:    UserRegisterErrorMsg,
	UserNameEmpty:       UserNameIsEmptyMsg,
	UserPassEmpty:       UserPasswordIsEmptyMsg,
	InternalServerError: InternalServerErrorMsg,
}

func ResultOk(code int, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response{
		ErrCode: code,
		Data:    data,
		Msg:     msg,
	})
}

func ResultFail(code int, data interface{}, msg string, ctx *gin.Context) {
	if msg == "" {
		ctx.JSON(http.StatusOK, response{
			ErrCode: code,
			Data:    data,
			ErrMsg:  customError[code],
		})
	} else {
		ctx.JSON(http.StatusOK, response{
			ErrCode: code,
			Data:    data,
			ErrMsg:  msg,
		})
	}
}

func Ok(c *gin.Context) {
	ResultOk(SUCCESS, map[string]interface{}{}, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	ResultOk(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	ResultOk(SUCCESS, data, "success", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	ResultOk(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	ResultFail(ERROR, map[string]interface{}{}, "failed", c)
}

func FailWithMessage(code int, message string, c *gin.Context) {
	ResultFail(code, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, code int, message string, c *gin.Context) {
	ResultFail(code, data, message, c)
}
