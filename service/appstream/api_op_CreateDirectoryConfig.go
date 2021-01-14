// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Directory Config object in AppStream 2.0. This object includes the
// configuration information required to join fleets and image builders to
// Microsoft Active Directory domains.
func (c *Client) CreateDirectoryConfig(ctx context.Context, params *CreateDirectoryConfigInput, optFns ...func(*Options)) (*CreateDirectoryConfigOutput, error) {
	if params == nil {
		params = &CreateDirectoryConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDirectoryConfig", params, optFns, addOperationCreateDirectoryConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDirectoryConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDirectoryConfigInput struct {

	// The fully qualified name of the directory (for example, corp.example.com).
	//
	// This member is required.
	DirectoryName *string

	// The distinguished names of the organizational units for computer accounts.
	//
	// This member is required.
	OrganizationalUnitDistinguishedNames []string

	// The credentials for the service account used by the fleet or image builder to
	// connect to the directory.
	ServiceAccountCredentials *types.ServiceAccountCredentials
}

type CreateDirectoryConfigOutput struct {

	// Information about the directory configuration.
	DirectoryConfig *types.DirectoryConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDirectoryConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDirectoryConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDirectoryConfig{}, middleware.After)
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
	if err = addOpCreateDirectoryConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDirectoryConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDirectoryConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "CreateDirectoryConfig",
	}
}
