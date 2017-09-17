package jsonapi

import (
	"net/http"
	"time"
)

type Module struct {
	start    time.Time
	response http.ResponseWriter
}
