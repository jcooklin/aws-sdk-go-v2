// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new flow. The request must include one source. The request optionally
// can include outputs (up to 50) and entitlements (up to 50).
func (c *Client) CreateFlow(ctx context.Context, params *CreateFlowInput, optFns ...func(*Options)) (*CreateFlowOutput, error) {
	if params == nil {
		params = &CreateFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFlow", params, optFns, addOperationCreateFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new flow. The request must include one source. The request optionally
// can include outputs (up to 50) and entitlements (up to 50).
type CreateFlowInput struct {

	// The name of the flow.
	//
	// This member is required.
	Name *string

	// The Availability Zone that you want to create the flow in. These options are
	// limited to the Availability Zones within the current AWS Region.
	AvailabilityZone *string

	// The entitlements that you want to grant on a flow.
	Entitlements []types.GrantEntitlementRequest

	// The outputs that you want to add to this flow.
	Outputs []types.AddOutputRequest

	// The settings for the source of the flow.
	Source *types.SetSourceRequest

	// The settings for source failover
	SourceFailoverConfig *types.FailoverConfig

	Sources []types.SetSourceRequest

	// The VPC interfaces you want on the flow.
	VpcInterfaces []types.VpcInterfaceRequest
}

type CreateFlowOutput struct {

	// The settings for a flow, including its source, outputs, and entitlements.
	Flow *types.Flow

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFlow{}, middleware.After)
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
	if err = addOpCreateFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "CreateFlow",
	}
}
