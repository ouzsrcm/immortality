package middleware

import (
	"fmt"
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

		fmt.Println("token: " + token)

		userStore := users.NewUserStore()
		tokenExists, _ := userStore.TokenExists(token)
		if !tokenExists {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
