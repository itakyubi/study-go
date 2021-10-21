package common

import "net/http"

type Code string

func (c Code) String() string {
	if msg, ok := templates[c]; ok {
		return msg
	}
	return templates[ErrUnknown]
}

const (
	ErrRequestMethodNotFound = "ErrRequestMethodNotFound"
	ErrUnknown               = "UnknownError"
)

var templates = map[Code]string{
	ErrRequestMethodNotFound: "The request method is not found.",
	ErrUnknown:               "There is a unknown error{{if .error}} ({{.error}}){{end}}. If the attempt to retry does not work, please contact us.",
}

func getHTTPStatus(c Code) int {
	switch c {
	case ErrRequestMethodNotFound:
		return http.StatusNotFound
	case ErrUnknown:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
