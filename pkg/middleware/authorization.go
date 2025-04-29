package middleware

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// AuthorizationMiddleware is a middleware for handling authorization.
type AuthorizationMiddleware struct {
	SkipPaths   map[string]struct{} // Paths to skip authorization check
	TokenHeader string              // Header field to read token from, e.g., "Authorization"
}

// NewAuthorizationMiddleware creates a new AuthorizationMiddleware.
// skipPaths: list of paths that don't require authorization
func NewAuthorizationMiddleware(skipPaths []string) *AuthorizationMiddleware {
	skipMap := make(map[string]struct{})
	for _, path := range skipPaths {
		skipMap[path] = struct{}{}
	}
	return &AuthorizationMiddleware{
		SkipPaths:   skipMap,
		TokenHeader: "Authorization", // 默认从 Authorization Header 取token
	}
}

// Handle implements the Middleware interface for AuthorizationMiddleware.
func (a *AuthorizationMiddleware) Handle() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil
		}
		return false, nil
	})
}

// ShouldSkip determines whether the middleware should skip a specific path.
func (a *AuthorizationMiddleware) ShouldSkip(path string) bool {
	_, ok := a.SkipPaths[path]
	return ok
}
