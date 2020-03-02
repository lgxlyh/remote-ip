package ip

import (
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strings.Split(r.RemoteAddr, ":")[0]))
}
