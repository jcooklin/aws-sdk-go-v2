// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an empty cluster. Each cluster supports five nodes. You use the
// CreateJob action separately to create the jobs for each of these nodes. The
// cluster does not ship until these five node jobs have been created.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	if params == nil {
		params = &CreateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCluster", params, optFns, addOperationCreateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {

	// The ID for the address that you want the cluster shipped to.
	//
	// This member is required.
	AddressId *string

	// The type of job for this cluster. Currently, the only job type supported for
	// clusters is LOCAL_USE.
	//
	// This member is required.
	JobType types.JobType

	// The resources associated with the cluster job. These resources include Amazon S3
	// buckets and optional AWS Lambda functions written in the Python language.
	//
	// This member is required.
	Resources *types.JobResource

	// The RoleARN that you want to associate with this cluster. RoleArn values are
	// created by using the CreateRole
	// (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) API
	// action in AWS Identity and Access Management (IAM).
	//
	// This member is required.
	RoleARN *string

	// The shipping speed for each node in this cluster. This speed doesn't dictate how
	// soon you'll get each Snowball Edge device, rather it represents how quickly each
	// device moves to its destination while in transit. Regional shipping speeds are
	// as follows:
	//
	// * In Australia, you have access to express shipping. Typically,
	// Snow devices shipped express are delivered in about a day.
	//
	// * In the European
	// Union (EU), you have access to express shipping. Typically, Snow devices shipped
	// express are delivered in about a day. In addition, most countries in the EU have
	// access to standard shipping, which typically takes less than a week, one way.
	//
	// *
	// In India, Snow device are delivered in one to seven days.
	//
	// * In the United
	// States of America (US), you have access to one-day shipping and two-day
	// shipping.
	//
	// * In Australia, you have access to express shipping. Typically,
	// devices shipped express are delivered in about a day.
	//
	// * In the European Union
	// (EU), you have access to express shipping. Typically, Snow devices shipped
	// express are delivered in about a day. In addition, most countries in the EU have
	// access to standard shipping, which typically takes less than a week, one way.
	//
	// *
	// In India, Snow device are delivered in one to seven days.
	//
	// * In the US, you have
	// access to one-day shipping and two-day shipping.
	//
	// This member is required.
	ShippingOption types.ShippingOption

	// An optional description of this specific cluster, for example Environmental Data
	// Cluster-01.
	Description *string

	// The forwarding address ID for a cluster. This field is not supported in most
	// regions.
	ForwardingAddressId *string

	// The KmsKeyARN value that you want to associate with this cluster. KmsKeyARN
	// values are created by using the CreateKey
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateKey.html) API
	// action in AWS Key Management Service (AWS KMS).
	KmsKeyARN *string

	// The Amazon Simple Notification Service (Amazon SNS) notification settings for
	// this cluster.
	Notification *types.Notification

	// The type of AWS Snow Family device to use for this cluster. For cluster jobs,
	// AWS Snow Family currently supports only the EDGE device type.
	SnowballType types.SnowballType

	// The tax documents required in your AWS Region.
	TaxDocuments *types.TaxDocuments
}

type CreateClusterOutput struct {

	// The automatically generated ID for a cluster.
	ClusterId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCluster{}, middleware.After)
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
	if err = addOpCreateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "CreateCluster",
	}
}
