// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateScriptInput
type CreateScriptInput struct {
	_ struct{} `type:"structure"`

	// Descriptive label that is associated with a script. Script names do not need
	// to be unique. You can use UpdateScript to change this value later.
	Name *string `min:"1" type:"string"`

	// Location of the Amazon S3 bucket where a zipped file containing your Realtime
	// scripts is stored. The storage location must specify the Amazon S3 bucket
	// name, the zip file name (the "key"), and a role ARN that allows Amazon GameLift
	// to access the Amazon S3 storage location. The S3 bucket must be in the same
	// region where you want to create a new script. By default, Amazon GameLift
	// uploads the latest version of the zip file; if you have S3 object versioning
	// turned on, you can use the ObjectVersion parameter to specify an earlier
	// version.
	StorageLocation *S3Location `type:"structure"`

	// Version that is associated with a build or script. Version strings do not
	// need to be unique. You can use UpdateScript to change this value later.
	Version *string `min:"1" type:"string"`

	// Data object containing your Realtime scripts and dependencies as a zip file.
	// The zip file can have one or multiple files. Maximum size of a zip file is
	// 5 MB.
	//
	// When using the AWS CLI tool to create a script, this parameter is set to
	// the zip file name. It must be prepended with the string "fileb://" to indicate
	// that the file data is a binary object. For example: --zip-file fileb://myRealtimeScript.zip.
	//
	// ZipFile is automatically base64 encoded/decoded by the SDK.
	ZipFile []byte `type:"blob"`
}

// String returns the string representation
func (s CreateScriptInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateScriptInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateScriptInput"}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}
	if s.StorageLocation != nil {
		if err := s.StorageLocation.Validate(); err != nil {
			invalidParams.AddNested("StorageLocation", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateScriptOutput
type CreateScriptOutput struct {
	_ struct{} `type:"structure"`

	// The newly created script record with a unique script ID. The new script's
	// storage location reflects an Amazon S3 location: (1) If the script was uploaded
	// from an S3 bucket under your account, the storage location reflects the information
	// that was provided in the CreateScript request; (2) If the script file was
	// uploaded from a local zip file, the storage location reflects an S3 location
	// controls by the Amazon GameLift service.
	Script *Script `type:"structure"`
}

// String returns the string representation
func (s CreateScriptOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateScript = "CreateScript"

// CreateScriptRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Creates a new script record for your Realtime Servers script. Realtime scripts
// are JavaScript that provide configuration settings and optional custom game
// logic for your game. The script is deployed when you create a Realtime Servers
// fleet to host your game sessions. Script logic is executed during an active
// game session.
//
// To create a new script record, specify a script name and provide the script
// file(s). The script files and all dependencies must be zipped into a single
// file. You can pull the zip file from either of these locations:
//
//    * A locally available directory. Use the ZipFile parameter for this option.
//
//    * An Amazon Simple Storage Service (Amazon S3) bucket under your AWS account.
//    Use the StorageLocation parameter for this option. You'll need to have
//    an Identity Access Management (IAM) role that allows the Amazon GameLift
//    service to access your S3 bucket.
//
// If the call is successful, a new script record is created with a unique script
// ID. If the script file is provided as a local file, the file is uploaded
// to an Amazon GameLift-owned S3 bucket and the script record's storage location
// reflects this location. If the script file is provided as an S3 bucket, Amazon
// GameLift accesses the file at this storage location as needed for deployment.
//
// Learn more
//
// Amazon GameLift Realtime Servers (https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-intro.html)
//
// Set Up a Role for Amazon GameLift Access (https://docs.aws.amazon.com/gamelift/latest/developerguide/setting-up-role.html)
//
// Related operations
//
//    * CreateScript
//
//    * ListScripts
//
//    * DescribeScript
//
//    * UpdateScript
//
//    * DeleteScript
//
//    // Example sending a request using CreateScriptRequest.
//    req := client.CreateScriptRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateScript
func (c *Client) CreateScriptRequest(input *CreateScriptInput) CreateScriptRequest {
	op := &aws.Operation{
		Name:       opCreateScript,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateScriptInput{}
	}

	req := c.newRequest(op, input, &CreateScriptOutput{})
	return CreateScriptRequest{Request: req, Input: input, Copy: c.CreateScriptRequest}
}

// CreateScriptRequest is the request type for the
// CreateScript API operation.
type CreateScriptRequest struct {
	*aws.Request
	Input *CreateScriptInput
	Copy  func(*CreateScriptInput) CreateScriptRequest
}

// Send marshals and sends the CreateScript API request.
func (r CreateScriptRequest) Send(ctx context.Context) (*CreateScriptResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateScriptResponse{
		CreateScriptOutput: r.Request.Data.(*CreateScriptOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateScriptResponse is the response type for the
// CreateScript API operation.
type CreateScriptResponse struct {
	*CreateScriptOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateScript request.
func (r *CreateScriptResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}