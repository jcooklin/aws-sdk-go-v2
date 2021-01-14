// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the history for the specified alarm. You can filter the results by
// date range or item type. If an alarm name is not specified, the histories for
// either all metric alarms or all composite alarms are returned. CloudWatch
// retains the history of an alarm even if you delete the alarm.
func (c *Client) DescribeAlarmHistory(ctx context.Context, params *DescribeAlarmHistoryInput, optFns ...func(*Options)) (*DescribeAlarmHistoryOutput, error) {
	if params == nil {
		params = &DescribeAlarmHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAlarmHistory", params, optFns, addOperationDescribeAlarmHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAlarmHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAlarmHistoryInput struct {

	// The name of the alarm.
	AlarmName *string

	// Use this parameter to specify whether you want the operation to return metric
	// alarms or composite alarms. If you omit this parameter, only metric alarms are
	// returned.
	AlarmTypes []types.AlarmType

	// The ending date to retrieve alarm history.
	EndDate *time.Time

	// The type of alarm histories to retrieve.
	HistoryItemType types.HistoryItemType

	// The maximum number of alarm history records to retrieve.
	MaxRecords *int32

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string

	// Specified whether to return the newest or oldest alarm history first. Specify
	// TimestampDescending to have the newest event history returned first, and specify
	// TimestampAscending to have the oldest history returned first.
	ScanBy types.ScanBy

	// The starting date to retrieve alarm history.
	StartDate *time.Time
}

type DescribeAlarmHistoryOutput struct {

	// The alarm histories, in JSON format.
	AlarmHistoryItems []types.AlarmHistoryItem

	// The token that marks the start of the next batch of returned results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAlarmHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAlarmHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAlarmHistory{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAlarmHistory(options.Region), middleware.Before); err != nil {
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

// DescribeAlarmHistoryAPIClient is a client that implements the
// DescribeAlarmHistory operation.
type DescribeAlarmHistoryAPIClient interface {
	DescribeAlarmHistory(context.Context, *DescribeAlarmHistoryInput, ...func(*Options)) (*DescribeAlarmHistoryOutput, error)
}

var _ DescribeAlarmHistoryAPIClient = (*Client)(nil)

// DescribeAlarmHistoryPaginatorOptions is the paginator options for
// DescribeAlarmHistory
type DescribeAlarmHistoryPaginatorOptions struct {
	// The maximum number of alarm history records to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAlarmHistoryPaginator is a paginator for DescribeAlarmHistory
type DescribeAlarmHistoryPaginator struct {
	options   DescribeAlarmHistoryPaginatorOptions
	client    DescribeAlarmHistoryAPIClient
	params    *DescribeAlarmHistoryInput
	nextToken *string
	firstPage bool
}

// NewDescribeAlarmHistoryPaginator returns a new DescribeAlarmHistoryPaginator
func NewDescribeAlarmHistoryPaginator(client DescribeAlarmHistoryAPIClient, params *DescribeAlarmHistoryInput, optFns ...func(*DescribeAlarmHistoryPaginatorOptions)) *DescribeAlarmHistoryPaginator {
	options := DescribeAlarmHistoryPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeAlarmHistoryInput{}
	}

	return &DescribeAlarmHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAlarmHistoryPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeAlarmHistory page.
func (p *DescribeAlarmHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAlarmHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeAlarmHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAlarmHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "DescribeAlarmHistory",
	}
}
