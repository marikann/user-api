package errors

import (
	"net/http"
	"user-api/pkg/errors"
)

const repoOp = "Repository"
const handlerOp = "Handler"
const serviceOp = "Service"

var (
	InvalidCreateUserBody = errors.New(handlerOp, "Invalid User Body", 1, http.StatusBadRequest)
	UserAlreadyExist      = errors.New(serviceOp, "User with that email already exists", 2, http.StatusConflict)
	InsertUser            = errors.New(repoOp, "Insert User Error ", 3, http.StatusInternalServerError)
	UUIDParseError        = errors.New(handlerOp, "UUID Parse ERROR ", 4, http.StatusBadRequest)
	UserNotFound          = errors.New(serviceOp, "User Not Found", 5, http.StatusNotFound)
	FindOne               = errors.New(repoOp, "Find One Error", 6, http.StatusInternalServerError)
	InvalidUpdateUserBody = errors.New(handlerOp, "Invalid User Body", 7, http.StatusBadRequest)
	Cursor                = errors.New(repoOp, "Cursor Error", 8, http.StatusInternalServerError)
	CountDocuments        = errors.New(repoOp, "Count Documents Error", 9, http.StatusInternalServerError)
	FindOneAndUpdate      = errors.New(repoOp, "Find One And Update Error", 10, http.StatusInternalServerError)
	FindOneAndDelete      = errors.New(repoOp, "Find One And Delete Error", 11, http.StatusInternalServerError)
	CursorAll             = errors.New(repoOp, "Cursor All Error", 12, http.StatusInternalServerError)
)
