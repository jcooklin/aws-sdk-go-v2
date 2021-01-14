// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the association between an Amazon Lex bot and a
// messaging platform. This operation requires permissions for the
// lex:GetBotChannelAssociation action.
func (c *Client) GetBotChannelAssociation(ctx context.Context, params *GetBotChannelAssociationInput, optFns ...func(*Options)) (*GetBotChannelAssociationOutput, error) {
	if params == nil {
		params = &GetBotChannelAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBotChannelAssociation", params, optFns, addOperationGetBotChannelAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBotChannelAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotChannelAssociationInput struct {

	// An alias pointing to the specific version of the Amazon Lex bot to which this
	// association is being made.
	//
	// This member is required.
	BotAlias *string

	// The name of the Amazon Lex bot.
	//
	// This member is required.
	BotName *string

	// The name of the association between the bot and the channel. The name is case
	// sensitive.
	//
	// This member is required.
	Name *string
}

type GetBotChannelAssociationOutput struct {

	// An alias pointing to the specific version of the Amazon Lex bot to which this
	// association is being made.
	BotAlias *string

	// Provides information that the messaging platform needs to communicate with the
	// Amazon Lex bot.
	BotConfiguration map[string]string

	// The name of the Amazon Lex bot.
	BotName *string

	// The date that the association between the bot and the channel was created.
	CreatedDate *time.Time

	// A description of the association between the bot and the channel.
	Description *string

	// If status is FAILED, Amazon Lex provides the reason that it failed to create the
	// association.
	FailureReason *string

	// The name of the association between the bot and the channel.
	Name *string

	// The status of the bot channel.
	//
	// * CREATED - The channel has been created and is
	// ready for use.
	//
	// * IN_PROGRESS - Channel creation is in progress.
	//
	// * FAILED -
	// There was an error creating the channel. For information about the reason for
	// the failure, see the failureReason field.
	Status types.ChannelStatus

	// The type of the messaging platform.
	Type types.ChannelType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBotChannelAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBotChannelAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBotChannelAssociation{}, middleware.After)
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
	if err = addOpGetBotChannelAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBotChannelAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBotChannelAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBotChannelAssociation",
	}
}
