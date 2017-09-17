package ers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tokopedia/galadriel/src/common"
)

type AppError struct {
	ID         int    `jsonapi:"-"              json:"-"`
	Code       string `jsonapi:"code"           json:"code"`
	Title      string `jsonapi:"name=title"     json:"title"`
	Detail     string `jsonapi:"name=detail"    json:"detail"`
	DetailId   string `jsonapi:"-"`
	DetailEn   string `jsonapi:"-"`
	Messages   string `jsonapi:"name=messages"  json:"messages,omitempty"`
	Status     string `jsonapi:"name=status"    json:"status,omitempty"`
	HttpStatus int    `jsonapi:"name=http_status"    json:"http_status,omitempty"`
}

var MarshalError error = AppError{Code: "100001", Title: "Library Error", DetailEn: "Cannot marshall data.", Status: "401", HttpStatus: http.StatusBadRequest}
var UnmarshalError error = AppError{Code: "100002", Title: "Library Error", DetailEn: "Cannot unmarshall data.", Status: "401"}
var ReadDataError error = AppError{Code: "100003", Title: "Library Error", DetailEn: "Cannot read data.", Status: "401"}
var InvalidRequestParamError error = AppError{Code: "100004", Title: "Library Error", DetailEn: "Invalid request parameter", Status: "401", HttpStatus: http.StatusPreconditionFailed}
var HttpRequestError error = AppError{Code: "100005", Title: "Library Error", DetailEn: "HTTP request error", Status: "401"}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (ae AppError) GetID() string {
	return strconv.Itoa(ae.ID)
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (ae AppError) GetName() string {
	return "errors"
}

// Implement parent method
func (ae AppError) Error() string {
	return fmt.Sprintf("%s: %s - %s", ae.Code, ae.Title, ae.DetailEn)
}

func ParseError(lang string, err error) AppError {
	appError, _ := err.(AppError)

	if lang == common.LANG_EN {
		appError.Detail = fmt.Sprintf("%s", appError.DetailEn)
	} else {
		appError.Detail = fmt.Sprintf("%s", appError.DetailId)
	}

	if appError.Messages != "" {
		appError.Detail = fmt.Sprintf("%s , messages : %s", appError.Detail, appError.Messages)
	}

	return appError
}
