package ers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// HTTPError is used for errors
type HTTPError struct {
	err     error
	message string
	status  int
	Errors  []AppError `json:"errors,omitempty"`
}

// NewHTTPError creates a new error with message and status code.
// `err` will be logged (but never sent to a client), `msg` will be sent and `status` is the http status code.
// `err` can be nil.
func NewHTTPError(err error, msg string, status int) HTTPError {
	return HTTPError{err: err, message: msg, status: status}
}

func MarshalHTTPError(lang string, err error) []byte {
	//newError := galError.(GaladrielError)
	newError := ParseError(lang, err)
	status, e := strconv.Atoi(newError.Status)
	if e != nil {
		status = http.StatusBadRequest
	}

	input := NewHTTPError(newError, newError.Detail, status)

	if len(input.Errors) == 0 {
		input.Errors = []AppError{{
			Code:     newError.Code,
			Title:    newError.Title,
			Detail:   newError.Detail,
			DetailId: newError.DetailId,
			DetailEn: newError.DetailEn,
		}}
	}

	data, err := json.Marshal(input)

	if err != nil {
		log.Println(err)
		return []byte("{}")
	}

	return data
}
