package http

import (
	"github.com/gin-gonic/gin"
)

type AuthErrorResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      interface{} `json:"id"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func BasicAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(401, AuthErrorResponse{
				Jsonrpc: "2.0",
				Error: struct {
					Code    int    `json:"code"`
					Message string `json:"message"`
				}{Code: -32000, Message: "Unauthorized"},
			})
			c.Abort()
			return
		}

		if !isValidUser(username, password) {
			c.JSON(401, AuthErrorResponse{
				Jsonrpc: "2.0",
				Error: struct {
					Code    int    `json:"code"`
					Message string `json:"message"`
				}{Code: -32000, Message: "Authorization failed"},
			})
			c.Abort()
			return
		}

		c.Set("auth:username", username)
		c.Set("auth:password", password)
		c.Next()
	}
}

func isValidUser(username, password string) bool {

	return true
	/*user, err := models.GetUserByUsername(username)
	if err != nil || user == nil {
		return false
	}
	return username == user.Username && password == user.Password*/
}
