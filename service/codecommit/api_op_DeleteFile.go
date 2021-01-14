// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a specified file from a specified branch. A commit is created on the
// branch that contains the revision. The file still exists in the commits earlier
// to the commit that contains the deletion.
func (c *Client) DeleteFile(ctx context.Context, params *DeleteFileInput, optFns ...func(*Options)) (*DeleteFileOutput, error) {
	if params == nil {
		params = &DeleteFileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFile", params, optFns, addOperationDeleteFileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFileInput struct {

	// The name of the branch where the commit that deletes the file is made.
	//
	// This member is required.
	BranchName *string

	// The fully qualified path to the file that to be deleted, including the full name
	// and extension of that file. For example, /examples/file.md is a fully qualified
	// path to a file named file.md in a folder named examples.
	//
	// This member is required.
	FilePath *string

	// The ID of the commit that is the tip of the branch where you want to create the
	// commit that deletes the file. This must be the HEAD commit for the branch. The
	// commit that deletes the file is created from this commit ID.
	//
	// This member is required.
	ParentCommitId *string

	// The name of the repository that contains the file to delete.
	//
	// This member is required.
	RepositoryName *string

	// The commit message you want to include as part of deleting the file. Commit
	// messages are limited to 256 KB. If no message is specified, a default message is
	// used.
	CommitMessage *string

	// The email address for the commit that deletes the file. If no email address is
	// specified, the email address is left blank.
	Email *string

	// If a file is the only object in the folder or directory, specifies whether to
	// delete the folder or directory that contains the file. By default, empty folders
	// are deleted. This includes empty folders that are part of the directory
	// structure. For example, if the path to a file is dir1/dir2/dir3/dir4, and dir2
	// and dir3 are empty, deleting the last file in dir4 also deletes the empty
	// folders dir4, dir3, and dir2.
	KeepEmptyFolders bool

	// The name of the author of the commit that deletes the file. If no name is
	// specified, the user's ARN is used as the author name and committer name.
	Name *string
}

type DeleteFileOutput struct {

	// The blob ID removed from the tree as part of deleting the file.
	//
	// This member is required.
	BlobId *string

	// The full commit ID of the commit that contains the change that deletes the file.
	//
	// This member is required.
	CommitId *string

	// The fully qualified path to the file to be deleted, including the full name and
	// extension of that file.
	//
	// This member is required.
	FilePath *string

	// The full SHA-1 pointer of the tree information for the commit that contains the
	// delete file change.
	//
	// This member is required.
	TreeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteFileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteFile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteFile{}, middleware.After)
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
	if err = addOpDeleteFileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "DeleteFile",
	}
}
