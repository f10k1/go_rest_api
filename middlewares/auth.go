package middlewares

import (
	"context"
	"net/http"
	"rest_api/database"
	"strings"
)

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var clientId = request.URL.Query().Get("id")
		clientProfile, ok := database.Database[clientId]
		if !ok || clientId == "" {
			http.Error(writer, "Forbidden", http.StatusForbidden)
			return
		}

		token := request.Header.Get("Authorization")
		if !isValidToken(clientProfile, token) {
			http.Error(writer, "Forbidden", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(request.Context(), "clientProfile", clientProfile)
		request = request.WithContext(ctx)

		next.ServeHTTP(writer, request)
	}
}

func isValidToken(clientProfile database.ClientProfile, token string) bool {
	if strings.HasPrefix(token, "Bearer") {
		return strings.TrimPrefix(token, "Bearer ") == clientProfile.Token
	}

	return false
}
