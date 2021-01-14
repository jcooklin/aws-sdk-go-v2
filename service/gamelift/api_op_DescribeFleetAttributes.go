// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves core properties, including configuration, status, and metadata, for a
// fleet. To get attributes for one or more fleets, provide a list of fleet IDs or
// fleet ARNs. To get attributes for all fleets, do not specify a fleet identifier.
// When requesting attributes for multiple fleets, use the pagination parameters to
// retrieve results as a set of sequential pages. If successful, a FleetAttributes
// object is returned for each fleet requested, unless the fleet identifier is not
// found. Some API operations may limit the number of fleet IDs allowed in one
// request. If a request exceeds this limit, the request fails and the error
// message includes the maximum allowed number. Learn more Setting up GameLift
// Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
// * CreateFleet
//
// * ListFleets
//
// * DeleteFleet
//
// * Describe
// fleets:
//
// * DescribeFleetAttributes
//
// * DescribeFleetCapacity
//
// *
// DescribeFleetPortSettings
//
// * DescribeFleetUtilization
//
// *
// DescribeRuntimeConfiguration
//
// * DescribeEC2InstanceLimits
//
// *
// DescribeFleetEvents
//
// * UpdateFleetAttributes
//
// * StartFleetActions or
// StopFleetActions
func (c *Client) DescribeFleetAttributes(ctx context.Context, params *DescribeFleetAttributesInput, optFns ...func(*Options)) (*DescribeFleetAttributesOutput, error) {
	if params == nil {
		params = &DescribeFleetAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetAttributes", params, optFns, addOperationDescribeFleetAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeFleetAttributesInput struct {

	// A list of unique fleet identifiers to retrieve attributes for. You can use
	// either the fleet ID or ARN value. To retrieve attributes for all current fleets,
	// do not include this parameter. If the list of fleet identifiers includes fleets
	// that don't currently exist, the request succeeds but no attributes for that
	// fleet are returned.
	FleetIds []string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is ignored when the
	// request specifies one or a list of fleet IDs.
	Limit *int32

	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value. This parameter is ignored
	// when the request specifies one or a list of fleet IDs.
	NextToken *string
}

// Represents the returned data in response to a request operation.
type DescribeFleetAttributesOutput struct {

	// A collection of objects containing attribute metadata for each requested fleet
	// ID. Attribute objects are returned only for fleets that currently exist.
	FleetAttributes []types.FleetAttributes

	// Token that indicates where to resume retrieving results on the next call to this
	// operation. If no token is returned, these results represent the end of the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFleetAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetAttributes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetAttributes(options.Region), middleware.Before); err != nil {
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

// DescribeFleetAttributesAPIClient is a client that implements the
// DescribeFleetAttributes operation.
type DescribeFleetAttributesAPIClient interface {
	DescribeFleetAttributes(context.Context, *DescribeFleetAttributesInput, ...func(*Options)) (*DescribeFleetAttributesOutput, error)
}

var _ DescribeFleetAttributesAPIClient = (*Client)(nil)

// DescribeFleetAttributesPaginatorOptions is the paginator options for
// DescribeFleetAttributes
type DescribeFleetAttributesPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is ignored when the
	// request specifies one or a list of fleet IDs.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetAttributesPaginator is a paginator for DescribeFleetAttributes
type DescribeFleetAttributesPaginator struct {
	options   DescribeFleetAttributesPaginatorOptions
	client    DescribeFleetAttributesAPIClient
	params    *DescribeFleetAttributesInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetAttributesPaginator returns a new
// DescribeFleetAttributesPaginator
func NewDescribeFleetAttributesPaginator(client DescribeFleetAttributesAPIClient, params *DescribeFleetAttributesInput, optFns ...func(*DescribeFleetAttributesPaginatorOptions)) *DescribeFleetAttributesPaginator {
	options := DescribeFleetAttributesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeFleetAttributesInput{}
	}

	return &DescribeFleetAttributesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetAttributesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeFleetAttributes page.
func (p *DescribeFleetAttributesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetAttributesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeFleetAttributes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFleetAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeFleetAttributes",
	}
}
