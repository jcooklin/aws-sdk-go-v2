// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// EnableHealthServiceAccessForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html)
// is already in progress. Wait for the action to complete before trying again. To
// get the current status, use the DescribeHealthServiceStatusForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeHealthServiceStatusForOrganization.html)
// operation.
type ConcurrentModificationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified pagination token (nextToken) is not valid.
type InvalidPaginationToken struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidPaginationToken) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPaginationToken) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPaginationToken) ErrorCode() string             { return "InvalidPaginationToken" }
func (e *InvalidPaginationToken) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified locale is not supported.
type UnsupportedLocale struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnsupportedLocale) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedLocale) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedLocale) ErrorCode() string             { return "UnsupportedLocale" }
func (e *UnsupportedLocale) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
