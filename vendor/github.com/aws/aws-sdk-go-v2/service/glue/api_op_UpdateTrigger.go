// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateTriggerRequest
type UpdateTriggerInput struct {
	_ struct{} `type:"structure"`

	// The name of the trigger to update.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The new values with which to update the trigger.
	//
	// TriggerUpdate is a required field
	TriggerUpdate *TriggerUpdate `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateTriggerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTriggerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTriggerInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.TriggerUpdate == nil {
		invalidParams.Add(aws.NewErrParamRequired("TriggerUpdate"))
	}
	if s.TriggerUpdate != nil {
		if err := s.TriggerUpdate.Validate(); err != nil {
			invalidParams.AddNested("TriggerUpdate", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateTriggerResponse
type UpdateTriggerOutput struct {
	_ struct{} `type:"structure"`

	// The resulting trigger definition.
	Trigger *Trigger `type:"structure"`
}

// String returns the string representation
func (s UpdateTriggerOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTrigger = "UpdateTrigger"

// UpdateTriggerRequest returns a request value for making API operation for
// AWS Glue.
//
// Updates a trigger definition.
//
//    // Example sending a request using UpdateTriggerRequest.
//    req := client.UpdateTriggerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateTrigger
func (c *Client) UpdateTriggerRequest(input *UpdateTriggerInput) UpdateTriggerRequest {
	op := &aws.Operation{
		Name:       opUpdateTrigger,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTriggerInput{}
	}

	req := c.newRequest(op, input, &UpdateTriggerOutput{})
	return UpdateTriggerRequest{Request: req, Input: input, Copy: c.UpdateTriggerRequest}
}

// UpdateTriggerRequest is the request type for the
// UpdateTrigger API operation.
type UpdateTriggerRequest struct {
	*aws.Request
	Input *UpdateTriggerInput
	Copy  func(*UpdateTriggerInput) UpdateTriggerRequest
}

// Send marshals and sends the UpdateTrigger API request.
func (r UpdateTriggerRequest) Send(ctx context.Context) (*UpdateTriggerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTriggerResponse{
		UpdateTriggerOutput: r.Request.Data.(*UpdateTriggerOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTriggerResponse is the response type for the
// UpdateTrigger API operation.
type UpdateTriggerResponse struct {
	*UpdateTriggerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTrigger request.
func (r *UpdateTriggerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
