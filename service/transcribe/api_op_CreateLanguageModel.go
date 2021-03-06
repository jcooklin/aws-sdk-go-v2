// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new custom language model. When creating a new language model, you
// must specify if you want a Wideband (audio sample rates over 16,000 Hz) or
// Narrowband (audio sample rates under 16,000 Hz) base model. You then include the
// S3 URI location of your training and tuning files, the language for the model, a
// unique name, and any tags you want associated with your model.
func (c *Client) CreateLanguageModel(ctx context.Context, params *CreateLanguageModelInput, optFns ...func(*Options)) (*CreateLanguageModelOutput, error) {
	if params == nil {
		params = &CreateLanguageModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLanguageModel", params, optFns, c.addOperationCreateLanguageModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLanguageModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLanguageModelInput struct {

	// The Amazon Transcribe standard language model, or base model, used to create
	// your custom language model. Amazon Transcribe offers two options for base
	// models: Wideband and Narrowband. If the audio you want to transcribe has a
	// sample rate of 16,000 Hz or greater, choose WideBand. To transcribe audio with a
	// sample rate less than 16,000 Hz, choose NarrowBand.
	//
	// This member is required.
	BaseModelName types.BaseModelName

	// Contains your data access role ARN (Amazon Resource Name) and the Amazon S3
	// locations of your training (S3Uri) and tuning (TuningDataS3Uri) data.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// The language of your custom language model; note that the language code you
	// select must match the language of your training and tuning data.
	//
	// This member is required.
	LanguageCode types.CLMLanguageCode

	// The name of your new custom language model. This name is case sensitive, cannot
	// contain spaces, and must be unique within an Amazon Web Services account. If you
	// try to create a language model with the same name as a previous language model,
	// you get a ConflictException error.
	//
	// This member is required.
	ModelName *string

	// Optionally add tags, each in the form of a key:value pair, to your new language
	// model. See also: .
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateLanguageModelOutput struct {

	// The Amazon Transcribe standard language model, or base model, you used when
	// creating your custom language model. If your audio has a sample rate of 16,000
	// Hz or greater, this value should be WideBand. If your audio has a sample rate of
	// less than 16,000 Hz, this value should be NarrowBand.
	BaseModelName types.BaseModelName

	// Lists your data access role ARN (Amazon Resource Name) and the Amazon S3
	// locations your provided for your training (S3Uri) and tuning (TuningDataS3Uri)
	// data.
	InputDataConfig *types.InputDataConfig

	// The language code you selected for your custom language model.
	LanguageCode types.CLMLanguageCode

	// The unique name you chose for your custom language model.
	ModelName *string

	// The status of your custom language model. When the status shows as COMPLETED,
	// your model is ready to use.
	ModelStatus types.ModelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLanguageModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLanguageModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLanguageModel{}, middleware.After)
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
	if err = addOpCreateLanguageModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLanguageModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLanguageModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "CreateLanguageModel",
	}
}
