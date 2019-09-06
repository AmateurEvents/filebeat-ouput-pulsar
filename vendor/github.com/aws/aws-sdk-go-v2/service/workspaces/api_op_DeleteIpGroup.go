// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DeleteIpGroupRequest
type DeleteIpGroupInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the IP access control group.
	//
	// GroupId is a required field
	GroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteIpGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteIpGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteIpGroupInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DeleteIpGroupResult
type DeleteIpGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteIpGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteIpGroup = "DeleteIpGroup"

// DeleteIpGroupRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Deletes the specified IP access control group.
//
// You cannot delete an IP access control group that is associated with a directory.
//
//    // Example sending a request using DeleteIpGroupRequest.
//    req := client.DeleteIpGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DeleteIpGroup
func (c *Client) DeleteIpGroupRequest(input *DeleteIpGroupInput) DeleteIpGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteIpGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteIpGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteIpGroupOutput{})
	return DeleteIpGroupRequest{Request: req, Input: input, Copy: c.DeleteIpGroupRequest}
}

// DeleteIpGroupRequest is the request type for the
// DeleteIpGroup API operation.
type DeleteIpGroupRequest struct {
	*aws.Request
	Input *DeleteIpGroupInput
	Copy  func(*DeleteIpGroupInput) DeleteIpGroupRequest
}

// Send marshals and sends the DeleteIpGroup API request.
func (r DeleteIpGroupRequest) Send(ctx context.Context) (*DeleteIpGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteIpGroupResponse{
		DeleteIpGroupOutput: r.Request.Data.(*DeleteIpGroupOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteIpGroupResponse is the response type for the
// DeleteIpGroup API operation.
type DeleteIpGroupResponse struct {
	*DeleteIpGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteIpGroup request.
func (r *DeleteIpGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
