package route

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/tokopedia/panics"
)

func MiddlewareCors(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		r.Header.Add("Access-Control-Allow-Origin", "*")
		r.Header.Add("Access-Control-Allow-Methods", "GET, POST, PATCH")
		r.Header.Add("Access-Control-Allow-Headers", "origin, x-requested-with, content-type")
		handler(w, r, ps)
		return
	}
}

func MiddlewareAuthorization(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		basicAuthPrefix := " "
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, basicAuthPrefix) {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		handler(w, r, ps)
	}
}

func MiddlewareAccessLog(handler httprouter.Handle) httprouter.Handle {
	return panics.CaptureHTTPRouterHandler(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		messages := fmt.Sprintf("%s %s %s", r.Host, r.Method, r.RequestURI)
		go fmt.Println(messages)
		handler(w, r, ps)
	})
}
