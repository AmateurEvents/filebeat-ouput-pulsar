// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/StartAppReplicationRequest
type StartAppReplicationInput struct {
	_ struct{} `type:"structure"`

	// ID of the application to replicate.
	AppId *string `locationName:"appId" type:"string"`
}

// String returns the string representation
func (s StartAppReplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/StartAppReplicationResponse
type StartAppReplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartAppReplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartAppReplication = "StartAppReplication"

// StartAppReplicationRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Starts replicating an application.
//
//    // Example sending a request using StartAppReplicationRequest.
//    req := client.StartAppReplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/StartAppReplication
func (c *Client) StartAppReplicationRequest(input *StartAppReplicationInput) StartAppReplicationRequest {
	op := &aws.Operation{
		Name:       opStartAppReplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartAppReplicationInput{}
	}

	req := c.newRequest(op, input, &StartAppReplicationOutput{})
	return StartAppReplicationRequest{Request: req, Input: input, Copy: c.StartAppReplicationRequest}
}

// StartAppReplicationRequest is the request type for the
// StartAppReplication API operation.
type StartAppReplicationRequest struct {
	*aws.Request
	Input *StartAppReplicationInput
	Copy  func(*StartAppReplicationInput) StartAppReplicationRequest
}

// Send marshals and sends the StartAppReplication API request.
func (r StartAppReplicationRequest) Send(ctx context.Context) (*StartAppReplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartAppReplicationResponse{
		StartAppReplicationOutput: r.Request.Data.(*StartAppReplicationOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartAppReplicationResponse is the response type for the
// StartAppReplication API operation.
type StartAppReplicationResponse struct {
	*StartAppReplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartAppReplication request.
func (r *StartAppReplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
