package debug

import (
	"net/http"
	"net/http/pprof"
)

func WithProfile() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", pprof.Index)
	mux.HandleFunc("/cmdline", pprof.Cmdline)
	mux.HandleFunc("/profile", pprof.Profile)
	mux.HandleFunc("/symbol", pprof.Symbol)
	mux.HandleFunc("/trace", pprof.Trace)

	return mux
}
