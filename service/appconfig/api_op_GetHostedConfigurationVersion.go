// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get information about a specific configuration version.
func (c *Client) GetHostedConfigurationVersion(ctx context.Context, params *GetHostedConfigurationVersionInput, optFns ...func(*Options)) (*GetHostedConfigurationVersionOutput, error) {
	if params == nil {
		params = &GetHostedConfigurationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHostedConfigurationVersion", params, optFns, addOperationGetHostedConfigurationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHostedConfigurationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetHostedConfigurationVersionInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// The configuration profile ID.
	//
	// This member is required.
	ConfigurationProfileId *string

	// The version.
	//
	// This member is required.
	VersionNumber int32
}

type GetHostedConfigurationVersionOutput struct {

	// The application ID.
	ApplicationId *string

	// The configuration profile ID.
	ConfigurationProfileId *string

	// The content of the configuration or the configuration data.
	Content []byte

	// A standard MIME type describing the format of the configuration content. For
	// more information, see Content-Type
	// (https://docs.aws.amazon.com/https:/www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType *string

	// A description of the configuration.
	Description *string

	// The configuration version.
	VersionNumber int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetHostedConfigurationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetHostedConfigurationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetHostedConfigurationVersion{}, middleware.After)
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
	if err = addOpGetHostedConfigurationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHostedConfigurationVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetHostedConfigurationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "GetHostedConfigurationVersion",
	}
}
