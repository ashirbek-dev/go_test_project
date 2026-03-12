package http

import (
	"gateway/core/app"
	"gateway/core/app/token/queries"
	postgresrepo "gateway/infrastructure/storage/postgres/repositories"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthErrorResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      interface{} `json:"id"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func BasicAuthMiddleware(appService *app.ApplicationService) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
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

		token := strings.TrimPrefix(authHeader, "Bearer ")

		tokenService := appService.GetTokenService(postgresrepo.Token{})
		res, err := tokenService.Queries.ValidateToken.Handle(queries.ValidateTokenQuery{Token: token})
		if err != nil || res == nil || !res.IsValid {
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

		c.Set("service_id", res.ServiceId)
		c.Next()
	}
}
