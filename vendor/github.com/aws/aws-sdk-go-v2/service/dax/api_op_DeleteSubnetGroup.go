// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DeleteSubnetGroupRequest
type DeleteSubnetGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the subnet group to delete.
	//
	// SubnetGroupName is a required field
	SubnetGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSubnetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSubnetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSubnetGroupInput"}

	if s.SubnetGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DeleteSubnetGroupResponse
type DeleteSubnetGroupOutput struct {
	_ struct{} `type:"structure"`

	// A user-specified message for this action (i.e., a reason for deleting the
	// subnet group).
	DeletionMessage *string `type:"string"`
}

// String returns the string representation
func (s DeleteSubnetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSubnetGroup = "DeleteSubnetGroup"

// DeleteSubnetGroupRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Deletes a subnet group.
//
// You cannot delete a subnet group if it is associated with any DAX clusters.
//
//    // Example sending a request using DeleteSubnetGroupRequest.
//    req := client.DeleteSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DeleteSubnetGroup
func (c *Client) DeleteSubnetGroupRequest(input *DeleteSubnetGroupInput) DeleteSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteSubnetGroupOutput{})
	return DeleteSubnetGroupRequest{Request: req, Input: input, Copy: c.DeleteSubnetGroupRequest}
}

// DeleteSubnetGroupRequest is the request type for the
// DeleteSubnetGroup API operation.
type DeleteSubnetGroupRequest struct {
	*aws.Request
	Input *DeleteSubnetGroupInput
	Copy  func(*DeleteSubnetGroupInput) DeleteSubnetGroupRequest
}

// Send marshals and sends the DeleteSubnetGroup API request.
func (r DeleteSubnetGroupRequest) Send(ctx context.Context) (*DeleteSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSubnetGroupResponse{
		DeleteSubnetGroupOutput: r.Request.Data.(*DeleteSubnetGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSubnetGroupResponse is the response type for the
// DeleteSubnetGroup API operation.
type DeleteSubnetGroupResponse struct {
	*DeleteSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSubnetGroup request.
func (r *DeleteSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
