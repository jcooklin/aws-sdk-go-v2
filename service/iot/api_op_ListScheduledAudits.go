// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of your scheduled audits.
func (c *Client) ListScheduledAudits(ctx context.Context, params *ListScheduledAuditsInput, optFns ...func(*Options)) (*ListScheduledAuditsOutput, error) {
	if params == nil {
		params = &ListScheduledAuditsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListScheduledAudits", params, optFns, addOperationListScheduledAuditsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListScheduledAuditsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListScheduledAuditsInput struct {

	// The maximum number of results to return at one time. The default is 25.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string
}

type ListScheduledAuditsOutput struct {

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// The list of scheduled audits.
	ScheduledAudits []types.ScheduledAuditMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListScheduledAuditsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListScheduledAudits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListScheduledAudits{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListScheduledAudits(options.Region), middleware.Before); err != nil {
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

// ListScheduledAuditsAPIClient is a client that implements the ListScheduledAudits
// operation.
type ListScheduledAuditsAPIClient interface {
	ListScheduledAudits(context.Context, *ListScheduledAuditsInput, ...func(*Options)) (*ListScheduledAuditsOutput, error)
}

var _ ListScheduledAuditsAPIClient = (*Client)(nil)

// ListScheduledAuditsPaginatorOptions is the paginator options for
// ListScheduledAudits
type ListScheduledAuditsPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 25.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListScheduledAuditsPaginator is a paginator for ListScheduledAudits
type ListScheduledAuditsPaginator struct {
	options   ListScheduledAuditsPaginatorOptions
	client    ListScheduledAuditsAPIClient
	params    *ListScheduledAuditsInput
	nextToken *string
	firstPage bool
}

// NewListScheduledAuditsPaginator returns a new ListScheduledAuditsPaginator
func NewListScheduledAuditsPaginator(client ListScheduledAuditsAPIClient, params *ListScheduledAuditsInput, optFns ...func(*ListScheduledAuditsPaginatorOptions)) *ListScheduledAuditsPaginator {
	options := ListScheduledAuditsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListScheduledAuditsInput{}
	}

	return &ListScheduledAuditsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListScheduledAuditsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListScheduledAudits page.
func (p *ListScheduledAuditsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListScheduledAuditsOutput, error) {
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

	result, err := p.client.ListScheduledAudits(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListScheduledAudits(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListScheduledAudits",
	}
}
