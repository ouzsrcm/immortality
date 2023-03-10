package middleware

import (
	"immortality/pkg/common"
	"immortality/pkg/users"
	"log"
	"net/http"
	"strings"
)

type AuthMiddleWare struct {
	tokens map[string]uint
}

func (m *AuthMiddleWare) Init() {
	log.Print("AuthMiddleWare init")

	userStore := users.NewUserStore()
	tokens, _ := userStore.GetTokens()

	m.tokens = make(map[string]uint)
	for _, token := range tokens {
		m.tokens[token.Token] = token.ID
	}
	log.Print("tokens", m.tokens)
}

func (m *AuthMiddleWare) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if common.Contains(common.Extensions(), r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		if strings.Contains(r.URL.Path, "auth") {
			next.ServeHTTP(w, r)
			return
		}

		if strings.Contains(r.URL.Path, "swagger") {
			next.ServeHTTP(w, r)
			return
		}

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
