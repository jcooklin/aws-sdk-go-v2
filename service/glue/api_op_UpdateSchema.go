// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the description, compatibility setting, or version checkpoint for a
// schema set. For updating the compatibility setting, the call will not validate
// compatibility for the entire set of schema versions with the new compatibility
// setting. If the value for Compatibility is provided, the VersionNumber (a
// checkpoint) is also required. The API will validate the checkpoint version
// number for consistency. If the value for the VersionNumber (checkpoint) is
// provided, Compatibility is optional and this can be used to set/reset a
// checkpoint for the schema. This update will happen only if the schema is in the
// AVAILABLE state.
func (c *Client) UpdateSchema(ctx context.Context, params *UpdateSchemaInput, optFns ...func(*Options)) (*UpdateSchemaOutput, error) {
	if params == nil {
		params = &UpdateSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSchema", params, optFns, addOperationUpdateSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSchemaInput struct {

	// This is a wrapper structure to contain schema identity fields. The structure
	// contains:
	//
	// * SchemaId$SchemaArn: The Amazon Resource Name (ARN) of the schema.
	// One of SchemaArn or SchemaName has to be provided.
	//
	// * SchemaId$SchemaName: The
	// name of the schema. One of SchemaArn or SchemaName has to be provided.
	//
	// This member is required.
	SchemaId *types.SchemaId

	// The new compatibility setting for the schema.
	Compatibility types.Compatibility

	// The new description for the schema.
	Description *string

	// Version number required for check pointing. One of VersionNumber or
	// Compatibility has to be provided.
	SchemaVersionNumber *types.SchemaVersionNumber
}

type UpdateSchemaOutput struct {

	// The name of the registry that contains the schema.
	RegistryName *string

	// The Amazon Resource Name (ARN) of the schema.
	SchemaArn *string

	// The name of the schema.
	SchemaName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSchema{}, middleware.After)
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
	if err = addOpUpdateSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSchema(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateSchema",
	}
}
