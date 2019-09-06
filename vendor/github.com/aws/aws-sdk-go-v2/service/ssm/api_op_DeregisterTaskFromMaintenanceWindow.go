// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeregisterTaskFromMaintenanceWindowRequest
type DeregisterTaskFromMaintenanceWindowInput struct {
	_ struct{} `type:"structure"`

	// The ID of the maintenance window the task should be removed from.
	//
	// WindowId is a required field
	WindowId *string `min:"20" type:"string" required:"true"`

	// The ID of the task to remove from the maintenance window.
	//
	// WindowTaskId is a required field
	WindowTaskId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterTaskFromMaintenanceWindowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterTaskFromMaintenanceWindowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterTaskFromMaintenanceWindowInput"}

	if s.WindowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowId"))
	}
	if s.WindowId != nil && len(*s.WindowId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowId", 20))
	}

	if s.WindowTaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowTaskId"))
	}
	if s.WindowTaskId != nil && len(*s.WindowTaskId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowTaskId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeregisterTaskFromMaintenanceWindowResult
type DeregisterTaskFromMaintenanceWindowOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the maintenance window the task was removed from.
	WindowId *string `min:"20" type:"string"`

	// The ID of the task removed from the maintenance window.
	WindowTaskId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s DeregisterTaskFromMaintenanceWindowOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterTaskFromMaintenanceWindow = "DeregisterTaskFromMaintenanceWindow"

// DeregisterTaskFromMaintenanceWindowRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Removes a task from a maintenance window.
//
//    // Example sending a request using DeregisterTaskFromMaintenanceWindowRequest.
//    req := client.DeregisterTaskFromMaintenanceWindowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeregisterTaskFromMaintenanceWindow
func (c *Client) DeregisterTaskFromMaintenanceWindowRequest(input *DeregisterTaskFromMaintenanceWindowInput) DeregisterTaskFromMaintenanceWindowRequest {
	op := &aws.Operation{
		Name:       opDeregisterTaskFromMaintenanceWindow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterTaskFromMaintenanceWindowInput{}
	}

	req := c.newRequest(op, input, &DeregisterTaskFromMaintenanceWindowOutput{})
	return DeregisterTaskFromMaintenanceWindowRequest{Request: req, Input: input, Copy: c.DeregisterTaskFromMaintenanceWindowRequest}
}

// DeregisterTaskFromMaintenanceWindowRequest is the request type for the
// DeregisterTaskFromMaintenanceWindow API operation.
type DeregisterTaskFromMaintenanceWindowRequest struct {
	*aws.Request
	Input *DeregisterTaskFromMaintenanceWindowInput
	Copy  func(*DeregisterTaskFromMaintenanceWindowInput) DeregisterTaskFromMaintenanceWindowRequest
}

// Send marshals and sends the DeregisterTaskFromMaintenanceWindow API request.
func (r DeregisterTaskFromMaintenanceWindowRequest) Send(ctx context.Context) (*DeregisterTaskFromMaintenanceWindowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterTaskFromMaintenanceWindowResponse{
		DeregisterTaskFromMaintenanceWindowOutput: r.Request.Data.(*DeregisterTaskFromMaintenanceWindowOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterTaskFromMaintenanceWindowResponse is the response type for the
// DeregisterTaskFromMaintenanceWindow API operation.
type DeregisterTaskFromMaintenanceWindowResponse struct {
	*DeregisterTaskFromMaintenanceWindowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterTaskFromMaintenanceWindow request.
func (r *DeregisterTaskFromMaintenanceWindowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
