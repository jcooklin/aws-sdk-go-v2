// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the listeners for a custom routing accelerator.
func (c *Client) ListCustomRoutingListeners(ctx context.Context, params *ListCustomRoutingListenersInput, optFns ...func(*Options)) (*ListCustomRoutingListenersOutput, error) {
	if params == nil {
		params = &ListCustomRoutingListenersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomRoutingListeners", params, optFns, addOperationListCustomRoutingListenersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomRoutingListenersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomRoutingListenersInput struct {

	// The Amazon Resource Name (ARN) of the accelerator to list listeners for.
	//
	// This member is required.
	AcceleratorArn *string

	// The number of listener objects that you want to return with this call. The
	// default value is 10.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string
}

type ListCustomRoutingListenersOutput struct {

	// The list of listeners for a custom routing accelerator.
	Listeners []types.CustomRoutingListener

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCustomRoutingListenersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCustomRoutingListeners{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCustomRoutingListeners{}, middleware.After)
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
	if err = addOpListCustomRoutingListenersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomRoutingListeners(options.Region), middleware.Before); err != nil {
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

// ListCustomRoutingListenersAPIClient is a client that implements the
// ListCustomRoutingListeners operation.
type ListCustomRoutingListenersAPIClient interface {
	ListCustomRoutingListeners(context.Context, *ListCustomRoutingListenersInput, ...func(*Options)) (*ListCustomRoutingListenersOutput, error)
}

var _ ListCustomRoutingListenersAPIClient = (*Client)(nil)

// ListCustomRoutingListenersPaginatorOptions is the paginator options for
// ListCustomRoutingListeners
type ListCustomRoutingListenersPaginatorOptions struct {
	// The number of listener objects that you want to return with this call. The
	// default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomRoutingListenersPaginator is a paginator for
// ListCustomRoutingListeners
type ListCustomRoutingListenersPaginator struct {
	options   ListCustomRoutingListenersPaginatorOptions
	client    ListCustomRoutingListenersAPIClient
	params    *ListCustomRoutingListenersInput
	nextToken *string
	firstPage bool
}

// NewListCustomRoutingListenersPaginator returns a new
// ListCustomRoutingListenersPaginator
func NewListCustomRoutingListenersPaginator(client ListCustomRoutingListenersAPIClient, params *ListCustomRoutingListenersInput, optFns ...func(*ListCustomRoutingListenersPaginatorOptions)) *ListCustomRoutingListenersPaginator {
	options := ListCustomRoutingListenersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListCustomRoutingListenersInput{}
	}

	return &ListCustomRoutingListenersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomRoutingListenersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListCustomRoutingListeners page.
func (p *ListCustomRoutingListenersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomRoutingListenersOutput, error) {
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

	result, err := p.client.ListCustomRoutingListeners(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomRoutingListeners(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "ListCustomRoutingListeners",
	}
}
