// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of certificate authorities added for the current account and
// Region.
func (c *Client) ListWebsiteCertificateAuthorities(ctx context.Context, params *ListWebsiteCertificateAuthoritiesInput, optFns ...func(*Options)) (*ListWebsiteCertificateAuthoritiesOutput, error) {
	if params == nil {
		params = &ListWebsiteCertificateAuthoritiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWebsiteCertificateAuthorities", params, optFns, addOperationListWebsiteCertificateAuthoritiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWebsiteCertificateAuthoritiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWebsiteCertificateAuthoritiesInput struct {

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	// The maximum number of results to be included in the next page.
	MaxResults *int32

	// The pagination token used to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string
}

type ListWebsiteCertificateAuthoritiesOutput struct {

	// The pagination token used to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// Information about the certificates.
	WebsiteCertificateAuthorities []types.WebsiteCaSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListWebsiteCertificateAuthoritiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWebsiteCertificateAuthorities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWebsiteCertificateAuthorities{}, middleware.After)
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
	if err = addOpListWebsiteCertificateAuthoritiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWebsiteCertificateAuthorities(options.Region), middleware.Before); err != nil {
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

// ListWebsiteCertificateAuthoritiesAPIClient is a client that implements the
// ListWebsiteCertificateAuthorities operation.
type ListWebsiteCertificateAuthoritiesAPIClient interface {
	ListWebsiteCertificateAuthorities(context.Context, *ListWebsiteCertificateAuthoritiesInput, ...func(*Options)) (*ListWebsiteCertificateAuthoritiesOutput, error)
}

var _ ListWebsiteCertificateAuthoritiesAPIClient = (*Client)(nil)

// ListWebsiteCertificateAuthoritiesPaginatorOptions is the paginator options for
// ListWebsiteCertificateAuthorities
type ListWebsiteCertificateAuthoritiesPaginatorOptions struct {
	// The maximum number of results to be included in the next page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWebsiteCertificateAuthoritiesPaginator is a paginator for
// ListWebsiteCertificateAuthorities
type ListWebsiteCertificateAuthoritiesPaginator struct {
	options   ListWebsiteCertificateAuthoritiesPaginatorOptions
	client    ListWebsiteCertificateAuthoritiesAPIClient
	params    *ListWebsiteCertificateAuthoritiesInput
	nextToken *string
	firstPage bool
}

// NewListWebsiteCertificateAuthoritiesPaginator returns a new
// ListWebsiteCertificateAuthoritiesPaginator
func NewListWebsiteCertificateAuthoritiesPaginator(client ListWebsiteCertificateAuthoritiesAPIClient, params *ListWebsiteCertificateAuthoritiesInput, optFns ...func(*ListWebsiteCertificateAuthoritiesPaginatorOptions)) *ListWebsiteCertificateAuthoritiesPaginator {
	options := ListWebsiteCertificateAuthoritiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListWebsiteCertificateAuthoritiesInput{}
	}

	return &ListWebsiteCertificateAuthoritiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWebsiteCertificateAuthoritiesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListWebsiteCertificateAuthorities page.
func (p *ListWebsiteCertificateAuthoritiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWebsiteCertificateAuthoritiesOutput, error) {
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

	result, err := p.client.ListWebsiteCertificateAuthorities(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWebsiteCertificateAuthorities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "ListWebsiteCertificateAuthorities",
	}
}
