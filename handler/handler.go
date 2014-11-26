package handler

import (
	"github.com/grooveshark/golib/gslog"
	"net/http"
)

func GetPage(w http.ResponseWriter, r *http.Request) {
	gslog.Debug("HANDLER: GetPage called with header: %+v, host: %s, requestURI: %s, remoteAddr: %s", r.Header, r.Host, r.RequestURI, r.RemoteAddr)
	w.Write([]byte(r.URL.Path))
}
