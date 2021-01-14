// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the resource counts across accounts and regions that are present in your
// AWS Config aggregator. You can request the resource counts by providing filters
// and GroupByKey. For example, if the input contains accountID 12345678910 and
// region us-east-1 in filters, the API returns the count of resources in account
// ID 12345678910 and region us-east-1. If the input contains ACCOUNT_ID as a
// GroupByKey, the API returns resource counts for all source accounts that are
// present in your aggregator.
func (c *Client) GetAggregateDiscoveredResourceCounts(ctx context.Context, params *GetAggregateDiscoveredResourceCountsInput, optFns ...func(*Options)) (*GetAggregateDiscoveredResourceCountsOutput, error) {
	if params == nil {
		params = &GetAggregateDiscoveredResourceCountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAggregateDiscoveredResourceCounts", params, optFns, addOperationGetAggregateDiscoveredResourceCountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAggregateDiscoveredResourceCountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAggregateDiscoveredResourceCountsInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// Filters the results based on the ResourceCountFilters object.
	Filters *types.ResourceCountFilters

	// The key to group the resource counts.
	GroupByKey types.ResourceCountGroupKey

	// The maximum number of GroupedResourceCount objects returned on each page. The
	// default is 1000. You cannot specify a number greater than 1000. If you specify
	// 0, AWS Config uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

type GetAggregateDiscoveredResourceCountsOutput struct {

	// The total number of resources that are present in an aggregator with the filters
	// that you provide.
	//
	// This member is required.
	TotalDiscoveredResources int64

	// The key passed into the request object. If GroupByKey is not provided, the
	// result will be empty.
	GroupByKey *string

	// Returns a list of GroupedResourceCount objects.
	GroupedResourceCounts []types.GroupedResourceCount

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAggregateDiscoveredResourceCountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAggregateDiscoveredResourceCounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAggregateDiscoveredResourceCounts{}, middleware.After)
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
	if err = addOpGetAggregateDiscoveredResourceCountsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAggregateDiscoveredResourceCounts(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAggregateDiscoveredResourceCounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetAggregateDiscoveredResourceCounts",
	}
}
