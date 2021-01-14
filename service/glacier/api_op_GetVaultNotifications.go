// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation retrieves the notification-configuration subresource of the
// specified vault. For information about setting a notification configuration on a
// vault, see SetVaultNotifications. If a notification configuration for a vault is
// not set, the operation returns a 404 Not Found error. For more information about
// vault notifications, see Configuring Vault Notifications in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/configuring-notifications.html).
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions by
// default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access
// Management (IAM)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
// For conceptual information and underlying REST API, see Configuring Vault
// Notifications in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/configuring-notifications.html)
// and Get Vault Notification Configuration
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-get.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) GetVaultNotifications(ctx context.Context, params *GetVaultNotificationsInput, optFns ...func(*Options)) (*GetVaultNotificationsOutput, error) {
	if params == nil {
		params = &GetVaultNotificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVaultNotifications", params, optFns, addOperationGetVaultNotificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVaultNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for retrieving the notification configuration set on an Amazon
// Glacier vault.
type GetVaultNotificationsInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string
}

// Contains the Amazon S3 Glacier response to your request.
type GetVaultNotificationsOutput struct {

	// Returns the notification configuration set on the vault.
	VaultNotificationConfig *types.VaultNotificationConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetVaultNotificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVaultNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVaultNotifications{}, middleware.After)
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
	if err = addOpGetVaultNotificationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVaultNotifications(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetVaultNotifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "GetVaultNotifications",
	}
}
