// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a gateway capability configuration or defines a new capability
// configuration. Each gateway capability defines data sources for a gateway. A
// capability configuration can contain multiple data source configurations. If you
// define OPC-UA sources for a gateway in the AWS IoT SiteWise console, all of your
// OPC-UA sources are stored in one capability configuration. To list all
// capability configurations for a gateway, use DescribeGateway
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeGateway.html).
func (c *Client) UpdateGatewayCapabilityConfiguration(ctx context.Context, params *UpdateGatewayCapabilityConfigurationInput, optFns ...func(*Options)) (*UpdateGatewayCapabilityConfigurationOutput, error) {
	if params == nil {
		params = &UpdateGatewayCapabilityConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGatewayCapabilityConfiguration", params, optFns, addOperationUpdateGatewayCapabilityConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGatewayCapabilityConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGatewayCapabilityConfigurationInput struct {

	// The JSON document that defines the configuration for the gateway capability. For
	// more information, see Configuring data sources (CLI)
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/configure-sources.html#configure-source-cli)
	// in the AWS IoT SiteWise User Guide.
	//
	// This member is required.
	CapabilityConfiguration *string

	// The namespace of the gateway capability configuration to be updated. For
	// example, if you configure OPC-UA sources from the AWS IoT SiteWise console, your
	// OPC-UA capability configuration has the namespace
	// iotsitewise:opcuacollector:version, where version is a number such as 1.
	//
	// This member is required.
	CapabilityNamespace *string

	// The ID of the gateway to be updated.
	//
	// This member is required.
	GatewayId *string
}

type UpdateGatewayCapabilityConfigurationOutput struct {

	// The namespace of the gateway capability.
	//
	// This member is required.
	CapabilityNamespace *string

	// The synchronization status of the capability configuration. The sync status can
	// be one of the following:
	//
	// * IN_SYNC – The gateway is running the capability
	// configuration.
	//
	// * OUT_OF_SYNC – The gateway hasn't received the capability
	// configuration.
	//
	// * SYNC_FAILED – The gateway rejected the capability
	// configuration.
	//
	// After you update a capability configuration, its sync status is
	// OUT_OF_SYNC until the gateway receives and applies or rejects the updated
	// configuration.
	//
	// This member is required.
	CapabilitySyncStatus types.CapabilitySyncStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateGatewayCapabilityConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGatewayCapabilityConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGatewayCapabilityConfiguration{}, middleware.After)
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
	if err = addEndpointPrefix_opUpdateGatewayCapabilityConfigurationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateGatewayCapabilityConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGatewayCapabilityConfiguration(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateGatewayCapabilityConfigurationMiddleware struct {
}

func (*endpointPrefix_opUpdateGatewayCapabilityConfigurationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateGatewayCapabilityConfigurationMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "edge." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opUpdateGatewayCapabilityConfigurationMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opUpdateGatewayCapabilityConfigurationMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGatewayCapabilityConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "UpdateGatewayCapabilityConfiguration",
	}
}
