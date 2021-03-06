// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The request failed due to an unknown error.
type InternalServiceError struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalServiceError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceError) ErrorCode() string             { return "InternalServiceError" }
func (e *InternalServiceError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// One of the arguments provided is invalid for this request.
type InvalidArgumentException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidArgumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArgumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArgumentException) ErrorCode() string             { return "InvalidArgumentException" }
func (e *InvalidArgumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The user is not authorized to perform this request.
type NotAuthorizedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NotAuthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotAuthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotAuthorizedException) ErrorCode() string             { return "NotAuthorizedException" }
func (e *NotAuthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
