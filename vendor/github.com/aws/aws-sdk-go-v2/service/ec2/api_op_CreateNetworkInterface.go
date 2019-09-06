// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CreateNetworkInterface.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateNetworkInterfaceRequest
type CreateNetworkInterfaceInput struct {
	_ struct{} `type:"structure"`

	// A description for the network interface.
	Description *string `locationName:"description" type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The IDs of one or more security groups.
	Groups []string `locationName:"SecurityGroupId" locationNameList:"SecurityGroupId" type:"list"`

	// Indicates the type of network interface. To create an Elastic Fabric Adapter
	// (EFA), specify efa. For more information, see Elastic Fabric Adapter (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/efa.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	InterfaceType NetworkInterfaceCreationType `type:"string" enum:"true"`

	// The number of IPv6 addresses to assign to a network interface. Amazon EC2
	// automatically selects the IPv6 addresses from the subnet range. You can't
	// use this option if specifying specific IPv6 addresses. If your subnet has
	// the AssignIpv6AddressOnCreation attribute set to true, you can specify 0
	// to override this setting.
	Ipv6AddressCount *int64 `locationName:"ipv6AddressCount" type:"integer"`

	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your
	// subnet. You can't use this option if you're specifying a number of IPv6 addresses.
	Ipv6Addresses []InstanceIpv6Address `locationName:"ipv6Addresses" locationNameList:"item" type:"list"`

	// The primary private IPv4 address of the network interface. If you don't specify
	// an IPv4 address, Amazon EC2 selects one for you from the subnet's IPv4 CIDR
	// range. If you specify an IP address, you cannot indicate any IP addresses
	// specified in privateIpAddresses as primary (only one IP address can be designated
	// as primary).
	PrivateIpAddress *string `locationName:"privateIpAddress" type:"string"`

	// One or more private IPv4 addresses.
	PrivateIpAddresses []PrivateIpAddressSpecification `locationName:"privateIpAddresses" locationNameList:"item" type:"list"`

	// The number of secondary private IPv4 addresses to assign to a network interface.
	// When you specify a number of secondary IPv4 addresses, Amazon EC2 selects
	// these IP addresses within the subnet's IPv4 CIDR range. You can't specify
	// this option and specify more than one private IP address using privateIpAddresses.
	//
	// The number of IP addresses you can assign to a network interface varies by
	// instance type. For more information, see IP Addresses Per ENI Per Instance
	// Type (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI)
	// in the Amazon Virtual Private Cloud User Guide.
	SecondaryPrivateIpAddressCount *int64 `locationName:"secondaryPrivateIpAddressCount" type:"integer"`

	// The ID of the subnet to associate with the network interface.
	//
	// SubnetId is a required field
	SubnetId *string `locationName:"subnetId" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateNetworkInterfaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNetworkInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNetworkInterfaceInput"}

	if s.SubnetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of CreateNetworkInterface.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateNetworkInterfaceResult
type CreateNetworkInterfaceOutput struct {
	_ struct{} `type:"structure"`

	// Information about the network interface.
	NetworkInterface *NetworkInterface `locationName:"networkInterface" type:"structure"`
}

// String returns the string representation
func (s CreateNetworkInterfaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateNetworkInterface = "CreateNetworkInterface"

// CreateNetworkInterfaceRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a network interface in the specified subnet.
//
// For more information about network interfaces, see Elastic Network Interfaces
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html) in the
// Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateNetworkInterfaceRequest.
//    req := client.CreateNetworkInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateNetworkInterface
func (c *Client) CreateNetworkInterfaceRequest(input *CreateNetworkInterfaceInput) CreateNetworkInterfaceRequest {
	op := &aws.Operation{
		Name:       opCreateNetworkInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNetworkInterfaceInput{}
	}

	req := c.newRequest(op, input, &CreateNetworkInterfaceOutput{})
	return CreateNetworkInterfaceRequest{Request: req, Input: input, Copy: c.CreateNetworkInterfaceRequest}
}

// CreateNetworkInterfaceRequest is the request type for the
// CreateNetworkInterface API operation.
type CreateNetworkInterfaceRequest struct {
	*aws.Request
	Input *CreateNetworkInterfaceInput
	Copy  func(*CreateNetworkInterfaceInput) CreateNetworkInterfaceRequest
}

// Send marshals and sends the CreateNetworkInterface API request.
func (r CreateNetworkInterfaceRequest) Send(ctx context.Context) (*CreateNetworkInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNetworkInterfaceResponse{
		CreateNetworkInterfaceOutput: r.Request.Data.(*CreateNetworkInterfaceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNetworkInterfaceResponse is the response type for the
// CreateNetworkInterface API operation.
type CreateNetworkInterfaceResponse struct {
	*CreateNetworkInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNetworkInterface request.
func (r *CreateNetworkInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
