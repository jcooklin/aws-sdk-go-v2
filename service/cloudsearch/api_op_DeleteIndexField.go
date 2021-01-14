// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes an IndexField from the search domain. For more information, see
// Configuring Index Fields
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-index-fields.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DeleteIndexField(ctx context.Context, params *DeleteIndexFieldInput, optFns ...func(*Options)) (*DeleteIndexFieldOutput, error) {
	if params == nil {
		params = &DeleteIndexFieldInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIndexField", params, optFns, addOperationDeleteIndexFieldMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIndexFieldOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DeleteIndexField operation. Specifies the
// name of the domain you want to update and the name of the index field you want
// to delete.
type DeleteIndexFieldInput struct {

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string

	// The name of the index field your want to remove from the domain's indexing
	// options.
	//
	// This member is required.
	IndexFieldName *string
}

// The result of a DeleteIndexField request.
type DeleteIndexFieldOutput struct {

	// The status of the index field being deleted.
	//
	// This member is required.
	IndexField *types.IndexFieldStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteIndexFieldMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteIndexField{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteIndexField{}, middleware.After)
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
	if err = addOpDeleteIndexFieldValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIndexField(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteIndexField(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DeleteIndexField",
	}
}
