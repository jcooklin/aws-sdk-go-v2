// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads one or more pieces of evidence to the specified control in the
// assessment in AWS Audit Manager.
func (c *Client) BatchImportEvidenceToAssessmentControl(ctx context.Context, params *BatchImportEvidenceToAssessmentControlInput, optFns ...func(*Options)) (*BatchImportEvidenceToAssessmentControlOutput, error) {
	if params == nil {
		params = &BatchImportEvidenceToAssessmentControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchImportEvidenceToAssessmentControl", params, optFns, addOperationBatchImportEvidenceToAssessmentControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchImportEvidenceToAssessmentControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchImportEvidenceToAssessmentControlInput struct {

	// The identifier for the specified assessment.
	//
	// This member is required.
	AssessmentId *string

	// The identifier for the specified control.
	//
	// This member is required.
	ControlId *string

	// The identifier for the specified control set.
	//
	// This member is required.
	ControlSetId *string

	// The list of manual evidence objects.
	//
	// This member is required.
	ManualEvidence []types.ManualEvidence
}

type BatchImportEvidenceToAssessmentControlOutput struct {

	// A list of errors returned by the BatchImportEvidenceToAssessmentControl API.
	Errors []types.BatchImportEvidenceToAssessmentControlError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchImportEvidenceToAssessmentControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchImportEvidenceToAssessmentControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchImportEvidenceToAssessmentControl{}, middleware.After)
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
	if err = addOpBatchImportEvidenceToAssessmentControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchImportEvidenceToAssessmentControl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchImportEvidenceToAssessmentControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "BatchImportEvidenceToAssessmentControl",
	}
}
