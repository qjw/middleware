package middleware

import (
	"time"
	"net/http"
)

func NoCache(w http.ResponseWriter, r *http.Request){
	header := w.Header()
	header.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	header.Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	header.Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
}


