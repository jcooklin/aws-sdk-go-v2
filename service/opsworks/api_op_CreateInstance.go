// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an instance in a specified stack. For more information, see Adding an
// Instance to a Layer
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-add.html).
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) CreateInstance(ctx context.Context, params *CreateInstanceInput, optFns ...func(*Options)) (*CreateInstanceOutput, error) {
	if params == nil {
		params = &CreateInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInstance", params, optFns, addOperationCreateInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstanceInput struct {

	// The instance type, such as t2.micro. For a list of supported instance types,
	// open the stack in the console, choose Instances, and choose + Instance. The Size
	// list contains the currently supported types. For more information, see Instance
	// Families and Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html). The
	// parameter values that you use to specify the various types are in the API Name
	// column of the Available Instance Types table.
	//
	// This member is required.
	InstanceType *string

	// An array that contains the instance's layer IDs.
	//
	// This member is required.
	LayerIds []string

	// The stack ID.
	//
	// This member is required.
	StackId *string

	// The default AWS OpsWorks Stacks agent version. You have the following
	// options:
	//
	// * INHERIT - Use the stack's default agent version setting.
	//
	// *
	// version_number - Use the specified agent version. This value overrides the
	// stack's default setting. To update the agent version, edit the instance
	// configuration and specify a new version. AWS OpsWorks Stacks then automatically
	// installs that version on the instance.
	//
	// The default setting is INHERIT. To
	// specify an agent version, you must use the complete version number, not the
	// abbreviated number shown on the console. For a list of available agent version
	// numbers, call DescribeAgentVersions. AgentVersion cannot be set to Chef 12.2.
	AgentVersion *string

	// A custom AMI ID to be used to create the instance. The AMI should be based on
	// one of the supported operating systems. For more information, see Using Custom
	// AMIs
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html).
	// If you specify a custom AMI, you must set Os to Custom.
	AmiId *string

	// The instance architecture. The default option is x86_64. Instance types do not
	// necessarily support both architectures. For a list of the architectures that are
	// supported by the different instance types, see Instance Families and Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html).
	Architecture types.Architecture

	// For load-based or time-based instances, the type. Windows stacks can use only
	// time-based instances.
	AutoScalingType types.AutoScalingType

	// The instance Availability Zone. For more information, see Regions and Endpoints
	// (https://docs.aws.amazon.com/general/latest/gr/rande.html).
	AvailabilityZone *string

	// An array of BlockDeviceMapping objects that specify the instance's block
	// devices. For more information, see Block Device Mapping
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html).
	// Note that block device mappings are not supported for custom AMIs.
	BlockDeviceMappings []types.BlockDeviceMapping

	// Whether to create an Amazon EBS-optimized instance.
	EbsOptimized *bool

	// The instance host name.
	Hostname *string

	// Whether to install operating system and package updates when the instance boots.
	// The default value is true. To control when updates are installed, set this value
	// to false. You must then update your instances manually by using CreateDeployment
	// to run the update_dependencies stack command or by manually running yum (Amazon
	// Linux) or apt-get (Ubuntu) on the instances. We strongly recommend using the
	// default value of true to ensure that your instances have the latest security
	// updates.
	InstallUpdatesOnBoot *bool

	// The instance's operating system, which must be set to one of the following.
	//
	// * A
	// supported Linux operating system: An Amazon Linux version, such as Amazon Linux
	// 2018.03, Amazon Linux 2017.09, Amazon Linux 2017.03, Amazon Linux 2016.09,
	// Amazon Linux 2016.03, Amazon Linux 2015.09, or Amazon Linux 2015.03.
	//
	// * A
	// supported Ubuntu operating system, such as Ubuntu 16.04 LTS, Ubuntu 14.04 LTS,
	// or Ubuntu 12.04 LTS.
	//
	// * CentOS Linux 7
	//
	// * Red Hat Enterprise Linux 7
	//
	// * A
	// supported Windows operating system, such as Microsoft Windows Server 2012 R2
	// Base, Microsoft Windows Server 2012 R2 with SQL Server Express, Microsoft
	// Windows Server 2012 R2 with SQL Server Standard, or Microsoft Windows Server
	// 2012 R2 with SQL Server Web.
	//
	// * A custom AMI: Custom.
	//
	// For more information
	// about the supported operating systems, see AWS OpsWorks Stacks Operating Systems
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html).
	// The default option is the current Amazon Linux version. If you set this
	// parameter to Custom, you must use the CreateInstance action's AmiId parameter to
	// specify the custom AMI that you want to use. Block device mappings are not
	// supported if the value is Custom. For more information about supported operating
	// systems, see Operating Systems
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html)For
	// more information about how to use custom AMIs with AWS OpsWorks Stacks, see
	// Using Custom AMIs
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html).
	Os *string

	// The instance root device type. For more information, see Storage for the Root
	// Device
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ComponentsAMIs.html#storage-for-the-root-device).
	RootDeviceType types.RootDeviceType

	// The instance's Amazon EC2 key-pair name.
	SshKeyName *string

	// The ID of the instance's subnet. If the stack is running in a VPC, you can use
	// this parameter to override the stack's default subnet ID value and direct AWS
	// OpsWorks Stacks to launch the instance in a different subnet.
	SubnetId *string

	// The instance's tenancy option. The default option is no tenancy, or if the
	// instance is running in a VPC, inherit tenancy settings from the VPC. The
	// following are valid values for this parameter: dedicated, default, or host.
	// Because there are costs associated with changes in tenancy options, we recommend
	// that you research tenancy options before choosing them for your instances. For
	// more information about dedicated hosts, see Dedicated Hosts Overview
	// (http://aws.amazon.com/ec2/dedicated-hosts/) and Amazon EC2 Dedicated Hosts
	// (http://aws.amazon.com/ec2/dedicated-hosts/). For more information about
	// dedicated instances, see Dedicated Instances
	// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/dedicated-instance.html)
	// and Amazon EC2 Dedicated Instances
	// (http://aws.amazon.com/ec2/purchasing-options/dedicated-instances/).
	Tenancy *string

	// The instance's virtualization type, paravirtual or hvm.
	VirtualizationType *string
}

// Contains the response to a CreateInstance request.
type CreateInstanceOutput struct {

	// The instance ID.
	InstanceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateInstance{}, middleware.After)
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
	if err = addOpCreateInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "CreateInstance",
	}
}
