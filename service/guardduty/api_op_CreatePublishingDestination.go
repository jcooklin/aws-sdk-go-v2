// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a publishing destination to export findings to. The resource to export
// findings to must exist before you use this operation.
func (c *Client) CreatePublishingDestination(ctx context.Context, params *CreatePublishingDestinationInput, optFns ...func(*Options)) (*CreatePublishingDestinationOutput, error) {
	if params == nil {
		params = &CreatePublishingDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePublishingDestination", params, optFns, addOperationCreatePublishingDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePublishingDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePublishingDestinationInput struct {

	// The properties of the publishing destination, including the ARNs for the
	// destination and the KMS key used for encryption.
	//
	// This member is required.
	DestinationProperties *types.DestinationProperties

	// The type of resource for the publishing destination. Currently only Amazon S3
	// buckets are supported.
	//
	// This member is required.
	DestinationType types.DestinationType

	// The ID of the GuardDuty detector associated with the publishing destination.
	//
	// This member is required.
	DetectorId *string

	// The idempotency token for the request.
	ClientToken *string
}

type CreatePublishingDestinationOutput struct {

	// The ID of the publishing destination that is created.
	//
	// This member is required.
	DestinationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePublishingDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePublishingDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePublishingDestination{}, middleware.After)
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
	if err = addIdempotencyToken_opCreatePublishingDestinationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePublishingDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePublishingDestination(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreatePublishingDestination struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePublishingDestination) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePublishingDestination) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePublishingDestinationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePublishingDestinationInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreatePublishingDestinationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePublishingDestination{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePublishingDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "CreatePublishingDestination",
	}
}
