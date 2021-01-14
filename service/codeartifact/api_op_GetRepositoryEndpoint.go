// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the endpoint of a repository for a specific package format. A repository
// has one endpoint for each package format:
//
// * npm
//
// * pypi
//
// * maven
//
// * nuget
func (c *Client) GetRepositoryEndpoint(ctx context.Context, params *GetRepositoryEndpointInput, optFns ...func(*Options)) (*GetRepositoryEndpointOutput, error) {
	if params == nil {
		params = &GetRepositoryEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRepositoryEndpoint", params, optFns, addOperationGetRepositoryEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRepositoryEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRepositoryEndpointInput struct {

	// The name of the domain that contains the repository.
	//
	// This member is required.
	Domain *string

	// Returns which endpoint of a repository to return. A repository has one endpoint
	// for each package format:
	//
	// * npm
	//
	// * pypi
	//
	// * maven
	//
	// * nuget
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the repository.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the AWS account that owns the domain that
	// contains the repository. It does not include dashes or spaces.
	DomainOwner *string
}

type GetRepositoryEndpointOutput struct {

	// A string that specifies the URL of the returned endpoint.
	RepositoryEndpoint *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRepositoryEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRepositoryEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRepositoryEndpoint{}, middleware.After)
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
	if err = addOpGetRepositoryEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRepositoryEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRepositoryEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "GetRepositoryEndpoint",
	}
}
