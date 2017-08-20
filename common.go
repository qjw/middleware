package middleware

import "net/http"

type ServeHTTP func(http.ResponseWriter, *http.Request)(bool)
