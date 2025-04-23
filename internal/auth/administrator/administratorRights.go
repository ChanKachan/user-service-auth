package administrator

import "net/http"

// Проверяет, есть у пользователя права админа
func administratorRightsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}
