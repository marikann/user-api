package errors

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Public     PublicError
	StatusCode int
}

type PublicError struct {
	Op        string
	Desc      string
	ErrorCode int
}

func (e *Error) Log() {
	fields := logrus.Fields{
		"StatusCode": e.StatusCode,
		"Op":         e.Public.Op,
		"ErrorCode":  e.Public.ErrorCode,
	}
	if e.StatusCode >= 400 && e.StatusCode < 500 {
		logrus.WithFields(fields).Info(e.Public.Desc)
	} else if e.StatusCode >= 500 {
		logrus.WithFields(fields).Error(e.Public.Desc)
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Operation: %s, Description: %s, ErrorCode: %d", e.Public.Op, e.Public.Desc, e.Public.ErrorCode)
}

func New(op string, desc string, errorCode int, statusCode int) *Error {
	return &Error{Public: PublicError{
		Op:        op,
		Desc:      desc,
		ErrorCode: errorCode,
	}, StatusCode: statusCode}
}

func (e *Error) WrapDesc(desc string) *Error {
	return &Error{Public: PublicError{
		Op:        e.Public.Op,
		Desc:      fmt.Sprintf("Error Message: %s , Internal: %s", e.Public.Desc, desc),
		ErrorCode: e.Public.ErrorCode,
	}, StatusCode: e.StatusCode}
}
