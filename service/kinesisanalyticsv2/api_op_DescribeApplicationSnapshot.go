// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a snapshot of application state data.
func (c *Client) DescribeApplicationSnapshot(ctx context.Context, params *DescribeApplicationSnapshotInput, optFns ...func(*Options)) (*DescribeApplicationSnapshotOutput, error) {
	if params == nil {
		params = &DescribeApplicationSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApplicationSnapshot", params, optFns, addOperationDescribeApplicationSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeApplicationSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeApplicationSnapshotInput struct {

	// The name of an existing application.
	//
	// This member is required.
	ApplicationName *string

	// The identifier of an application snapshot. You can retrieve this value using .
	//
	// This member is required.
	SnapshotName *string
}

type DescribeApplicationSnapshotOutput struct {

	// An object containing information about the application snapshot.
	//
	// This member is required.
	SnapshotDetails *types.SnapshotDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeApplicationSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeApplicationSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeApplicationSnapshot{}, middleware.After)
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
	if err = addOpDescribeApplicationSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplicationSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeApplicationSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DescribeApplicationSnapshot",
	}
}
