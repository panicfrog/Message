package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseStruct struct {
	Sc ApiStatus     `json:"sc"`
	Message string   `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

// 返回http结果 v需要能进行json解析
func sendResponse(c *gin.Context, bussinessStateCode ApiStatus, message string, v interface{} ) {
	res := ResponseStruct{bussinessStateCode, message, v}
	response, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	var m map[string]interface{}

	err = json.Unmarshal(response, &m)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, m)
}

func sendSuccess(c *gin.Context, message string, v interface{}) {
	sendResponse(c, ApiStatusSuccess, message, v)
}

func sendFail(c *gin.Context, message string) {
	sendResponse(c, ApiStatusFailed, message, nil)
}

func sendParamsError(c *gin.Context, message string) {
	sendResponse(c, ApiStatusParamsError, message, nil)
}

func sendServerInternelError(c *gin.Context, message string) {
	sendResponse(c, ApiStatusInternelError, message, nil)
}

func sendAuthError(c *gin.Context, message string) {
	if message == "" {
		message = "鉴权失败"
	}
	sendResponse(c, ApiStatusUnauthUnauthorized, message, nil)
}

func sendHTTPError(c *gin.Context,  state int, message string) {
	c.String(state, "%s", message)
}