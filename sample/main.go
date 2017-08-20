package main

import (
	"net/http"
	"log"
	"github.com/qjw/middleware"
	"runtime"
)

func NoCache(w http.ResponseWriter, r *http.Request) {
	middleware.NoCache()(w,r)

	pc := make([]uintptr, 10)  // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	w.Write([]byte(f.Name()))
}

func Cors(w http.ResponseWriter, r *http.Request) {
	middleware.Cors(&middleware.CorsConfig{
		AllowMethods:[]string{"POST"},
		AllowOrigins:[]string{"http://127.0.0.2"},
	})(w,r)

	pc := make([]uintptr, 10)  // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	w.Write([]byte(f.Name()))
}

func main() {
	http.HandleFunc("/nocache", NoCache)
	http.HandleFunc("/cors", Cors)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}