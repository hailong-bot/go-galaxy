package middleware

import "net/http"

type LoginVerificationMiddleware struct {
}

func NewLoginVerificationMiddleware() *LoginVerificationMiddleware {
	return &LoginVerificationMiddleware{}
}

func (m *LoginVerificationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		if r.Header.Get("token") == "123456" {
			// Passthrough to next handler if need
			next(w, r)
			return
		}
		w.Write([]byte("权限不足，无法进行"))
		return
	}
}
