package restapi

import (
	"immortality/pkg/users"
	"net/http"
)

type AuthMiddleWare struct {
	tokens map[string]uint
}

func (m *AuthMiddleWare) Init() {
	userStore := users.UserStore{}
	userStore.Connect()
	tokens, _ := userStore.GetTokens()
	for _, token := range tokens {
		m.tokens[token.Token] = token.ID
	}
}

func (m *AuthMiddleWare) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if _, ok := m.tokens[token]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
