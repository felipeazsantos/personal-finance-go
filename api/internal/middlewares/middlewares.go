package middlewares

import (
	"net/http"
)

type MiddlewareRoute struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

func RegisterRoutes(routes []MiddlewareRoute) {
	for _, route := range routes {
		http.HandleFunc(route.Pattern, VerifyRoutes(route))
	}
}

func VerifyRoutes(route MiddlewareRoute) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if route.Method == r.Method {
			route.Handler.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Route not found"))
	}
}
