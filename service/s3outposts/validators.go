// Code generated by smithy-go-codegen DO NOT EDIT.

package s3outposts

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateEndpoint struct {
}

func (*validateOpCreateEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEndpoint struct {
}

func (*validateOpDeleteEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSharedEndpoints struct {
}

func (*validateOpListSharedEndpoints) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSharedEndpoints) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSharedEndpointsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSharedEndpointsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateEndpoint{}, middleware.After)
}

func addOpDeleteEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteEndpoint{}, middleware.After)
}

func addOpListSharedEndpointsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSharedEndpoints{}, middleware.After)
}

func validateOpCreateEndpointInput(v *CreateEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateEndpointInput"}
	if v.OutpostId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutpostId"))
	}
	if v.SubnetId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetId"))
	}
	if v.SecurityGroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityGroupId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEndpointInput(v *DeleteEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEndpointInput"}
	if v.EndpointId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndpointId"))
	}
	if v.OutpostId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutpostId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSharedEndpointsInput(v *ListSharedEndpointsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSharedEndpointsInput"}
	if v.OutpostId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutpostId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
