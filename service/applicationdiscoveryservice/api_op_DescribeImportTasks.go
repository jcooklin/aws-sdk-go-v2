// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of import tasks for your account, including status information,
// times, IDs, the Amazon S3 Object URL for the import file, and more.
func (c *Client) DescribeImportTasks(ctx context.Context, params *DescribeImportTasksInput, optFns ...func(*Options)) (*DescribeImportTasksOutput, error) {
	if params == nil {
		params = &DescribeImportTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImportTasks", params, optFns, addOperationDescribeImportTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImportTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImportTasksInput struct {

	// An array of name-value pairs that you provide to filter the results for the
	// DescribeImportTask request to a specific subset of results. Currently, wildcard
	// values aren't supported for filters.
	Filters []types.ImportTaskFilter

	// The maximum number of results that you want this request to return, up to 100.
	MaxResults *int32

	// The token to request a specific page of results.
	NextToken *string
}

type DescribeImportTasksOutput struct {

	// The token to request the next page of results.
	NextToken *string

	// A returned array of import tasks that match any applied filters, up to the
	// specified number of maximum results.
	Tasks []types.ImportTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeImportTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImportTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImportTasks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImportTasks(options.Region), middleware.Before); err != nil {
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

// DescribeImportTasksAPIClient is a client that implements the DescribeImportTasks
// operation.
type DescribeImportTasksAPIClient interface {
	DescribeImportTasks(context.Context, *DescribeImportTasksInput, ...func(*Options)) (*DescribeImportTasksOutput, error)
}

var _ DescribeImportTasksAPIClient = (*Client)(nil)

// DescribeImportTasksPaginatorOptions is the paginator options for
// DescribeImportTasks
type DescribeImportTasksPaginatorOptions struct {
	// The maximum number of results that you want this request to return, up to 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeImportTasksPaginator is a paginator for DescribeImportTasks
type DescribeImportTasksPaginator struct {
	options   DescribeImportTasksPaginatorOptions
	client    DescribeImportTasksAPIClient
	params    *DescribeImportTasksInput
	nextToken *string
	firstPage bool
}

// NewDescribeImportTasksPaginator returns a new DescribeImportTasksPaginator
func NewDescribeImportTasksPaginator(client DescribeImportTasksAPIClient, params *DescribeImportTasksInput, optFns ...func(*DescribeImportTasksPaginatorOptions)) *DescribeImportTasksPaginator {
	options := DescribeImportTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeImportTasksInput{}
	}

	return &DescribeImportTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeImportTasksPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeImportTasks page.
func (p *DescribeImportTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeImportTasksOutput, error) {
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

	result, err := p.client.DescribeImportTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeImportTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DescribeImportTasks",
	}
}
