package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	. "github.com/projectkita/project-harapan-backend-golang/src/common"
	. "github.com/projectkita/project-harapan-backend-golang/src/error"
	writer "github.com/projectkita/project-harapan-backend-golang/src/responsewriter"
)

func GetHelloHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	respWriter := writer.New(w)

	response := ResponseData{Message: "Hello World"}

	jsonData, err := json.Marshal(response)
	if err != nil {
		respWriter.ErrorWriter(LANG_EN, MarshalError)
		return
	}

	w.Write([]byte(jsonData))
}

func GetHelloByIDHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	respWriter := writer.New(w)

	queryValues := r.URL.Query()

	id, err := strconv.Atoi(queryValues.Get("id"))
	if err != nil {
		respWriter.ErrorWriter(LANG_EN, InvalidRequestParamError)
		return
	}

	response := ResponseData{ID: id}

	if id < 1 {
		response.Message = "Not Valid ID"
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		respWriter.ErrorWriter(LANG_EN, MarshalError)
		return
	}

	w.Write([]byte(jsonData))
}
