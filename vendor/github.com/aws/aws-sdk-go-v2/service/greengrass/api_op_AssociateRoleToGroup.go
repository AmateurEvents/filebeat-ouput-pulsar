// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/AssociateRoleToGroupRequest
type AssociateRoleToGroupInput struct {
	_ struct{} `type:"structure"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`

	// The ARN of the role you wish to associate with this group. The existence
	// of the role is not validated.
	//
	// RoleArn is a required field
	RoleArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateRoleToGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateRoleToGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateRoleToGroupInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateRoleToGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/AssociateRoleToGroupResponse
type AssociateRoleToGroupOutput struct {
	_ struct{} `type:"structure"`

	// The time, in milliseconds since the epoch, when the role ARN was associated
	// with the group.
	AssociatedAt *string `type:"string"`
}

// String returns the string representation
func (s AssociateRoleToGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateRoleToGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AssociatedAt != nil {
		v := *s.AssociatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AssociatedAt", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opAssociateRoleToGroup = "AssociateRoleToGroup"

// AssociateRoleToGroupRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Associates a role with a group. Your Greengrass core will use the role to
// access AWS cloud services. The role's permissions should allow Greengrass
// core Lambda functions to perform actions against the cloud.
//
//    // Example sending a request using AssociateRoleToGroupRequest.
//    req := client.AssociateRoleToGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/AssociateRoleToGroup
func (c *Client) AssociateRoleToGroupRequest(input *AssociateRoleToGroupInput) AssociateRoleToGroupRequest {
	op := &aws.Operation{
		Name:       opAssociateRoleToGroup,
		HTTPMethod: "PUT",
		HTTPPath:   "/greengrass/groups/{GroupId}/role",
	}

	if input == nil {
		input = &AssociateRoleToGroupInput{}
	}

	req := c.newRequest(op, input, &AssociateRoleToGroupOutput{})
	return AssociateRoleToGroupRequest{Request: req, Input: input, Copy: c.AssociateRoleToGroupRequest}
}

// AssociateRoleToGroupRequest is the request type for the
// AssociateRoleToGroup API operation.
type AssociateRoleToGroupRequest struct {
	*aws.Request
	Input *AssociateRoleToGroupInput
	Copy  func(*AssociateRoleToGroupInput) AssociateRoleToGroupRequest
}

// Send marshals and sends the AssociateRoleToGroup API request.
func (r AssociateRoleToGroupRequest) Send(ctx context.Context) (*AssociateRoleToGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateRoleToGroupResponse{
		AssociateRoleToGroupOutput: r.Request.Data.(*AssociateRoleToGroupOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateRoleToGroupResponse is the response type for the
// AssociateRoleToGroup API operation.
type AssociateRoleToGroupResponse struct {
	*AssociateRoleToGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateRoleToGroup request.
func (r *AssociateRoleToGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
