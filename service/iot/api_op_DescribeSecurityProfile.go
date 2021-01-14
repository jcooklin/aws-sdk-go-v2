// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a Device Defender security profile.
func (c *Client) DescribeSecurityProfile(ctx context.Context, params *DescribeSecurityProfileInput, optFns ...func(*Options)) (*DescribeSecurityProfileOutput, error) {
	if params == nil {
		params = &DescribeSecurityProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSecurityProfile", params, optFns, addOperationDescribeSecurityProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSecurityProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSecurityProfileInput struct {

	// The name of the security profile whose information you want to get.
	//
	// This member is required.
	SecurityProfileName *string
}

type DescribeSecurityProfileOutput struct {

	// Please use DescribeSecurityProfileResponse$additionalMetricsToRetainV2 instead.
	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here.
	//
	// Deprecated: Use additionalMetricsToRetainV2.
	AdditionalMetricsToRetain []string

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here.
	AdditionalMetricsToRetainV2 []types.MetricToRetain

	// Where the alerts are sent. (Alerts are always sent to the console.)
	AlertTargets map[string]types.AlertTarget

	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors []types.Behavior

	// The time the security profile was created.
	CreationDate *time.Time

	// The time the security profile was last modified.
	LastModifiedDate *time.Time

	// The ARN of the security profile.
	SecurityProfileArn *string

	// A description of the security profile (associated with the security profile when
	// it was created or updated).
	SecurityProfileDescription *string

	// The name of the security profile.
	SecurityProfileName *string

	// The version of the security profile. A new version is generated whenever the
	// security profile is updated.
	Version int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSecurityProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSecurityProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSecurityProfile{}, middleware.After)
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
	if err = addOpDescribeSecurityProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSecurityProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSecurityProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeSecurityProfile",
	}
}
