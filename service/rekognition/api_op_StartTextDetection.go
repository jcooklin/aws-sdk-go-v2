// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts asynchronous detection of text in a stored video. Amazon Rekognition
// Video can detect text in a video stored in an Amazon S3 bucket. Use Video to
// specify the bucket name and the filename of the video. StartTextDetection
// returns a job identifier (JobId) which you use to get the results of the
// operation. When text detection is finished, Amazon Rekognition Video publishes a
// completion status to the Amazon Simple Notification Service topic that you
// specify in NotificationChannel. To get the results of the text detection
// operation, first check that the status value published to the Amazon SNS topic
// is SUCCEEDED. if so, call GetTextDetection and pass the job identifier (JobId)
// from the initial call to StartTextDetection.
func (c *Client) StartTextDetection(ctx context.Context, params *StartTextDetectionInput, optFns ...func(*Options)) (*StartTextDetectionOutput, error) {
	if params == nil {
		params = &StartTextDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTextDetection", params, optFns, addOperationStartTextDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTextDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTextDetectionInput struct {

	// Video file stored in an Amazon S3 bucket. Amazon Rekognition video start
	// operations such as StartLabelDetection use Video to specify a video for
	// analysis. The supported file formats are .mp4, .mov and .avi.
	//
	// This member is required.
	Video *types.Video

	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartTextDetection requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidentaly started more
	// than once.
	ClientRequestToken *string

	// Optional parameters that let you set criteria the text must meet to be included
	// in your response.
	Filters *types.StartTextDetectionFilters

	// An identifier returned in the completion status published by your Amazon Simple
	// Notification Service topic. For example, you can use JobTag to group related
	// jobs and identify them in the completion notification.
	JobTag *string

	// The Amazon Simple Notification Service topic to which Amazon Rekognition
	// publishes the completion status of a video analysis operation. For more
	// information, see api-video.
	NotificationChannel *types.NotificationChannel
}

type StartTextDetectionOutput struct {

	// Identifier for the text detection job. Use JobId to identify the job in a
	// subsequent call to GetTextDetection.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartTextDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartTextDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTextDetection{}, middleware.After)
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
	if err = addOpStartTextDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTextDetection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartTextDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartTextDetection",
	}
}
