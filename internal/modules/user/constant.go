package user

import (
	"net/http"

	"redler/internal/pkg/lib"
)

const (
	EmailExistsErr = iota + 1
	EmailNotFoundErr
	UserNotFoundErr
	ParseBodyErr
	SomethingWentWrongErr
)

var HttpErrors = map[int]*lib.HttpResponseStruct{
	EmailExistsErr: lib.HttpResponse(http.StatusBadRequest).Errors(lib.H{
		"email": "Email already exists",
	}),
	EmailNotFoundErr: lib.HttpResponse(http.StatusNotFound).Errors(lib.H{
		"email": "Email not found",
	}),
	ParseBodyErr:          lib.HttpResponse(http.StatusBadRequest).Message("Unable to parse body"),
	UserNotFoundErr:       lib.HttpResponse(http.StatusNotFound).Message("User not found"),
	SomethingWentWrongErr: lib.HttpResponse(http.StatusInternalServerError).Message("Something went wrong"),
}
