// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists dashboards in an AWS account.
func (c *Client) ListDashboards(ctx context.Context, params *ListDashboardsInput, optFns ...func(*Options)) (*ListDashboardsOutput, error) {
	if params == nil {
		params = &ListDashboardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDashboards", params, optFns, addOperationListDashboardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDashboardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDashboardsInput struct {

	// The ID of the AWS account that contains the dashboards that you're listing.
	//
	// This member is required.
	AwsAccountId *string

	// The maximum number of results to be returned per request.
	MaxResults int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string
}

type ListDashboardsOutput struct {

	// A structure that contains all of the dashboards in your AWS account. This
	// structure provides basic information about the dashboards.
	DashboardSummaryList []types.DashboardSummary

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDashboardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDashboards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDashboards{}, middleware.After)
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
	if err = addOpListDashboardsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDashboards(options.Region), middleware.Before); err != nil {
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

// ListDashboardsAPIClient is a client that implements the ListDashboards
// operation.
type ListDashboardsAPIClient interface {
	ListDashboards(context.Context, *ListDashboardsInput, ...func(*Options)) (*ListDashboardsOutput, error)
}

var _ ListDashboardsAPIClient = (*Client)(nil)

// ListDashboardsPaginatorOptions is the paginator options for ListDashboards
type ListDashboardsPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDashboardsPaginator is a paginator for ListDashboards
type ListDashboardsPaginator struct {
	options   ListDashboardsPaginatorOptions
	client    ListDashboardsAPIClient
	params    *ListDashboardsInput
	nextToken *string
	firstPage bool
}

// NewListDashboardsPaginator returns a new ListDashboardsPaginator
func NewListDashboardsPaginator(client ListDashboardsAPIClient, params *ListDashboardsInput, optFns ...func(*ListDashboardsPaginatorOptions)) *ListDashboardsPaginator {
	options := ListDashboardsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListDashboardsInput{}
	}

	return &ListDashboardsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDashboardsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDashboards page.
func (p *ListDashboardsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDashboardsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDashboards(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDashboards(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListDashboards",
	}
}
