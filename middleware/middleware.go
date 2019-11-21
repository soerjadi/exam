package middleware

import (
	"fmt"
	"net/http"

	"github.com/soerjadi/exam/utils"
)

// MuxMiddleware represent the data-struct for middleware
type MuxMiddleware struct {
	// todo(*): code here
}

var (
	logger *utils.Logger
)

func init() {
	logger = utils.LogBuilder(true)
}

// LoggingMiddleware middleware that log every request
func (m *MuxMiddleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("[%s] %s", r.Method, r.RequestURI)
		logger.Debug(str)

		next.ServeHTTP(w, r)
	})
}

// InitMiddleware initialize middleware
func InitMiddleware() *MuxMiddleware {
	return &MuxMiddleware{}
}
