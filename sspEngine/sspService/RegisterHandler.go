package sspService

import (
	"io"
	"net/http"
)

func RegisterHandler() {

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.Write(msg)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}
