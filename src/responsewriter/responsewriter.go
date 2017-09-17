package jsonapi

import (
	"encoding/json"
	"log"
	"net/http"
	//"strconv"
	"time"

	gjsonapi "github.com/google/jsonapi"
	"github.com/manyminds/api2go/jsonapi"
)

import . "github.com/projectkita/project-harapan-backend-golang/src/error"

func New(res http.ResponseWriter) *Module {
	return &Module{
		start:    time.Now(),
		response: res,
	}
}

func (m *Module) SuccessWriter(dataResponse interface{}) {
	dataBytes, err := json.Marshal(dataResponse)
	if err != nil {
		log.Fatalln(err)
	}

	m.response.Header().Set("Content-Type", "application/json")
	m.response.WriteHeader(http.StatusOK)
	m.response.Write(dataBytes)
}

func (m *Module) SuccessWriterJsonAPI(dataResponse interface{}) {
	dataBytes, err := jsonapi.Marshal(dataResponse)
	if err != nil {
		log.Fatalln(err)
	}

	m.response.Header().Set("Content-Type", "application/json")
	m.response.WriteHeader(http.StatusOK)
	m.response.Write(dataBytes)
}

func (m *Module) SuccessWriterGJsonAPI(dataResponse interface{}) {

	m.response.Header().Set("Content-Type", "application/json")
	m.response.WriteHeader(http.StatusOK)

	if err := gjsonapi.MarshalPayload(m.response, dataResponse); err != nil {
		http.Error(m.response, err.Error(), http.StatusInternalServerError)
	}
}

func (m *Module) HttpErrorWriter(lang string, err error) {
	m.HttpErrorWriterResponse(lang, err)
}

func (m *Module) HttpErrorWriterWithoutPrintError(lang string, err error) {
	m.HttpErrorWriterResponse(lang, err)
}

func (m *Module) HttpErrorWriterResponse(lang string, err error) {
	appErr, ok := err.(AppError)
	if ok && appErr.HttpStatus != 0 {
		m.response.WriteHeader(appErr.HttpStatus)
	}
	m.response.Header().Set("Content-Type", "application/json")
	m.response.Write(MarshalHTTPError(lang, err))
}

func (m *Module) ErrorWriter(lang string, err error) {
	m.response.Header().Set("Content-Type", "application/json")

	apiError := ParseError(lang, err)
	jsonData, _ := jsonapi.MarshalWithURLs(apiError, nil)

	m.response.Write(jsonData)
}
