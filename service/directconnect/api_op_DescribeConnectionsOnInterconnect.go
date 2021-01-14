// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deprecated. Use DescribeHostedConnections instead. Lists the connections that
// have been provisioned on the specified interconnect. Intended for use by AWS
// Direct Connect Partners only.
//
// Deprecated: This operation has been deprecated.
func (c *Client) DescribeConnectionsOnInterconnect(ctx context.Context, params *DescribeConnectionsOnInterconnectInput, optFns ...func(*Options)) (*DescribeConnectionsOnInterconnectOutput, error) {
	if params == nil {
		params = &DescribeConnectionsOnInterconnectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnectionsOnInterconnect", params, optFns, addOperationDescribeConnectionsOnInterconnectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectionsOnInterconnectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectionsOnInterconnectInput struct {

	// The ID of the interconnect.
	//
	// This member is required.
	InterconnectId *string
}

type DescribeConnectionsOnInterconnectOutput struct {

	// The connections.
	Connections []types.Connection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeConnectionsOnInterconnectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConnectionsOnInterconnect{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConnectionsOnInterconnect{}, middleware.After)
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
	if err = addOpDescribeConnectionsOnInterconnectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnectionsOnInterconnect(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConnectionsOnInterconnect(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DescribeConnectionsOnInterconnect",
	}
}
