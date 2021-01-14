// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about a subscription definition version.
func (c *Client) GetSubscriptionDefinitionVersion(ctx context.Context, params *GetSubscriptionDefinitionVersionInput, optFns ...func(*Options)) (*GetSubscriptionDefinitionVersionOutput, error) {
	if params == nil {
		params = &GetSubscriptionDefinitionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSubscriptionDefinitionVersion", params, optFns, addOperationGetSubscriptionDefinitionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSubscriptionDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSubscriptionDefinitionVersionInput struct {

	// The ID of the subscription definition.
	//
	// This member is required.
	SubscriptionDefinitionId *string

	// The ID of the subscription definition version. This value maps to the
	// ''Version'' property of the corresponding ''VersionInformation'' object, which
	// is returned by ''ListSubscriptionDefinitionVersions'' requests. If the version
	// is the last one that was associated with a subscription definition, the value
	// also maps to the ''LatestVersion'' property of the corresponding
	// ''DefinitionInformation'' object.
	//
	// This member is required.
	SubscriptionDefinitionVersionId *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
}

type GetSubscriptionDefinitionVersionOutput struct {

	// The ARN of the subscription definition version.
	Arn *string

	// The time, in milliseconds since the epoch, when the subscription definition
	// version was created.
	CreationTimestamp *string

	// Information about the subscription definition version.
	Definition *types.SubscriptionDefinitionVersion

	// The ID of the subscription definition version.
	Id *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// The version of the subscription definition version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSubscriptionDefinitionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSubscriptionDefinitionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSubscriptionDefinitionVersion{}, middleware.After)
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
	if err = addOpGetSubscriptionDefinitionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSubscriptionDefinitionVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSubscriptionDefinitionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetSubscriptionDefinitionVersion",
	}
}
