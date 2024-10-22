package utils

import "net/http"

func ExtractRequestParam(r *http.Request, param string) string {
	return r.URL.Query().Get(param)
}
