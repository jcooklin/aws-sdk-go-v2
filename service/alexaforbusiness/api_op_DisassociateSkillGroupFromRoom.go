// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a skill group from a specified room. This disables all skills in
// the skill group on all devices in the room.
func (c *Client) DisassociateSkillGroupFromRoom(ctx context.Context, params *DisassociateSkillGroupFromRoomInput, optFns ...func(*Options)) (*DisassociateSkillGroupFromRoomOutput, error) {
	if params == nil {
		params = &DisassociateSkillGroupFromRoomInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateSkillGroupFromRoom", params, optFns, addOperationDisassociateSkillGroupFromRoomMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateSkillGroupFromRoomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateSkillGroupFromRoomInput struct {

	// The ARN of the room from which the skill group is to be disassociated. Required.
	RoomArn *string

	// The ARN of the skill group to disassociate from a room. Required.
	SkillGroupArn *string
}

type DisassociateSkillGroupFromRoomOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateSkillGroupFromRoomMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateSkillGroupFromRoom{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateSkillGroupFromRoom{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateSkillGroupFromRoom(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateSkillGroupFromRoom(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "DisassociateSkillGroupFromRoom",
	}
}
