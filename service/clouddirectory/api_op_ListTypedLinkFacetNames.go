// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of TypedLink facet names for a particular schema. For
// more information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) ListTypedLinkFacetNames(ctx context.Context, params *ListTypedLinkFacetNamesInput, optFns ...func(*Options)) (*ListTypedLinkFacetNamesOutput, error) {
	if params == nil {
		params = &ListTypedLinkFacetNamesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTypedLinkFacetNames", params, optFns, addOperationListTypedLinkFacetNamesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTypedLinkFacetNamesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypedLinkFacetNamesInput struct {

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns.
	//
	// This member is required.
	SchemaArn *string

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type ListTypedLinkFacetNamesOutput struct {

	// The names of typed link facets that exist within the schema.
	FacetNames []string

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTypedLinkFacetNamesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTypedLinkFacetNames{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTypedLinkFacetNames{}, middleware.After)
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
	if err = addOpListTypedLinkFacetNamesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTypedLinkFacetNames(options.Region), middleware.Before); err != nil {
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

// ListTypedLinkFacetNamesAPIClient is a client that implements the
// ListTypedLinkFacetNames operation.
type ListTypedLinkFacetNamesAPIClient interface {
	ListTypedLinkFacetNames(context.Context, *ListTypedLinkFacetNamesInput, ...func(*Options)) (*ListTypedLinkFacetNamesOutput, error)
}

var _ ListTypedLinkFacetNamesAPIClient = (*Client)(nil)

// ListTypedLinkFacetNamesPaginatorOptions is the paginator options for
// ListTypedLinkFacetNames
type ListTypedLinkFacetNamesPaginatorOptions struct {
	// The maximum number of results to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTypedLinkFacetNamesPaginator is a paginator for ListTypedLinkFacetNames
type ListTypedLinkFacetNamesPaginator struct {
	options   ListTypedLinkFacetNamesPaginatorOptions
	client    ListTypedLinkFacetNamesAPIClient
	params    *ListTypedLinkFacetNamesInput
	nextToken *string
	firstPage bool
}

// NewListTypedLinkFacetNamesPaginator returns a new
// ListTypedLinkFacetNamesPaginator
func NewListTypedLinkFacetNamesPaginator(client ListTypedLinkFacetNamesAPIClient, params *ListTypedLinkFacetNamesInput, optFns ...func(*ListTypedLinkFacetNamesPaginatorOptions)) *ListTypedLinkFacetNamesPaginator {
	options := ListTypedLinkFacetNamesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListTypedLinkFacetNamesInput{}
	}

	return &ListTypedLinkFacetNamesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTypedLinkFacetNamesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTypedLinkFacetNames page.
func (p *ListTypedLinkFacetNamesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTypedLinkFacetNamesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListTypedLinkFacetNames(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListTypedLinkFacetNames(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListTypedLinkFacetNames",
	}
}
