package appErrors

import (
	"fmt"
	"strings"
)

type Error struct {
	ErrorCode int         `json:"error_code"`
	Errors    []ErrorBody `json:"errors"`
}

type ErrorBody struct {
	Field   *string `json:"field"`
	Message string  `json:"message"`
}

func (e Error) Error() string {
	builder := new(strings.Builder)

	builder.WriteString(fmt.Sprintf("Application Error - status:{%d}\n", e.ErrorCode))

	for _, err := range e.Errors {

		builder.WriteString(fmt.Sprintf(" - %s", err.Message))
	}

	return builder.String()
}
