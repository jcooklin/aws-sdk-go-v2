// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an input
func (c *Client) CreateInput(ctx context.Context, params *CreateInputInput, optFns ...func(*Options)) (*CreateInputOutput, error) {
	if params == nil {
		params = &CreateInputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInput", params, optFns, addOperationCreateInputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The name of the input
type CreateInputInput struct {

	// Destination settings for PUSH type inputs.
	Destinations []types.InputDestinationRequest

	// Settings for the devices.
	InputDevices []types.InputDeviceSettings

	// A list of security groups referenced by IDs to attach to the input.
	InputSecurityGroups []string

	// A list of the MediaConnect Flows that you want to use in this input. You can
	// specify as few as one Flow and presently, as many as two. The only requirement
	// is when you have more than one is that each Flow is in a separate Availability
	// Zone as this ensures your EML input is redundant to AZ issues.
	MediaConnectFlows []types.MediaConnectFlowRequest

	// Name of the input.
	Name *string

	// Unique identifier of the request to ensure the request is handled exactly once
	// in case of retries.
	RequestId *string

	// The Amazon Resource Name (ARN) of the role this input assumes during and after
	// creation.
	RoleArn *string

	// The source URLs for a PULL-type input. Every PULL type input needs exactly two
	// source URLs for redundancy. Only specify sources for PULL type Inputs. Leave
	// Destinations empty.
	Sources []types.InputSourceRequest

	// A collection of key-value pairs.
	Tags map[string]string

	// Placeholder documentation for InputType
	Type types.InputType

	// Settings for a private VPC Input. When this property is specified, the input
	// destination addresses will be created in a VPC rather than with public Internet
	// addresses. This property requires setting the roleArn property on Input
	// creation. Not compatible with the inputSecurityGroups property.
	Vpc *types.InputVpcRequest
}

// Placeholder documentation for CreateInputResponse
type CreateInputOutput struct {

	// Placeholder documentation for Input
	Input *types.Input

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateInputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateInput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateInput{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateInputMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateInputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInput(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateInput struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateInput) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateInput) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateInputInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateInputInput ")
	}

	if input.RequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.RequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateInputMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateInput{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateInput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "CreateInput",
	}
}
