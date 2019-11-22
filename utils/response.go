package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// printDebugf behaves like log.Printf only in the debug env
func printDebugf(format string, args ...interface{}) {
	if env := os.Getenv("GO_SERVER_DEBUG"); len(env) != 0 {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

// DefaultResponse is the default response template
type DefaultResponse struct {
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
	Code    int32       `json:"code"`
}

// EntriesResponse response
type EntriesResponse struct {
	Found int64       `json:"found"`
	Data  interface{} `json:"data"`
}

func (e *DefaultResponse) String() string {
	return fmt.Sprintf("message: %s, code: %d", e.Message, e.Code)
}

// Respond is response write to ResponseWriter
func Respond(w http.ResponseWriter, code int32, src interface{}) {
	var body []byte
	var err error

	switch s := src.(type) {
	case []byte:
		if !json.Valid(s) {
			Error(w, http.StatusInternalServerError, "invalid json")
			return
		}
		body = s
	case string:
		body = []byte(s)
	case *DefaultResponse, DefaultResponse:
		// avoid infinite loop
		if body, err = json.Marshal(src); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprint("{\"message\":\"failed to parse json\",\"result\": \"\", \"code\": $1}", code)))
			return
		}
	default:
		if body, err = json.Marshal(src); err != nil {
			Error(w, http.StatusInternalServerError, "failed to parse json")
			return
		}
	}
	w.WriteHeader(int(code))
	w.Write(body)
}

// Error is wrapped Respond when error response
func Error(w http.ResponseWriter, code int32, msg string) {
	e := &DefaultResponse{
		Message: msg,
		Code:    code,
		Result:  "",
	}
	printDebugf("%s", e.String())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, e)
}

// JSON is wrapped Respond when success response
func JSON(w http.ResponseWriter, code int32, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json := &DefaultResponse{
		Message: "success",
		Code:    code,
		Result:  src,
	}
	Respond(w, code, json)
}
