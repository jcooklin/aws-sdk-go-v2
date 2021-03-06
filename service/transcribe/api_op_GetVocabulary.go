// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a vocabulary.
func (c *Client) GetVocabulary(ctx context.Context, params *GetVocabularyInput, optFns ...func(*Options)) (*GetVocabularyOutput, error) {
	if params == nil {
		params = &GetVocabularyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVocabulary", params, optFns, c.addOperationGetVocabularyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVocabularyInput struct {

	// The name of the vocabulary to return information about. The name is case
	// sensitive.
	//
	// This member is required.
	VocabularyName *string

	noSmithyDocumentSerde
}

type GetVocabularyOutput struct {

	// The S3 location where the vocabulary is stored. Use this URI to get the contents
	// of the vocabulary. The URI is available for a limited time.
	DownloadUri *string

	// If the VocabularyState field is FAILED, this field contains information about
	// why the job failed.
	FailureReason *string

	// The language code of the vocabulary entries.
	LanguageCode types.LanguageCode

	// The date and time that the vocabulary was last modified.
	LastModifiedTime *time.Time

	// The name of the vocabulary to return.
	VocabularyName *string

	// The processing state of the vocabulary.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetVocabularyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetVocabulary{}, middleware.After)
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
	if err = addOpGetVocabularyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVocabulary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetVocabulary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "GetVocabulary",
	}
}
